package models

type RoleModel struct {
	Id       string       `json:"_id"`
	Rev      string       `json:"_rev"`
	Name     string       `json:"name" binding:"required"`
	Codename string       `json:"codename" binding:"required"`
	Desc     string       `json:"desc"`
	Menus    []MenuModels `json:"menus"`
	Isactive bool         `json:"isactive"`
	Table    string       `json:"table" binding:"required"`
}
