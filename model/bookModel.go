package model

type BookModel struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       string `json:"price"`
	Rating      string `json:"rating"`
	Discount    int    `json:"discount"`
}
