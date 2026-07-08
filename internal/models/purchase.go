package models

import "time"

type Purchase struct {
	ID int
	CustomerId int
	Total float64
	PurchaseDate time.Time
}

type PurchaseItem struct {
	ID int
	PurchaseId int
	ProductId int
	Quantity int
	Price float64
}
