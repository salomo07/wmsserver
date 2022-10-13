package models
import (
	"log"
	"wms/config"
	"github.com/gin-gonic/gin"
	"encoding/json"
)
//Buat validator/verifikator object
type UpdateObjt struct {
	Id string  `json:"_id"`
	Rev string `json:"_rev"`
}
func Find(c *gin.Context)(string){
	db:=c.Query("db")
	jsonData, _ := c.GetRawData()
	return config.Find(db+"/_find",string(jsonData))
}
func Insert(c *gin.Context)(string){
	db:=c.Query("db")
	jsonData, _ := c.GetRawData()
	return config.Insert(db,string(jsonData))
}
func Update(c *gin.Context)(string){
	db:=c.Query("db")
	jsonData, _ := c.GetRawData()

	var objt UpdateObjt
	err:=json.Unmarshal([]byte(string(jsonData)),&objt)
	if err != nil{
		return `{"error":"`+err.Error()+`"}`
	}
	return config.Update(db+"/"+objt.Id,string(jsonData))
}
func Delete(c *gin.Context)(string){
	db:=c.Query("db")
	jsonData, _ := c.GetRawData()
	var objt UpdateObjt
	
	err:=json.Unmarshal([]byte(string(jsonData)),&objt)
	log.Println(err,db+"/"+objt.Id+"?rev="+objt.Rev)
	if err != nil{
		return `{"error":"`+err.Error()+`"}`
	}
	return config.Delete(db+"/"+objt.Id+"?rev="+objt.Rev)
}
func CreateDatabase(c *gin.Context)(string){
	db:=c.Query("db")
	return config.CreateDatabase(db)
}
func DeleteDatabase(c *gin.Context)(string){
	db:=c.Query("db")
	return config.DeleteDatabase(db)
}