package kipris

import "encoding/xml"

type unit struct {
	XMLName xml.Name `xml:"item"`
	Claim   string   `xml:"claim"`
}
