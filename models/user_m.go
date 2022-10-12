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
func Find(db string,strquery string)(string){
	return config.Request("POST",db+"/_find",strquery)
}
func Insert(db string,strquery string)(string){
	return config.Request("POST",db,strquery)
}
func Update(db string,strquery string)(string){
	var objt UpdateObjt
	err:=json.Unmarshal([]byte(strquery),&objt)
	log.Println(err)
	if objt.Rev == "" || objt.Id==""{
		return `{"error":"bad request","reason":"_id & _rev fields must include"}`
	}
	return config.Request("PUT",db+"/"+objt.Id,strquery)
}
func Delete(db string,strquery string)(string){
	var objt UpdateObjt
	err:=json.Unmarshal([]byte(strquery),&objt)
	log.Println("XXXXXXXX",err)
	if objt.Rev == "" || objt.Id==""{
		return `{"error":"bad request","reason":"_id & _rev fields must include"}`
	}
	return config.Request("DELETE",db+"/"+objt.Id+"?rev="+objt.Rev,strquery)
}