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

func Find(c *gin.Context)(string){
	db:=c.Query("db")
	jsonData, _ := c.GetRawData()
	return config.FindDoc(db+"/_find",string(jsonData))
}
func Insert(c *gin.Context)(string){
	db:=c.Query("db")
	jsonData, _ := c.GetRawData()
	return config.InsertDoc(db,string(jsonData))
}
func Update(c *gin.Context)(string){
	db:=c.Query("db")
	jsonData, _ := c.GetRawData()

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
func DeleteDatabase(c *gin.Context)(string){
	db:=c.Query("db")
	return config.DeleteDatabase(db)
}
func GetView(c *gin.Context)(string){
	jsonData, _ := c.GetRawData() 
	db:=c.Query("db")
	ddoc:=c.Query("ddoc")
	view:=c.Query("view")
	if db=="" || ddoc=="" || view==""{
		return `{"error":"db, ddoc, view parameters must included"}`
	}else{
		return config.GetDataByView(db,ddoc,view,string(jsonData))
	}
}
func BulkDocs(c *gin.Context)(string){
	db:=c.Query("db")
	jsonData, _ := c.GetRawData()
	return config.BulkDocs(db,string(jsonData))
}
func CreateUserDB(c *gin.Context)(string){
	jsonData, _ := c.GetRawData()
	var objt UserDBObjt
	err:=json.Unmarshal([]byte(string(jsonData)),&objt)
	if err != nil{
		return `{"error":"`+err.Error()+`"}`
	}
	return config.CreateUserDB(objt.Name,string(jsonData))
}
