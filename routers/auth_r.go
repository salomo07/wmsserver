package routers
import (
	"github.com/gin-gonic/gin"
	"wms/controllers"
)

func AuthRouter(r *gin.Engine) {	
	r.Static("assets/", "./assets")
	
	master := r.Group("/auth")
	{
		master.POST("/getuser", func(c *gin.Context) {
			c.Header("Content-Type", "application/json; charset=utf-8")
			c.String(200,controllers.FindUser(c))
		})
		master.POST("/insertuser", func(c *gin.Context) {
			c.Header("Content-Type", "application/json; charset=utf-8")
			c.String(200,controllers.InsertUser(c))
		})
		master.POST("/updateuser", func(c *gin.Context) {
			c.Header("Content-Type", "application/json; charset=utf-8")
			c.String(200,controllers.UpdateUser(c))
		})
		master.POST("/deleteuser", func(c *gin.Context) {
			c.Header("Content-Type", "application/json; charset=utf-8")
			c.String(200,controllers.DeleteUser(c))
		})
	}
}