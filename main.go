package main
import (
	"github.com/gin-gonic/gin"
	_"log"
  	"wms/routers"
)

func main() {
	r := gin.Default()
	routers.CouchDBRouter(r)

	r.NoRoute(func(c *gin.Context) {
	    c.JSON(404, gin.H{"code": "404", "message": "Page not found"})
	})
	// port := os.Getenv("PORT")
	r.Run(":7777")
}
