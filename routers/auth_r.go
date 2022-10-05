package routers
import (
	"github.com/gin-gonic/gin"
	"wms/controllers"
)

func AuthRouter(r *gin.Engine) {	
	r.Static("assets/", "./assets")
	
	master := r.Group("/auth")
	{
		master.GET("/getuser", func(c *gin.Context) {
			c.JSON(200, gin.H{"result":controllers.GetUser(c)})
		})
	}
}