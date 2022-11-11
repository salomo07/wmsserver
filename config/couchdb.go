package config

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

var DB_STR_CON_DEFAULT string
var DB_BASE_URL string

func init() {
	er := godotenv.Load()
	if er != nil {
		panic("Fail to load .env file")
	}
	DB_STR_CON_DEFAULT = os.Getenv("DB_STR_CON")
	DB_BASE_URL = os.Getenv("HOST_DB")
}

func BulkDocs(db string, strquery string) string {
	return Request("POST", db+"/_bulk_docs", strquery)
}
func FindDoc(db string, strquery string) string {
	return Request("POST", db+"/_find", strquery)
}
func FindDoc2(userdb string, passdb string, db string, strquery string) string {
	return Request2(userdb, passdb, "POST", db+"/_find", strquery)
}
func InsertDoc(path string, strquery string) string {
	return Request("POST", path, strquery)
}
func UpdateDoc(path string, strquery string) string {
	return Request("PUT", path, strquery)
}
func DeleteDoc(path string) string {
	return Request("DELETE", path, "")
}
func CreateDatabase(db string) string {
	return Request("PUT", db, "")
}
func CreateDatabase2(userdb string, passdb string, db string) string {
	return Request2(userdb, passdb, "PUT", db, "")
}
func DeleteDatabase(db string) string {
	return Request("DELETE", db, "")
}
func CreateUserDB(username string, strquery string) string {
	return Request("PUT", "_users/org.couchdb.user:"+username, strquery)
}
func InsertAuthorDB(db string, strquery string) string {
	return Request("PUT", db+"/_security", strquery)
}
func GetAuthorDB(db string, strquery string) string {
	return Request("GET", db+"/_security", strquery)
}
func CreateReplication(strquery string) string {
	return Request("POST", "_replicate", strquery)
}
func GetDataByView(db string, dsgname string, viewname string, str string) string {
	return Request("POST", db+"/_design/"+dsgname+"/_view/"+viewname, str)
}
func Request(method string, pathURL string, strquery string) string {

	payload := strings.NewReader(strquery)
	client := &http.Client{}
	req, errCon := http.NewRequest(method, DB_STR_CON_DEFAULT+pathURL, payload)

	if errCon != nil {
		log.Println(errCon, " xxx")
		return string(`{"error":` + errCon.Error() + `}`)
	}
	req.Header.Add("Content-Type", "application/json")

	res, errDo := client.Do(req)
	if errDo != nil {
		log.Println(errDo, " yyy")
		return string(`{"error":` + errDo.Error() + `}`)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err, " zzz")
		return string(`{"error":` + err.Error() + `}`)
	}
	return string(body)
}
func Request2(userdb string, passdb string, method string, pathURL string, strquery string) string {
	DB_URL := "https://" + userdb + ":" + passdb + "@" + DB_BASE_URL + "/"
	log.Println(DB_URL)
	payload := strings.NewReader(strquery)
	client := &http.Client{}
	req, errCon := http.NewRequest(method, DB_URL+pathURL, payload)

	if errCon != nil {
		log.Println(errCon, " xxx")
		return string(`{"error":` + errCon.Error() + `}`)
	}
	req.Header.Add("Content-Type", "application/json")

	res, errDo := client.Do(req)
	if errDo != nil {
		log.Println(errDo, " yyy")
		return string(`{"error":` + errDo.Error() + `}`)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err, " zzz")
		return string(`{"error":` + err.Error() + `}`)
	}
	return string(body)
}
