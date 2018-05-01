package models

type AdFilter struct {
	Location    string 	`json:"location"`
	MinPrice    float32 `json:"minPrice"`
	MaxPrice    float32	`json:"maxPrice"`
	Type 		byte	`json:"type"`
	Rooms 		byte	`json:"rooms"`
	Text		string	`json:"text"`
}

type AdFilterRequest struct {
	Filter   AdFilter `json:"filter"`
	Paginate PaginateFiler `json:"paginate"`
}

type PaginateFiler struct {
	PerPage int  `json:"perPage"`
	Page    int  `json:"page"`
}