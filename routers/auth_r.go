package routers
import (
	"github.com/gin-gonic/gin"
	"log"
	"wms/controllers"
)

func AuthRouter(r *gin.Engine) {	
	r.Static("assets/", "./assets")
	
	master := r.Group("/auth")
	{
		master.POST("/getuser", func(c *gin.Context) {
			usr:=controllers.GetUser(c)
			log.Println(usr)
		})
	}
}