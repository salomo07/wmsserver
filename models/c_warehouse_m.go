package models

//Buat validator/verifikator object

type WarehouseModel struct {
	Id        string        `json:"_id"`
	Rev       string        `json:"_rev"`
	Name      string        `json:"name"`
	Codename  string        `json:"codename"`
	Minstock  int32         `json:"minstock"`
	Maxstock  int32         `json:"maxstock"`
	Urllayout string        `json:"urllayout"`
	Details   []ColumnModel `json:"details"`
	Desc      string        `json:"desc"`
	Table     string        `json:"table"`
	Isactive  bool          `json:"isactive" binding:"required"`
}
type ColumnModel struct {
	Id       string   `json:"_id"`
	Rev      string   `json:"_rev"`
	Name     string   `json:"name"`
	Codename string   `json:"codename"`
	Rows     []string `json:"rows"`
	Role     string   `json:"role"`
	Table    string   `json:"table"`
	Isactive bool     `json:"isactive" binding:"required"`
}
type RowModel struct {
	Id       string      `json:"_id"`
	Rev      string      `json:"_rev"`
	Name     string      `json:"name"`
	Codename string      `json:"codename"`
	Alt      string      `json:"alt"`
	Coords   string      `json:"coords"`
	Href     string      `json:"href"`
	Shape    string      `json:"shape"`
	Target   string      `json:"target"`
	Racks    []RackModel `json:"rackmodel"`
	Isactive bool        `json:"isactive" binding:"required"`
}
type RackModel struct {
	Id            string               `json:"_id"`
	Rev           string               `json:"_rev"`
	Name          string               `json:"name"`
	Codename      string               `json:"codename"`
	ProductDetail []ProductDetailModel `json:"productdetail"`
}

type ProductDetailModel struct {
	Id        string `json:"_id"`
	Rev       string `json:"_rev"`
	Idproduct string `json:"idproduct"`
	Qty       string `json:"codename"`
	ProdDate  string `json:"prodDate"`
	ExpDate   string `json:"expDate"`
}
