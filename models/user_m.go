package models
import (
	"log"
	"encoding/json"
	"wms/config"
)
//Buat validator/verifikator object
type UpdateObjt struct {
	Id string `json:"_id"`
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
	
	err:=json.Unmarshal([]byte(strquery),&objt)
	log.Println(strquery,"wms/"+objt.Id,err)

	if objt.Rev == "" || objt.Id==""{
		return `{"error":"_id & _rev fields must include"}`
	}
	return config.Request("PUT","wms/"+objt.Id,strquery)
}