package config

import(
	"log"
	"os"
  	"strings"
  	"net/http"
  	"io/ioutil"
	"github.com/joho/godotenv"
)
var DB_STR_CON string 

func Request(method string,pathURL string,strquery string)(string){
	er := godotenv.Load()
	if er !=nil{
		panic("Fail to load .env file")
	}
	DB_STR_CON=os.Getenv("DB_STR_CON")
	payload := strings.NewReader(strquery)
	client := &http.Client {}
	req, err := http.NewRequest(method, DB_STR_CON+pathURL, payload)

	if err != nil {
		log.Println(err," xxx")
		panic(err)
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Println(err," yyy")
		panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err," zzz")
		panic(err)
	}
	return string(body)
}