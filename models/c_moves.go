package models

//Buat validator/verifikator object

type MovesModel struct {
	Id        string `json:"_id"`
	Rev       string `json:"_rev"`
	IdProduct string `json:"idproduct"`
	Datetime  string `json:"datetime"`
	Fromrack  string `json:"fromrack"`
	Torack    string `json:"torack"`
	Moveby    string `json:"moveby"`
	Approveby string `json:"approveby"`
	Isactive  bool   `json:"isactive" binding:"required"`
	Table     string `json:"table" binding:"required"`
}
