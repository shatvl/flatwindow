package models

import (
	"encoding/xml"
	"gopkg.in/mgo.v2/bson"
)

type Uedb struct {
	XMLName xml.Name `xml:"uedb" json:"-"`
	Records Records  `xml:"records" json:"records"`
}

type Records struct {
	XMLName xml.Name `xml:"records" json:"-"`
	Records []Ad `xml:"record" json:"records"`
}

type Photos struct {
	XMLName xml.Name `xml:"photos" bson:"-" json:"-"`
	Photos  []Photo  `xml:"photo" json:"photos"`
}

type Photo struct {
	XMLName xml.Name `xml:"photo" bson:"-" json:"-"`
	Picture string   `xml:"picture,attr" json:"picture" bson:"picture"`
}

type Ad struct {
	XMLName          xml.Name 	    `xml:"record" bson:"-" json:"-"`
	ID               bson.ObjectId  `json:"id" bson:"_id,omitempty"`
	Unid             string   		`xml:"unid,attr" json:"unid"`
	Subject          string   		`xml:"subject" json:"subject"`
	RemunerationType byte     		`xml:"remuneration_type" json:"remunerationType" bson:"remunerationType"`
	Link             string   		`xml:"link" json:"link"`
	Phone            string   		`xml:"phone" json:"phone"`
	ContactPerson    string   		`xml:"contact_person" json:"contactPerson" bson:"contactPerson"`
	Photos           Photos   		`xml:"photos" json:"photos"`
	Condition        byte     		`xml:"condition" json:"condition"`
	Address          string   		`xml:"address" json:"address"`
	Area             int16    		`xml:"area" json:"area"`
	Region           byte     		`xml:"region" json:"region"`
	Metro            byte     		`xml:"metro" json:"metro"`
	Rooms            byte     		`xml:"rooms" json:"rooms"`
	Category         int32    		`xml:"category" json:"category"`
	Type             string   		`xml:"type" json:"type"`
	YearBuilt        int16    		`xml:"year_built" json:"yearBuilt" bson:"yearBuild"`
	Floor            byte     		`xml:"floor" json:"floor"`
	HouseType        byte     		`xml:"house_type" json:"houseType"`
	Size             float32  		`xml:"size" json:"size"`
	SizeLivingSpace  float32  		`xml:"size_living_space" json:"sizeLivingSpace" bson:"sizeLivingSpace"`
	SizeKitchen      float32  		`xml:"size_kitchen" json:"sizeKitchen" bson:"sizeKitchen"`
	Body             string   		`xml:"body" json:"body"`
	Price            float32  		`xml:"price" json:"price"`
	Currency         string   		`xml:"currency" json:"currency"`
	ModifiedDate	 string   		`xml:"modified_date" json:"modifiedDate" bson:"modifiedDate"`
	AgentType		 byte     		`json:"agentType" bson:"agentType"`
}