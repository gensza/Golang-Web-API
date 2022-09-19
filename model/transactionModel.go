package model

type Transaction struct {
	Id             int       `json:"id"`
	Noref          string    `json:"no_ref"`
	IdBook         int       `json:"id_book"`
	Qty            int       `json:"qty"`
	Transaction_at string    `json:"transaction_at"`
	Courier        []Courier `json:"courier"`
}

type Courier struct {
	IdCourier int    `json:"id"`
	IdUser    int    `json:"id_user"`
	Name      string `json:"name"`
}
