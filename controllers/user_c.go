package controllers
import (
	"github.com/gin-gonic/gin"	
	"wms/models"
)
func GetUser(c *gin.Context)(models.User){
	var user models.User	
	return user
}