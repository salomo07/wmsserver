package config

import(
	"os"
  	"strings"
  	"net/http"
  	"io/ioutil"
	"github.com/joho/godotenv"
)

func Request(method string,pathURL string,strquery string)(string){
	er := godotenv.Load()
	if er !=nil{
		panic("Fail to load .env file")
	}
	DB_HOST:=os.Getenv("DB_HOST")
	payload := strings.NewReader(strquery)
	client := &http.Client {}
	req, err := http.NewRequest(method, DB_HOST+pathURL, payload)

	if err != nil {
		panic(err)
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	return string(body)
}
func RequestWithoutPath(method string,strquery string)(string){
	er := godotenv.Load()
	if er !=nil{
		panic("Fail to load .env file")
	}
	DB_HOST:=os.Getenv("DB_HOST")
	payload := strings.NewReader(strquery)
	client := &http.Client {}
	req, err := http.NewRequest(method, DB_HOST, payload)

	if err != nil {
		panic(err)
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	return string(body)
}