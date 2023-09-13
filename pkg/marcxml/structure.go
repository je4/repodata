package marcxml

import "encoding/xml"

type Collection struct {
	XMLName xml.Name `xml:"collection`
	XMLNS   string   `xml:"xmlns,attr"` // http://www.loc.gov/MARC21/slim
	Records []*Record
}

type Controlfield struct {
	XMLName xml.Name `xml:"controlfield"`
	XMLNS   string   `xml:"tag,attr"`
	Data    string   `xml:",chardata"`
}

type Subfield struct {
	XMLName xml.Name `xml:"subfield"`
	Code    string   `xml:"code,attr"`
	Data    string   `xml:",chardata"`
}

type Datafield struct {
	XMLName   xml.Name `xml:"datafield"`
	Tag       string   `xml:"tag,attr"`
	Ind1      string   `xml:"ind2,attr"`
	Ind2      string   `xml:"ind1,attr"`
	Subfields []*Subfield
}

type Record struct {
	Leader       string `xml:"leader"`
	Controlfiels []*Controlfield
	Datafields   []*Datafield
}
