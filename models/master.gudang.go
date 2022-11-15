package models

// "wms/config"
// "github.com/gin-gonic/gin"
// "encoding/json"

//Buat validator/verifikator object
type GudangObjt struct {
	Id       string             `json:"_id"`  //required
	Rev      string             `json:"_rev"` //required
	Name     string             `json:"name"`
	Codename string             `json:"code"`
	StockMin float32            `json:"stockmin"`
	StockMax float32            `json:"stockmax"`
	Layout   string             `json:"layout"`
	Detail   []DetailGudangObjt `json:"details"`
	Table    string             `json:"table"`
	IsActive bool               `json:"isactive"`
}

type DetailGudangObjt struct {
	Name   string `json:"name"`   //required
	Coords string `json:"coords"` //required
	Href   string `json:"href"`
}
