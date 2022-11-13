package config

import (
	"context"
	"encoding/json"
	"log"
	"os"
	_ "time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v9"
	"github.com/joho/godotenv"
)

type ContactStruct struct {
	Mobile string `json:"mobile"`
	Email  string `json:"email"`
}
type UserStruct struct {
	Id        string        `json:"_id"`
	Rev       string        `json:"_rev"`
	Username  string        `json:"username"`
	Password  string        `json:"password"`
	Idcompany string        `json:"idcompany"`
	Contact   ContactStruct `json:"contact"`
}

var ctx = context.Background()

type CookieReq struct {
	User    UserStruct `json:"user"`
	Expired int64      `json:"expired"`
}

var REDIS_SERVER = ""
var ERROR_LOAD_ENV = ""

func init() {
	er := godotenv.Load()
	if er != nil {
		ERROR_LOAD_ENV = er.Error()
		log.Println("Error : Fail to load .env file")
	}
	log.Println(".env is loaded")
	REDIS_SERVER = os.Getenv("REDIS_HOST_LOCAL")
}
func CheckSession(c *gin.Context) (string, string) {
	// var rst =""
	tokenIn, _ := c.Cookie("token")
	if tokenIn == "" {
		return "", string(`{"error":"Cookie is not found"}`)
	} else {
		dataSession := GetData(tokenIn)
		if dataSession == "" {
			return "", string(`{"error":"Session not found"}`)
		} else {
			var objt CookieReq
			json.Unmarshal([]byte(dataSession), &objt)
			log.Println(objt)

			// currentTime := time.Now().Format("2006-01-02")
			// now,_:=time.Parse("2006-01-02", currentTime)
			// if objt.Expired.Unix()<=now.Unix(){
			// 	log.Println(`{"error":"Session is expired"}`)
			// 	return false
			// }else{
			// 	return true
			// }
			xxx, _ := json.Marshal(CookieReq{})
			return string(xxx), ""
		}
	}
}
func SetData(key string, data string) string {
	if ERROR_LOAD_ENV != "" {
		return `{"error":"` + ERROR_LOAD_ENV + `"}`
	}

	// rdb := redis.NewClient(&redis.Options{Addr:REDIS_SERVER,Password:"passredis",DB:0})
	rdb := redis.NewClient(&redis.Options{Addr: REDIS_SERVER, Password: "passredis", DB: 0})
	err := rdb.Set(ctx, key, data, 0).Err()
	if err != nil {
		return `{"error":"` + err.Error() + `"}`
	}
	return `{"ok":true}`
}
func GetData(key string) string {
	if ERROR_LOAD_ENV != "" {
		return `{"error":"` + ERROR_LOAD_ENV + `"}`
	}
	rdb := redis.NewClient(&redis.Options{Addr: REDIS_SERVER, Password: "passredis", DB: 0})
	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		return `{"error":"` + err.Error() + `"}`
	}
	return val
}
