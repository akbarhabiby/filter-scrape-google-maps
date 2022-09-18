package models

import "time"

type Scrapes struct {
	ID          int       `json:"-"`
	FullAddress string    `json:"full_address"`
	Number      string    `json:"number"`
	District    string    `json:"district"`
	City        string    `json:"city"`
	Province    string    `json:"province"`
	PostalCode  string    `json:"postal_code"`
	Country     string    `json:"country"`
	Latitude    string    `json:"latitude"`
	Longitude   string    `json:"longitude"`
	PlusCode    string    `json:"plus_code"`
	CreatedAt   time.Time `json:"-"`
}

type Result struct {
	Scrapes []Scrapes `json:"scrapes"`
}
