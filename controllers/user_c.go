package controllers
import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"	
	"wms/models"
)
func FindUser(c *gin.Context)(string){
	db:=c.Query("db")
	jsonData, _ := c.GetRawData()
	return models.Find(db,string(jsonData))
}
func InsertUser(c *gin.Context)(string){
	db:=c.Query("db")
	jsonData, _ := c.GetRawData()
	return models.Insert(db,string(jsonData))
}
func UpdateUser(c *gin.Context)(string){
	db:=c.Query("db")
	jsonData, _ := c.GetRawData()
	return models.Update(db,string(jsonData))
}
func DeleteUser(c *gin.Context)(string){
	db:=c.Query("db")
	jsonData, _ := c.GetRawData()
	return models.Delete(db,string(jsonData))
}
func GetMD5Hash(text string) string {
   hash := md5.Sum([]byte(text))
   return hex.EncodeToString(hash[:])
}