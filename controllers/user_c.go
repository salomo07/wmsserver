package controllers
import (
	"github.com/gin-gonic/gin"	
	"wms/models"
)
func GetUser(c *gin.Context)([] models.User){
	return models.Find(map[string]interface{}{"username":"salomo07"})
}