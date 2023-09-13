package main

import (
	"emperror.dev/errors"
	"github.com/BurntSushi/toml"
	elasticutils "github.com/je4/utils/v2/pkg/elasticsearch"
	"io/fs"
)

type RepoConfig struct {
	Elasticsearch *elasticutils.Elastic8Config
	Index         string
}

func LoadConfig(fsys fs.FS, path string) (*RepoConfig, error) {
	data, err := fs.ReadFile(fsys, path)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot read %s:%s", fsys, path)
	}
	cfg := &RepoConfig{}
	if _, err := toml.Decode(string(data), cfg); err != nil {
		return nil, errors.Wrapf(err, "cannot decode %s:%s", fsys, path)
	}
	return cfg, nil
}
