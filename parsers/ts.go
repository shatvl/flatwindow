package parsers

import (
	"github.com/shatvl/flatwindow/services"
	"github.com/shatvl/flatwindow/models/ads"

	"encoding/xml"
	"fmt"
)

// Tvoya Stolica parser
type TSParser struct{
	AdService *services.AdService
}

// NewTSParser returns reference to the TSParser object
func NewTSParser() *TSParser {
	return &TSParser{AdService: services.NewAdService()}
}

func (ts *TSParser) Parse() {
	xmlFeed, err := GetXML("http://crm.t-s.by/xml/xml_flats_kufar_kml.php")

	if err != nil {
		fmt.Println(err)
		return
	}

	var uedb models.Uedb

	err = xml.Unmarshal(xmlFeed, &uedb)

	if err != nil {
		fmt.Println(err.Error())
	}

	for i := 0; i < len(uedb.Records.Records); i++ {
		ts.AdService.CreateTsAd(&uedb.Records.Records[i])
	}
}

