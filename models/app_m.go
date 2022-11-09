package models

import (
	"encoding/json"
	"wms/config"

	"github.com/gin-gonic/gin"
)

//Buat validator/verifikator object
type UpdateObjt struct {
	Id  string `json:"_id"`
	Rev string `json:"_rev"`
}
type UserDBObjt struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Type string `json:"type"`
}

func FindDoc(db string, jsonData []byte) string {
	return config.FindDoc(db+"/_find", string(jsonData))
}
func InsertDoc(db string, jsonData []byte) string {
	return config.InsertDoc(db, string(jsonData))
}
func UpdateDoc(db string, jsonData []byte) string {
	var objt UpdateObjt
	err := json.Unmarshal([]byte(string(jsonData)), &objt)
	if err != nil {
		return `{"error":"` + err.Error() + `"}`
	}
	if objt.Id == "" || objt.Rev == "" {
		return `{"error":"_id & _rev must be includes"}`
	}
	return config.UpdateDoc(db+"/"+objt.Id, string(jsonData))
}
func DeleteDoc(db string, jsonData []byte) string {
	var objt UpdateObjt

	err := json.Unmarshal([]byte(string(jsonData)), &objt)
	if err != nil {
		return `{"error":"` + err.Error() + `"}`
	}
	if objt.Id == "" || objt.Rev == "" {
		return `{"error":"_id & _rev must be includes"}`
	}
	return config.DeleteDoc(db + "/" + objt.Id + "?rev=" + objt.Rev)
}
func CreateDatabase(db string) string {
	return config.CreateDatabase(db)
}
func DeleteDatabase(db string) string {
	return config.DeleteDatabase(db)
}
func GetView(db string, ddoc string, view string, jsonData []byte) string {
	return config.GetDataByView(db, ddoc, view, string(jsonData))
}
func BulkDocs(db string, jsonData []byte) string {
	return config.BulkDocs(db, string(jsonData))
}
func CreateUserDB(jsonData []byte) string {
	var objt UserDBObjt
	err := json.Unmarshal([]byte(string(jsonData)), &objt)
	if err != nil {
		return `{"error":"` + err.Error() + `"}`
	}
	if objt.Name == "" || objt.Password == "" {
		return `{"error":"Please insert name & password"}`
	}
	objt.Type="user"
	return config.CreateUserDB(objt.Name, string(jsonData))
}
func CreateReplication(c *gin.Context) string {
	jsonData, _ := c.GetRawData()
	return config.CreateReplication(string(jsonData))
}
func SetRedis(key string, val string) string {
	return config.SetData(key, val)
}
func GetRedis(key string) string {
	return config.GetData(key)
}
func InsertAuthorDB(db string,strquery string)(string){
	return InsertAuthorDB(db,strquery)
}
