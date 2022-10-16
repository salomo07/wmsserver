package config

import(
	"context"
	"os"
	"github.com/joho/godotenv"
	"github.com/go-redis/redis/v9"
)
var ctx = context.Background()
func init() {
  connect()
}
func connect (){
	er := godotenv.Load()
	if er !=nil{
		panic("Fail to load .env file")
	}
	REDIS_HOST:=os.Getenv("REDIS_HOST_LOCAL")
	REDIS_PASS:=os.Getenv("REDIS_PASS")
	REDIS_PORT:=os.Getenv("REDIS_PORT")
	opt, _ := redis.ParseURL("rediss://:"+REDIS_PASS+"@"+REDIS_HOST+":"+REDIS_PORT)
  	client := redis.NewClient(opt)
  	client.Set(ctx, "foo", "bar", 0)
	val := client.Get(ctx, "foo").Val()
	print(val)
}