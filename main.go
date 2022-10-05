package main
import (
	"github.com/gin-gonic/gin"
	_"log"
  	"wms/routers"
)

func main() {
	r := gin.Default()
	routers.AuthRouter(r)

	// port := os.Getenv("PORT")
	r.Run(":7777")
}
