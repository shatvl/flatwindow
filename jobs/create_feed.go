package jobs

import (
	"github.com/shatvl/flatwindow/services"
	"github.com/shatvl/flatwindow/repositories"
	"os"
	"encoding/xml"
	"github.com/shatvl/flatwindow/models"
)

type Feed struct {
	AdService *services.AdService
}

func NewFeed() *Feed {
	return &Feed{AdService: services.NewAdService()}
}

type XmlFeed struct {
	XMLName xml.Name    `xml:"records"`
	Records []*models.Ad `xml:"record"`
}

func (f *Feed) CreateTSFeed() {
	ads, err := f.AdService.GetAgentAdsForFeedByCode(repositories.TsTypeCode)

	if err != nil {
		return
	}

	xmlFeed := &XmlFeed{Records: ads}
	adsXml, err := xml.MarshalIndent(xmlFeed, " ", "  ")

	if err != nil {
		return
	}

	fout, err := os.Create("public/xml/" + "tsfeed.xml")
	defer fout.Close()

	if err != nil {
		return
	}

	fout.Write([]byte(xml.Header))
	fout.Write(adsXml)
}