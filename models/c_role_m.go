package models

type RoleModel struct {
	Id       string     `json:"_id"`
	Rev      string     `json:"_rev"`
	Name     string     `json:"name" binding:"required"`
	Codename string     `json:"codename" binding:"required"`
	Menus    MenuModels `json:"menus"`
	Role     string     `json:"role"`
	Table    string     `json:"table"`
	Desc     string     `json:"desc"`
}
