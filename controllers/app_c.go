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

type SessionErrorObjt struct {
	Error string `json:"error" binding:"required"`
}

// type ResultObjt struct {
// 	Docs []models.CompanyObjt `json:"docs"`
// }

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

		// CreateDBCompany(idCompany)
		dbname := "db_" + idCompany
		userdb := "admin" + idCompany
		password := strconv.FormatInt(time.Now().UnixNano(), 10)
		authJson := `{"admins": { "names": ["` + userdb + `"],"roles":["admins"]}}`
		userDBJson := `{"name": "` + userdb + `", "password": "` + password + `", "roles": ["members","admins"],"type":"user"}`
		sessionJson := `{"userdb":"` + userdb + `","passdb":"` + password + `",dbname:"` + dbname + `"}`
		log.Println("ID :" + dbname + " Password : " + password)
		log.Println("CreateUserDB ---- " + models.CreateUserDB(userDBJson))
		log.Println("CreateDatabase ---- " + models.CreateDatabase(dbname))
		log.Println("InsertAuthorDB ---- " + models.InsertAuthorDB(dbname, authJson))
		succRedis, errRedis := config.SetDataRedis(userdb, sessionJson)
		if errRedis != "" {
			return `{"error":"` + errRedis + `"}`
		}
		log.Println("SetData ---- " + succRedis)
	}
	return ""
}
func AddRolesDefault(c *gin.Context) string {
	return ""
	// log.Println("Add role ---- " + models.InsertDocByCompany(dbname, authJson))
}
func InsertAuthorDB(c *gin.Context) string {
	jsonData, _ := c.GetRawData()
	db := c.Query("db")
	return models.InsertAuthorDB(db, string(jsonData))
}
func CheckUserSession(c *gin.Context) (bool, string, models.UserModel) {
	apikey := c.GetHeader("Authorization")
	var usr models.UserModel
	if apikey == "" {
		return false, "You must have api key", usr
	}
	splitToken := strings.Split(apikey, "App ")

	if len(splitToken) != 2 {
		return false, "Your authorization must be : 'App {your apikey}'", usr
	}
	userId, err := base64.StdEncoding.DecodeString(splitToken[1])
	if err != nil {
		return false, err.Error(), usr
	}

	var errsession SessionErrorObjt
	usrData := models.GetRedis(string(userId))
	json.Unmarshal([]byte(usrData), &errsession)
	if errsession.Error != "" {
		return false, `Session is not found, please re-login`, usr
	} else {

		json.Unmarshal([]byte(usrData), &usr)
		return true, "", usr
	}
}
func TryLogin() string {
	log.Println(`{"selector":{"users":[""]},"fields": ["_id"]}`)
	findUser := `{"selector":{"users":[""]},"fields": ["_id]}`
	companyData := models.FindDocByRoot("mastercompany", []byte(findUser))
	var comObject models.CompanyModels
	json.Unmarshal([]byte(companyData), &comObject)
	// if len(comObject.Docs) == 0 {
	// 	return `{"error":"Youre not register yet"}`
	// } else {
	// 	company, _ := json.Marshal(comObject.Docs[0])
	// 	return string(company)
	// }
	return ""
}
func GetSessionCred(c *gin.Context) (string, string, string) {
	// valid, msg, userData := CheckUserSession(c)
	// log.Println(userData, msg)
	// var cred models.CredDBObjt
	// if valid == true {
	// 	//	Get CredDB from Redis
	// 	log.Println(models.GetRedis("admin" + userData.Idcompany))
	// 	json.Unmarshal([]byte(models.GetRedis("admin"+userData.Idcompany)), &cred)
	// 	return cred.Name + ":" + cred.Password, cred.DBName, ""
	// }

	return "", "", ""
}

func SetRedis(c *gin.Context) string {
	jsonData, _ := c.GetRawData()
	res, err := models.SetRedis(c.Query("key"), string(jsonData))
	if err != "" {
		return err
	}
	return res
}
func GetRedis(c *gin.Context) string {
	return models.GetRedis(c.Query("key"))
}
func InitializingData(idCompany string) {
	// masterwarehouse:=`{"name":"Gudang A","code":"GA","stokmin":3000,"stokmax":10000}`
	// models.BulkDocs("c"+idCompany,jsonData)
}

// func CreateDBCompany(user string, pass string, idCompany string) string {
// 	return models.CreateDatabase("c_" + idCompany)
// }
func Find(c *gin.Context) string {

	// if credDB != nil {
	// 	db := c.Query("db")
	// 	jsonData, _ := c.GetRawData()
	// 	rst = models.FindDoc(db, jsonData)
	// } else {
	// 	rst = err
	// }
	return ""
}
func Insert(c *gin.Context) string {
	creddb, dbname, msg := GetSessionCred(c)
	if creddb == "" {
		return `{"error":"` + msg + `"}`
	} else {
		log.Println(creddb, dbname, msg)
		return models.InsertDocByCompany(creddb, dbname, msg)
	}
	// return models.InsertDoc(db, jsonData)
}
func Insert2(c *gin.Context) string {
	db := c.Query("db")
	if db == "" {
		return `{"error":"Please insert db variable"}`
	}
	jsonData, _ := c.GetRawData()
	var user models.UserModel
	json.Unmarshal(jsonData, &user)
	log.Println(GetMD5Hash(user.Password))
	user.Password = GetMD5Hash(user.Password)
	rst, _ := json.Marshal(user)
	log.Println(string(rst))
	log.Println("success")
	return models.InsertDoc(db, rst)
	// return models.InsertDoc(db, jsonData)
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
	return models.CreateUserDB(string(jsonData))
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
