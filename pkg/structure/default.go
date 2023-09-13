package structure

import "time"

type LDR struct {
	Leader06TypeOfRecord        string `json:"leader_06_typeOfRecord"`
	Leader07BibliographicStatus string `json:"leader_07_bibliographicStatus"`
	Leader17EncodingLevel       string `json:"leader_17_encodingLevel"`
	Leader19MultipartLevel      string `json:"leader_19_multipartLevel"`
	LeaderFull                  string `json:"leader_full"`
}
type Controlfield struct {
	Zero01ID                    string `json:"001_id"`
	Zero05LatestTransactionTime string `json:"005_latestTransactionTime"`
	Zero07Full                  string `json:"007_full"`
	Zero0806TypeOfDate          string `json:"008_06_typeOfDate"`
	Zero080710DateFirst         string `json:"008_07-10_dateFirst"`
	Zero081114DateSecond        string `json:"008_11-14_dateSecond"`
	Zero081517Country           string `json:"008_15-17_country"`
	Zero0823Or29FormOfItem      string `json:"008_23or29_formOfItem"`
	Zero083537Language          string `json:"008_35-37_language"`
	Zero08Full                  string `json:"008_full"`
}

type Datafield struct {
	Tag      string   `json:"tag"`
	Ind1     string   `json:"ind1"`
	Ind2     string   `json:"ind2"`
	Subfield []string `json:"subfield"`
}

type StringMapStrings map[string][]string
type StringMapString map[string]string

type Location map[string][]StringMapStrings

type ACL struct {
	Content []string `json:"content"`
	Meta    []string `json:"meta"`
	Preview []string `json:"preview"`
}

type Default struct {
	Timestamp    time.Time    `json:"timestamp"`
	Ldr          LDR          `json:"LDR"`
	Controlfield Controlfield `json:"controlfield"`
	Datafield    []*Datafield `json:"datafield"`
	Fieldlists   any          `json:"fieldlists"`
	Mapping      struct {
		Extension  StringMapStrings             `json:"extension"`
		Identifier StringMapStrings             `json:"identifier"`
		Language   []string                     `json:"language"`
		Location   map[string][]StringMapString `json:"location"`
		Name       struct {
			Geographic []struct {
				Identifier string `json:"identifier"`
				NamePart   string `json:"namePart"`
			} `json:"geographic"`
			Personal []struct {
				Gnd struct {
					NamePart   string   `json:"namePart"`
					Date       string   `json:"date"`
					Role       []string `json:"role"`
					Identifier string   `json:"identifier"`
				} `json:"gnd"`
			} `json:"personal"`
		} `json:"name"`
		Note struct {
			Binding []struct {
				Main string `json:"main"`
			} `json:"binding"`
			Citation []struct {
				Add  string `json:"add,omitempty"`
				Main string `json:"main"`
			} `json:"citation"`
			General   []string `json:"general"`
			Ownership []struct {
				Main string `json:"main"`
			} `json:"ownership"`
			Publications              []string `json:"publications"`
			StatementOfResponsibility []string `json:"statementOfResponsibility"`
			VersionIdentification     []string `json:"versionIdentification"`
		} `json:"note"`
		OriginInfo struct {
			Publication []struct {
				Date      string   `json:"date"`
				Place     []string `json:"place"`
				Publisher []string `json:"publisher"`
			} `json:"publication"`
		} `json:"originInfo"`
		PhysicalDescription struct {
			Extent []struct {
				Extent string `json:"extent"`
			} `json:"extent"`
		} `json:"physicalDescription"`
		RecordIdentifier []string `json:"recordIdentifier"`
		TitleInfo        struct {
			Main []struct {
				SubTitle string `json:"subTitle"`
				Title    string `json:"title"`
			} `json:"main"`
		} `json:"titleInfo"`
	} `json:"mapping"`
	Facet struct {
		Strings []struct {
			Name   string   `json:"name"`
			String []string `json:"string"`
		} `json:"strings"`
		Objects []struct {
			Name    string `json:"name"`
			Objects []struct {
				Identifier []string `json:"Identifier"`
				Name       string   `json:"Name"`
				Role       []string `json:"Role"`
			} `json:"objects"`
		} `json:"objects"`
	} `json:"facet"`
	Sets  []string `json:"sets"`
	Flags []string `json:"flags"`
	ACL   ACL      `json:"acl"`
}
