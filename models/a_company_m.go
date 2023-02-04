package models

import (
	"encoding/json"
	"wms/config"
)

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

func FindCompanyByR(jsonData []byte) string {
	companyData := config.FindDoc("mastercompany", string(jsonData))
	var comObject CompanyModels
	json.Unmarshal([]byte(companyData), &comObject)
	return companyData
}
func FindCompanyByC(basiccred string, jsonData []byte) string {
	return config.FindDocByCompany(basiccred, "mastercompany", string(jsonData))
}
