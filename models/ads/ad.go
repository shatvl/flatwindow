package models

import (
	"encoding/xml"
)

type Uedb struct {
	XMLName xml.Name `xml:"uedb" json:"-"`
	Records Records  `xml:"records"`
}

type Records struct {
	XMLName xml.Name `xml:"records" json:"-"`
	Records []Ad `xml:"record"`
}

type Photos struct {
	XMLName xml.Name `xml:"photos" bson:"-" json:"-"`
	Photos  []Photo  `xml:"photo"`
}

type Photo struct {
	XMLName xml.Name `xml:"photo" bson:"-" json:"-"`
	Picture string   `xml:"picture,attr"`
}

type Ad struct {
	XMLName          xml.Name `xml:"record" bson:"-" json:"-"`
	Unid             string   `xml:"unid,attr"`
	Subject          string   `xml:"subject"`
	RemunerationType byte     `xml:"remuneration_type" json:"remunerationType" bson:"remunerationType"`
	Link             string   `xml:"link"`
	Phone            string   `xml:"phone"`
	ContactPerson    string   `xml:"contact_person" json:"contactPerson" bson:"contactPerson"`
	Photos           Photos   `xml:"photos"`
	Condition        byte     `xml:"condition"`
	Address          string   `xml:"address"`
	Area             int16    `xml:"area"`
	Region           byte     `xml:"region"`
	Metro            byte     `xml:"metro"`
	Rooms            byte     `xml:"rooms"`
	Category         int32    `xml:"category"`
	Type             string   `xml:"type"`
	YearBuilt        int16    `xml:"year_built" json:"yearBuilt" bson:"yearBuild"`
	Floor            byte     `xml:"floor"`
	HouseType        byte     `xml:"house_type"`
	Size             float32  `xml:"size"`
	SizeLivingSpace  float32  `xml:"size_living_space" json:"sizeLivingSpace" bson:"sizeLivingSpace"`
	SizeKitchen      float32  `xml:"size_kitchen" json:"sizeKitchen" bson:"sizeKitchen"`
	Body             string   `xml:"body"`
	Price            float32  `xml:"price"`
	Currency         string   `xml:"currency"`
	ModifiedDate	 string   `xml:"modified_date" json:"modifiedDate" bson:"modifiedDate"`
	AgentType		 byte     `json:"agentType" bson:"agentType"`
}