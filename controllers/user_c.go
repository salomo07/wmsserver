package controllers
import (
	"github.com/gin-gonic/gin"	
	"wms/models"
)
func GetUser(c *gin.Context)(models.Animal){
	rst:=models.Find()
	return rst
}