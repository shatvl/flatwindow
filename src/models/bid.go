package models

import (
	"gopkg.in/mgo.v2/bson"
)

const (
	DetailedQuery = iota
	ReviewQuery
)

type Bid struct {
	ID          bson.ObjectId `json:"id" bson:"_id,omitempty"`
	AdId        bson.ObjectId `json:"adId,omitempty" bson:"ad_id,omitempty"`
	Type        byte          `json:"type,string" bson:"type"`
	Fullname    string        `json:"fullname" bson:"fullname"`
	Email       string        `json:"email" bson:"email"`
	Phone       string        `json:"phone" bson:"phone"`
	City        string        `json:"city" bson:"city"`
	Message     string        `json:"message" bson:"message"`
	AboutMeType byte          `json:"aboutMeType,string" bson:"about_me_type"`
	Campaign    bool          `json:"campaign" bson:"-"`
	CopyEmail   bool          `json:"copyEmail" bson:"-"`
	PriceChange bool          `json:"priceChange" bson:"price_change"`
	Ads         Ad            `json:"ad,omitempty" bson:"ads"`
	Processed   bool          `json:"processed" bson:"processed"`
	AgentType   byte          `json:"-" bson:"agent_type"`
}

type UpdatedBid struct {
	ID        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Fullname  string        `json:"fullname" bson:"fullname"`
	Email     string        `json:"email" bson:"email"`
	Phone     string        `json:"phone" bson:"phone"`
	City      string        `json:"city" bson:"city"`
	Message   string        `json:"message" bson:"message"`
	Processed bool          `json:"processed" bson:"processed"`
}
