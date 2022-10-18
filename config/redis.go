package config

import(
	_"time"
	"log"
	"context"
	"os"
	"encoding/json"
	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v9"
)
type ContactStruct struct {
	Mobile string  `json:"mobile"`
	Email string `json:"email"`
}
type UserStruct struct {
	Id string  `json:"_id"`
	Rev string `json:"_rev"`
	Username string `json:"username"`
	Password string `json:"password"`
	Idcompany string `json:"idcompany"`
	Contact ContactStruct `json:"contact"`
}

var ctx = context.Background()
type CookieReq struct {
	User UserStruct `json:"user"`
	Expired int64 `json:"expired"`
}
func CheckSession(c *gin.Context)string{
	var rst =""
	tokenIn,_:=c.Cookie("token")
	if tokenIn ==""{
		rst =string(`{"error":"Cookie is not found"}`)
	}else {
		dataSession:=GetData(tokenIn)
		if dataSession==""{
			rst =string(`{"error":"Session not found"}`)
		}else{
			var objt CookieReq
			json.Unmarshal([]byte(dataSession),&objt)
			log.Println(objt)
			
			// currentTime := time.Now().Format("2006-01-02")
			// now,_:=time.Parse("2006-01-02", currentTime)
			// if objt.Expired.Unix()<=now.Unix(){
			// 	log.Println(`{"error":"Session is expired"}`)
			// 	return false
			// }else{
			// 	return true
			// }
			log.Println("xxxxxxxxx")
			xxx,_:=json.Marshal(CookieReq{})
			rst = string(xxx)
		}
		
	}
	return rst
}
func SetData (key string,data string){
	er := godotenv.Load()
	if er !=nil{
		panic("Fail to load .env file")
	}
	REDIS_STRING:=os.Getenv("REDIS_STRING_LOCAL")
	opt, _ := redis.ParseURL(REDIS_STRING)
  	client := redis.NewClient(opt)
  	client.Set(ctx, key, data,0)
  	log.Println(key," saved")
}
func GetData (key string)string{
	er := godotenv.Load()
	if er !=nil{
		panic("Fail to load .env file")
	}
	REDIS_STRING:=os.Getenv("REDIS_STRING_LOCAL")
	opt, _ := redis.ParseURL(REDIS_STRING)
  	client := redis.NewClient(opt)
  	val := client.Get(ctx, key).Val()
  	log.Println(val)
  	return val
}