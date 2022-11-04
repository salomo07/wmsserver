package models
import (
	// "wms/config"
	// "github.com/gin-gonic/gin"
	// "encoding/json"
)
//Buat validator/verifikator object
type GudangObjt struct {
	Id string  `json:"_id"`  //required
	Rev string `json:"_rev"` //required
	Name string `json:"name"`
	Code string `json:"code"`
	StockMin number `json:"stockmin"`
	StockMax number `json:"stockmax"`
	Layout string `json:"layout"`
	Detail []DetailGudangObjt `json:"details"`
	Table string `json:"table"`
	IsActive string `json:"isactive"`
}

type DetailGudangObjt struct {
	Name string  `json:"name"`  //required
	Coords string `json:"coords"` //required
	Href string `json:"href"`
}