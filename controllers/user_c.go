package controllers
import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"	
	"wms/models"
)
func FindUser(c *gin.Context)(string){
	jsonData, _ := c.GetRawData()
	return models.Find(string(jsonData))
}
func InsertUser(c *gin.Context)(string){
	jsonData, _ := c.GetRawData()
	return models.Insert(string(jsonData))
}
func Login(c *gin.Context)(string){
	jsonData, _ := c.GetRawData()
	return models.Find(string(jsonData))
}
func GetMD5Hash(text string) string {
   hash := md5.Sum([]byte(text))
   return hex.EncodeToString(hash[:])
}