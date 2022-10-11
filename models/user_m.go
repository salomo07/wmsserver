package models
import (
	"encoding/json"
	"wms/config"
)
//Buat validator/verifikator object
type UpdateObjt struct {
	Rev string `json:"_rev"`
}
func Find(strquery string)(string){
	return config.Request("POST","wms/_find",strquery)
}
func Insert(strquery string)(string){
	return config.Request("POST","wms",strquery)
}
func Update(strquery string)(string){
	var objt UpdateObjt
	json.Unmarshal(strquery,&objt)
	if objt.Rev == ""{
		return `{"error":"_rev field must include"}`
	}
	return config.Request("PUT","wms",strquery)
}