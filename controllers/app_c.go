package controllers
import (
	"strings"
	"encoding/json"
	"strconv"
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
type SetObjt struct {
	Data string  `json:"data" binding:"required"`
}
func RegisterCompany(c *gin.Context)(string){
	jsonData, _ := c.GetRawData()
	companyData:=models.Insert("mastercompany",jsonData)
	
	var objt CompanyObjt
	err:=json.Unmarshal([]byte(string(companyData)),&objt)
	
	idCompany:=""
	if err != nil{
		log.Println(`{"error":"`+err.Error()+`"}`)
		return `{"error":"`+err.Error()+`"}`
	}

	if objt.Ok {
		idCompany=objt.Id
		CreateDBCompany(idCompany)
		currentTime := time.Now().Format("2006-01-02")
		now,_:=time.Parse("2006-01-02", currentTime)
		ss := strconv.FormatInt(now.Unix(), 10)
		log.Println("admin"+string(idCompany))
		config.SetData("admin"+string(idCompany),ss)
		config.SetData("member"+string(idCompany),ss+"m")
		log.Println("admin"+string(idCompany),ss)
		log.Println("admin"+string(idCompany),ss+"m")
	}
	return ""
}
func SetRedis(c *gin.Context)(string){
	jsonData, _ := c.GetRawData()
	var objt SetObjt
	err:=json.Unmarshal([]byte(string(jsonData)),&objt)
	if err ==nil{
		return models.SetRedis(c.Query("key"),objt.Data)
	}
	return `{"error":`+err.Error()+`}`
}
func CheckSession(c *gin.Context)(string){
	apikey:=c.GetHeader("Authorization")
	if apikey==""{
		return `{"error":"You must have api key"}`
	}
	reqToken := apikey
	
	splitToken := strings.Split(reqToken, "Bearer ")
	reqToken = splitToken[1]
	log.Println(reqToken,len(splitToken),"dfdfdfdf")
	
	return reqToken
}
func GetRedis(c *gin.Context)(string){
	return models.GetRedis(c.Query("key"))
}
func InitializingData(idCompany string){
	// masterwarehouse:=`{"name":"Gudang A","code":"GA","stokmin":3000,"stokmax":10000}`
	// models.BulkDocs("c"+idCompany,jsonData)
}
func CreateDBCompany(idCompany string)(string){
	return models.CreateDatabase("c_"+idCompany)
}
func Find(c *gin.Context)(string){
	var rst,err =config.CheckSession(c)
	if err =="" {		
		db:=c.Query("db")
		jsonData, _ := c.GetRawData()
		rst=models.Find(db,jsonData)
	}else{rst=err}
	return rst
}
func Insert(c *gin.Context)(string){
	db:=c.Query("db")
	if db==""{
		return `{"error":"Please insert db variable"}`
	}
	jsonData, _ := c.GetRawData()
	return models.Insert(c.Query("db"),jsonData)
}
func Update(c *gin.Context)(string){
	db:=c.Query("db")
	if db==""{
		return `{"error":"Please insert db variable"}`
	}
	jsonData, _ := c.GetRawData()
	return models.Update(db,jsonData)
}
func Delete(c *gin.Context)(string){
	return models.Delete(c)
}
func CreateDatabase(c *gin.Context)(string){
	db:=c.Query("db")
	if db==""{
		return `{"error":"Please insert db variable"}`
	}
	return models.CreateDatabase(db)
}
func DeleteDatabase(c *gin.Context)(string){
	db:=c.Query("db")
	if db==""{
		return `{"error":"Please insert db variable"}`
	}
	return models.DeleteDatabase(db)
}
func CreateUserDB(c *gin.Context)(string){
	db:=c.Query("db")
	if db==""{
		return `{"error":"Please insert db variable"}`
	}
	jsonData, _ := c.GetRawData()
	return models.CreateUserDB(jsonData)
}
func GetView(c *gin.Context)(string){
	if c.Query("db")=="" || c.Query("ddoc")=="" || c.Query("view")==""{
		return `{"error":"db, ddoc, view parameters must included"}`
	}	
	jsonData, _ := c.GetRawData()
	return models.GetView(c.Query("db"),c.Query("ddoc"),c.Query("view"),jsonData)
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