package parser

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"

	mgo "gopkg.in/mgo.v2"
)

type Uedb struct {
	XMLName xml.Name `xml:"uedb"`
	Records Records  `xml:"records"`
}

type Records struct {
	XMLName xml.Name `xml:"records"`
	Records []Record `xml:"record"`
}

type Photos struct {
	XMLName xml.Name `xml:"photos"`
	Photos  []Photo  `xml:"photo"`
}

type Photo struct {
	XMLName xml.Name `xml:"photo"`
	Picture string   `xml:"picture,attr"`
}

type Record struct {
	XMLName          xml.Name `xml:"record"`
	Unid             string   `xml:"unid,attr"`
	Subject          string   `xml:"subject"`
	RemunerationType byte     `xml:"remuneration_type"`
	Link             string   `xml:"link"`
	Phone            string   `xml:"phone"`
	ContactPerson    string   `xml:"contact_person"`
	Photos           Photos   `xml:"photos"`
	Condition        byte     `xml:"condition"`
	Address          string   `xml:"address"`
	Area             int16    `xml:"area"`
	Region           byte     `xml:"region"`
	Metro            byte     `xml:"metro"`
	Rooms            byte     `xml:"rooms"`
	Category         int32    `xml:"category"`
	Type             string   `xml:"type"`
	YearBuilt        int16    `xml:"year_built"`
	Floor            byte     `xml:"floor"`
	HouseType        byte     `xml:"house_type"`
	Size             float32  `xml:"size"`
	SizeLivingSpace  float32  `xml:"size_living_space"`
	SizeKitchen      float32  `xml:"size_kitchen"`
	Body             string   `xml:"body"`
	Price            float32  `xml:"price"`
	Currency         string   `xml:"currency"`
}

// Parser feed interface
type Parser interface {
	Parse(url string)
}

// Tvoya Stolica parser
type TSParser struct{}

func NewParser() *TSParser {
	return &TSParser{}
}

func (ts *TSParser) Parse(session *mgo.Session) {
	xmlFeed, err := getXML("http://crm.t-s.by/xml/xml_flats_kufar_kml.php")

	if err != nil {
		fmt.Println(err)
		return
	}

	var uedb Uedb

	err = xml.Unmarshal(xmlFeed, &uedb)

	if err != nil {
		fmt.Println(err.Error())
	}

	for i := 0; i < len(uedb.Records.Records); i++ {
		session.DB("").C("ads").Insert(uedb.Records.Records[i])
		fmt.Println("Record subject: " + uedb.Records.Records[i].Subject)
	}
}

func getXML(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, fmt.Errorf("GET error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return []byte{}, fmt.Errorf("Status error: %v", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("Read body: %v", err)
	}

	return data, nil
}
