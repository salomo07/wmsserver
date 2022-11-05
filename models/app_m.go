package models
import (
	"wms/config"
	"github.com/gin-gonic/gin"
	"encoding/json"
)
//Buat validator/verifikator object
type UpdateObjt struct {
	Id string  `json:"_id"`
	Rev string `json:"_rev"`
}
type UserDBObjt struct {
	Name string  `json:"name"`
	Password string `json:"password"`
}

func Find(db string,jsonData []byte)(string){
	return config.FindDoc(db+"/_find",string(jsonData))
}
func Insert(db string,jsonData []byte)(string){
	return config.InsertDoc(db,string(jsonData))
}
func Update(db string,jsonData []byte)(string){
	var objt UpdateObjt
	err:=json.Unmarshal([]byte(string(jsonData)),&objt)
	if err != nil{
		return `{"error":"`+err.Error()+`"}`
	}
	return config.UpdateDoc(db+"/"+objt.Id,string(jsonData))
}
func Delete(c *gin.Context)(string){
	db:=c.Query("db")
	jsonData, _ := c.GetRawData()
	var objt UpdateObjt
	
	err:=json.Unmarshal([]byte(string(jsonData)),&objt)
	if err != nil{
		return `{"error":"`+err.Error()+`"}`
	}
	return config.DeleteDoc(db+"/"+objt.Id+"?rev="+objt.Rev)
}
func CreateDatabase(db string)(string){
	return config.CreateDatabase(db)
}
func DeleteDatabase(db string)(string){
	return config.DeleteDatabase(db)
}
func GetView(db string,ddoc string,view string, jsonData []byte)(string){ 
	return config.GetDataByView(db,ddoc,view,string(jsonData))
}
func BulkDocs(db string,jsonData []byte)(string){
	return config.BulkDocs(db,string(jsonData))
}
func CreateUserDB(jsonData []byte)(string){
	var objt UserDBObjt
	err:=json.Unmarshal([]byte(string(jsonData)),&objt)
	if err != nil{
		return `{"error":"`+err.Error()+`"}`
	}
	if objt.Name=="" || objt.Password==""{
		return `{"error":"Please insert name & password"}`
	}
	return config.CreateUserDB(objt.Name,string(jsonData))
}
func CreateReplication(c *gin.Context)(string){
	jsonData, _ := c.GetRawData()
	return config.CreateReplication(string(jsonData))
}
func SetRedis(key string,val string)(string){
	return config.SetData (key,val)
}
func GetRedis(key string)(string){
	return config.GetData (key)
}