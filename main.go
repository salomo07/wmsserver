package main
import (
	"os"
	"github.com/gin-gonic/gin"
	_"log"
  	"wms/routers"
)

func main() {
	r := gin.Default()
	// routers.ApiRouter(r)
	routers.AuthRouter(r)
	// routers.MasterRouter(r)
	// routers.AdministratorRouter(r)

	port := os.Getenv("PORT")
	r.Run(":"+port)
}
