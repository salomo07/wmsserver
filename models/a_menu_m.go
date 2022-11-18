package models

type MenuModels struct {
	Id       string         `json:"_id"`
	Rev      string         `json:"_rev"`
	Name     string         `json:"name"`
	Codename string         `json:"codename"`
	Icon     string         `json:"icon"`
	Url      string         `json:"url"`
	Desc     string         `json:"desc"`
	Submenus []SubmenuModel `json:"submenus"`
	Isactive bool           `json:"isactive"`
	Table    string         `json:"table" binding:"required"`
}
type SubmenuModel struct {
	Id       string `json:"_id"`
	Rev      string `json:"_rev"`
	Name     string `json:"name"`
	Codename string `json:"codename"`
	Icon     string `json:"icon"`
	Url      string `json:"url"`
	Desc     string `json:"desc"`
	Isactive bool   `json:"isactive"`
}
