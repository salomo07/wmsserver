package controllers
import (
	"log"
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"	
	"wms/models"
)
func Find(c *gin.Context)(string){
	
	log.Println(GetMD5Hash("321"))
	return models.Find(c)
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
	return models.CreateDatabase(c)
}
func GetView(c *gin.Context)(string){
	return models.GetView(c)
}
func GetMD5Hash(text string) string {
   hash := md5.Sum([]byte(text))
   return hex.EncodeToString(hash[:])
}