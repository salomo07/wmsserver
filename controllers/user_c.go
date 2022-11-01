package controllers
import (
	// "encoding/json"
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
	log.Println("ssss")
	companyData:=Insert(c)
	log.Println(companyData)
	// var objt CompanyObjt
	// err:=json.Unmarshal([]byte(string(companyData)),&objt)
	// if err != nil{
	// 	return `{"error":"`+err.Error()+`"}`
	// }
	// if objt.Ok {
	// 	idCompany:=objt.Id
	// 	log.Println("company_"+idCompany)
	// 	// log.Println(models.CreateDatabase("company_"+idCompany))
	// }
	return ""
}
func Find(c *gin.Context)(string){
	currentTime := time.Now().Format("2006-01-02")
	now,_:=time.Parse("2006-01-02", currentTime)
	log.Println(now.Unix())
	var rst,err =config.CheckSession(c)
	if err =="" {
		rst=models.Find(c)
	}else{rst=err}
	return rst
}
func Insert(c *gin.Context)(string){
	return models.Insert(c)
}
func Update(c *gin.Context)(string){
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
	return models.BulkDocs(c)
}
func GetMD5Hash(text string) string {
   hash := md5.Sum([]byte(text))
   return hex.EncodeToString(hash[:])
}