package models

const (
	LowPrice = iota + 1
	HighPrice
	Popular
	Best
)

type AdFilter struct {
	Location  string  `json:"location"`
	MinPrice  float32 `json:"minPrice,string"`
	MaxPrice  float32 `json:"maxPrice,string"`
	Type      byte    `json:"type,string"`
	Rooms     byte    `json:"rooms,string"`
	Text      string  `json:"text"`
	AgentType byte    `json:"agentType,string"`
	Sort      int 	  `json:"sort"`
}

type AdFilterRequest struct {
	Filter   AdFilter      `json:"filter"`
	Paginate PaginateFiler `json:"paginate"`
}

type BidFilterRequest struct {
	AgentType byte          `json:"agent_type"`
	Paginate  PaginateFiler `json:"paginate"`
}

type PaginateFiler struct {
	PerPage int `json:"perPage"`
	Page    int `json:"page"`
}
