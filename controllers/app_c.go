package controllers

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"log"
	"strconv"
	"strings"
	"time"
	"wms/config"
	"wms/models"

	"github.com/gin-gonic/gin"
)

type CompanyObjt struct {
	Id string `json:"id"`
	Ok bool   `json:"ok"`
}
type UserDBObjt struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Roles    string `json:"roles"`
}
type AdminsRole struct {
	Names string `json:"names"`
	Roles string `json:"roles"`
}
type AuthorDBObjt struct {
	Admins  AdminsRole `json:"admins"`
	Members AdminsRole `json:"members"`
}
type SetObjt struct {
	Data string `json:"data" binding:"required"`
}

func RegisterCompany(c *gin.Context) string {
	jsonData, _ := c.GetRawData()
	companyData := models.InsertDoc("mastercompany", jsonData)

	var objt CompanyObjt
	err := json.Unmarshal([]byte(string(companyData)), &objt)

	idCompany := ""
	if err != nil {
		log.Println(`{"error":"` + err.Error() + `"}`)
		return `{"error":"` + err.Error() + `"}`
	}

	if objt.Ok {
		idCompany = objt.Id
		currentTime := time.Now().Format("2006-01-02")
		now, _ := time.Parse("2006-01-02", currentTime)
		password := strconv.FormatInt(now.Unix(), 10)
		CreateDBCompany(idCompany)

		authJson := `{"admins": { "names": ["admin` + idCompany + `"]}, "members": { "names": ["jan"]}}`
		userJson := `{"name": "admin` + idCompany + `", "password": "` + password + `", "roles": ["members","admins"]}`
		log.Println("Password : " + password)
		models.CreateUserDB(datauser)
		models.InsertAuthorDB("c_"+idCompany, authJson)
		// log.Println(config.SetData("admin"+string(idCompany), ss))
	}
	return ""
}

func CheckSession(c *gin.Context) string {
	apikey := c.GetHeader("Authorization")
	if apikey == "" {
		return `{"error":"You must have api key"}`
	}
	splitToken := strings.Split(apikey, "App ")

	if len(splitToken) != 2 {
		return `{"error":"Your authorization must be : 'App {your apikey}'"}`
	}
	var decodedByte, _ = base64.StdEncoding.DecodeString(splitToken[1])
	return string(decodedByte)
}
func SetRedis(c *gin.Context) string {
	jsonData, _ := c.GetRawData()
	var objt SetObjt
	err := json.Unmarshal([]byte(string(jsonData)), &objt)
	if err == nil {
		return models.SetRedis(c.Query("key"), objt.Data)
	}
	return `{"error":` + err.Error() + `}`
}
func GetRedis(c *gin.Context) string {
	return models.GetRedis(c.Query("key"))
}
func InitializingData(idCompany string) {
	// masterwarehouse:=`{"name":"Gudang A","code":"GA","stokmin":3000,"stokmax":10000}`
	// models.BulkDocs("c"+idCompany,jsonData)
}
func CreateDBCompany(idCompany string) string {
	return models.CreateDatabase("c_" + idCompany)
}
func Find(c *gin.Context) string {
	var rst, err = config.CheckSession(c)
	if err == "" {
		db := c.Query("db")
		jsonData, _ := c.GetRawData()
		rst = models.FindDoc(db, jsonData)
	} else {
		rst = err
	}
	return rst
}
func Insert(c *gin.Context) string {
	db := c.Query("db")
	if db == "" {
		return `{"error":"Please insert db variable"}`
	}
	jsonData, _ := c.GetRawData()
	return models.InsertDoc(db, jsonData)
}
func Update(c *gin.Context) string {
	db := c.Query("db")
	if db == "" {
		return `{"error":"Please insert db variable"}`
	}
	jsonData, _ := c.GetRawData()
	return models.UpdateDoc(db, jsonData)
}
func Delete(c *gin.Context) string {
	db := c.Query("db")
	if db == "" {
		return `{"error":"Please insert db variable"}`
	}
	jsonData, _ := c.GetRawData()
	return models.DeleteDoc(db, jsonData)
}
func CreateDB(c *gin.Context) string {
	db := c.Query("db")
	if db == "" {
		return `{"error":"Please insert db variable"}`
	}
	return models.CreateDatabase(db)
}
func DeleteDB(c *gin.Context) string {
	db := c.Query("db")
	if db == "" {
		return `{"error":"Please insert db variable"}`
	}
	return models.DeleteDatabase(db)
}
func CreateUserDB(c *gin.Context) string {
	db := c.Query("db")
	if db == "" {
		return `{"error":"Please insert db variable"}`
	}
	jsonData, _ := c.GetRawData()
	return models.CreateUserDB(jsonData)
}
func GetView(c *gin.Context) string {
	if c.Query("db") == "" || c.Query("ddoc") == "" || c.Query("view") == "" {
		return `{"error":"db, ddoc, view parameters must included"}`
	}
	jsonData, _ := c.GetRawData()
	return models.GetView(c.Query("db"), c.Query("ddoc"), c.Query("view"), jsonData)
}
func BulkDocs(c *gin.Context) string {
	db := c.Query("db")
	jsonData, _ := c.GetRawData()
	return models.BulkDocs(db, jsonData)
}
func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}