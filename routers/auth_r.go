package routers
import (
	"github.com/gin-gonic/gin"
	"wms/controllers"
)

func CouchDBRouter(r *gin.Engine) {	
	r.Static("assets/", "./assets")
	
	master := r.Group("/couch")
	{
		master.POST("/getdoc", func(c *gin.Context) {
			c.Header("Content-Type", "application/json; charset=utf-8")
			c.String(200,controllers.Find(c))
		})
		master.POST("/insertdoc", func(c *gin.Context) {
			c.Header("Content-Type", "application/json; charset=utf-8")
			c.String(200,controllers.Insert(c))
		})
		master.POST("/updatedoc", func(c *gin.Context) {
			c.Header("Content-Type", "application/json; charset=utf-8")
			c.String(200,controllers.Update(c))
		})
		master.POST("/deletedoc", func(c *gin.Context) {
			c.Header("Content-Type", "application/json; charset=utf-8")
			c.String(200,controllers.Delete(c))
		})
		master.POST("/createdb", func(c *gin.Context) {
			c.Header("Content-Type", "application/json; charset=utf-8")
			c.String(200,controllers.CreateDatabase(c))
		})
	}
}