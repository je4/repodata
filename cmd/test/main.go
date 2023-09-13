package main

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/je4/repodata/v2/pkg/marcxml"
	"github.com/je4/repodata/v2/pkg/structure"
	elasticutils "github.com/je4/utils/v2/pkg/elasticsearch"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"path/filepath"
	"regexp"
)

var cfgFile = flag.String("cfg", "", "path to toml configuration file")

var subRegexp = regexp.MustCompile(`^\|([a-z0-9]+)\s+(.*)$`)

func main() {
	var x = []string{"hello", "hello2"}
	data, _ := json.Marshal(x)
	json.Unmarshal(data, &x)

	flag.Parse()
	zLogger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr})

	cfg, err := LoadConfig(os.DirFS(filepath.Dir(*cfgFile)), filepath.Base(*cfgFile))

	es8, err := elasticutils.NewClient2(cfg.Index, cfg.Elasticsearch, &zLogger)
	if err != nil {
		log.Panic().Stack().Err(err).Msg("cannot connect to elasticsearch")
	}

	query := &types.Query{
		Term: map[string]types.TermQuery{
			"flags": {Value: "e-rara"},
		},
	}
	srch, err := es8.Search().Query(query).Scroll("2m").Do(context.Background())
	if err != nil {
		log.Panic().Stack().Err(err).Msg("cannot query")
	}
	if srch.ScrollId_ == nil {
		log.Panic().Msg("no scrollid in query")
	}
	hits := srch.Hits.Hits
	scrollID := srch.ScrollId_
	log.Info().Msgf("hits: %v", len(hits))
	for {
		for _, hit := range hits {
			fmt.Println("id", hit.Id_)
			source := hit.Source_
			dataStruct := &structure.Default{}
			if err := json.Unmarshal(source, dataStruct); err != nil {
				log.Panic().Stack().Err(err).Msg("cannot unmarshal")
			}
			coll := &marcxml.Collection{
				XMLName: xml.Name{},
				XMLNS:   "http://www.loc.gov/MARC21/slim",
				Records: make([]*marcxml.Record, 0),
			}
			rec := &marcxml.Record{
				Leader:       dataStruct.Ldr.LeaderFull,
				Controlfiels: make([]*marcxml.Controlfield, 0), // []*marcxml.Controlfield{}
				Datafields:   make([]*marcxml.Datafield, 0),
			}
			for _, df := range dataStruct.Datafield {
				newDF := &marcxml.Datafield{
					XMLName:   xml.Name{},
					Tag:       df.Tag,
					Ind1:      df.Ind1,
					Ind2:      df.Ind2,
					Subfields: make([]*marcxml.Subfield, 0),
				}
				for _, sf := range df.Subfield {
					found := subRegexp.FindStringSubmatch(sf)
					if found == nil {
						zLogger.Error().Msgf("%s not a subfield", sf)
						break
					}
					newDF.Subfields = append(newDF.Subfields, &marcxml.Subfield{
						XMLName: xml.Name{},
						Code:    found[1],
						Data:    found[2],
					})
				}
				rec.Datafields = append(rec.Datafields, newDF)
			}
			coll.Records = append(coll.Records, rec)
			data, err := xml.MarshalIndent(coll, "", "  ")
			if err != nil {
				zLogger.Panic().Err(err)
			}
			//			fmt.Println(string(data))
			_ = data
		}
		if scrollID == nil {
			break
		}
		scroller, err := es8.Scroll().ScrollId(*srch.ScrollId_).Do(context.Background())
		if err != nil {
			log.Panic().Stack().Err(err).Msg("cannot scroll")
		}
		hits = scroller.Hits.Hits
	}
}
