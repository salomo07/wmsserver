package controllers
import (
	"github.com/gin-gonic/gin"	
	"wms/models"
)
func GetUser(c *gin.Context)(models.User){
	// rst:=models.Query()
	models.StoreData()
	var u models.User
	return u
}