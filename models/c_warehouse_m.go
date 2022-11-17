package models

//Buat validator/verifikator object

type WarehouseModel struct {
	Id       string `json:"_id"`
	Rev      string `json:"_rev"`
	Name     string `json:"name"`
	Codename string `json:"codename"`
	Menus    string `json:"menus"`
	Role     string `json:"role"`
	Table    string `json:"table"`
	Desc     string `json:"desc"`
}
type ColumnModel struct {
	Id       string `json:"_id"`
	Rev      string `json:"_rev"`
	Name     string `json:"name"`
	Codename string `json:"codename"`
	Rows     string `json:"rows"`
	Role     string `json:"role"`
	Table    string `json:"table"`
	Isactive bool   `json:"isactive" binding:"required"`
}
