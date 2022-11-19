package models

//Buat validator/verifikator object

type ProductModel struct {
	Id       string   `json:"_id"`
	Rev      string   `json:"_rev"`
	Sku      string   `json:"sku"`
	Type     string   `json:"type"`
	Imageurl []string `json:"imageurl"`
	Category string   `json:"category"`
	Desc     string   `json:"desc"`
	Isactive bool     `json:"isactive" binding:"required"`
	Table    string   `json:"table" binding:"required"`
}
