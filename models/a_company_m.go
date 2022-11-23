package models

type CompanyModels struct {
	Id      string         `json:"_id"`
	Rev     string         `json:"_rev"`
	Name    string         `json:"name"`
	Level   string         `json:"level"`
	Expired string         `json:"expired"`
	Contact ContactCompany `json:"contact"`
	Address string         `json:"address"`
	Users   []string       `json:"users"`
	Shipto  string         `json:"shipto"`
	Table   string         `json:"table" binding:"required"`
}
type ContactCompany struct {
	Email  string `json:"email"`
	Mobile string `json:"mobile" binding:"required"`
}
