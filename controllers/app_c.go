package controllers
import (
	"encoding/json"
	"log"
	"time"
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"	
	"wms/models"
	"wms/config"
)
type CompanyObjt struct {
	Id string  `json:"id"`
	Ok bool  `json:"ok"`
}
func RegisterCompany(c *gin.Context)(string){
	companyData:=models.Insert(c)
	var objt CompanyObjt
	err:=json.Unmarshal([]byte(string(companyData)),&objt)
	idCompany:=""
	if err != nil{
		return `{"errorx":"`+err.Error()+`"}`
	}
	if objt.Ok {
		idCompany=objt.Id
		CreateDBCompany(idCompany)
	}
	return ""
}
func InitializingData(c *gin.Context){
	BulkDocs(c)
}
func CreateDBCompany(idCompany string)(string){
	return models.CreateDatabase("c_"+idCompany)
}
func Find(c *gin.Context)(string){
	currentTime := time.Now().Format("2006-01-02")
	now,_:=time.Parse("2006-01-02", currentTime)
	log.Println(now.Unix())
	var rst,err =config.CheckSession(c)
	if err =="" {		
		db:=c.Query("db")
		jsonData, _ := c.GetRawData()
		rst=models.Find(db,jsonData)
	}else{rst=err}
	return rst
}
func Insert(c *gin.Context)(string){
	db:=""
	if c.Query("db")==""{
		db="mastercompany"
	}else{
		db=c.Query("db")
	}
	
	jsonData, _ := c.GetRawData()
	return models.Insert(c)
}
func Update(c *gin.Context)(string){
	db:=c.Query("db")
	jsonData, _ := c.GetRawData()
	return models.Update(c)
}
func Delete(c *gin.Context)(string){
	return models.Delete(c)
}
func CreateDatabase(c *gin.Context)(string){
	db:=c.Query("db")
	return models.CreateDatabase(db)
}
func DeleteDatabase(c *gin.Context)(string){
	return models.DeleteDatabase(c)
}
func CreateUserDB(c *gin.Context)(string){
	return models.CreateUserDB(c)
}
func GetView(c *gin.Context)(string){
	return models.GetView(c)
}
func BulkDocs(c *gin.Context)(string){
	db:=c.Query("db")
	jsonData, _ := c.GetRawData()
	return models.BulkDocs(db,jsonData)
}
func GetMD5Hash(text string) string {
   hash := md5.Sum([]byte(text))
   return hex.EncodeToString(hash[:])
}