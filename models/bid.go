package models

import (
	"gopkg.in/mgo.v2/bson"
)

const (
	DetailedQuery = iota
	ReviewQuery
)

type Bid struct {
	AdId        bson.ObjectId `json:"ad_id" bson:"ad_id,omitempty"`
	Type        byte          `json:"type" bson:"type"`
	Fullname    string        `json:"fullname" bson:"fullname"`
	Email       string        `json:"email" bson:"email"`
	Phone       string        `json:"phone" bson:"phone"`
	City        string        `json:"city" bson:"city"`
	Message     string        `json:"message" bson:"message"`
	AboutMeType byte          `json:"about_me_type" bson:"about_me_type"`
	Campaign    bool          `json:"campaign" bson:"-"`
	CopyEmail   bool          `json:"copy_email" bson:"-"`
	PriceChange bool          `json:"price_change" bson:"price_change"`
}