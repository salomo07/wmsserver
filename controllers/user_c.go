package controllers
import (
	"log"
	"time"
	"strings"
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"	
	"wms/models"
	"wms/config"
)
func Find(c *gin.Context)(string){
	currentTime := time.Now().Format("2006-01-02")
	now,_:=time.Parse("2006-01-02", currentTime)
	log.Println(now.Unix())
	var rst =config.CheckSession(c)
	if !strings.Contains(rst, "error") {
		rst=models.Find(c)
	}
	return rst
}
func Insert(c *gin.Context)(string){
	config.SetData("session-01a7d17fac6bae63fa5c69b76e019862",`{
	  "expired": 1666051200,
	  "user": "202cb962ac59075b964b07152d234b70"
	}`)
	return ""
	// models.Insert(c)
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