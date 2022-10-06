package models
import (
	"log"
	"time"
	"context"
	"wms/config"
)
var (db=config.Connect())

func Query()(User){
	selector:=map[string]interface{}{"username":"salomo07"}
	rows,errq := db.Query(context.TODO(),"_design/user", "_view/user",selector)

	log.Println(rows,errq,"ddd")
	var u User
	// err := rows.ScanDoc(&u)
	// if err != nil {
	//     panic(err)
	// }
	return u
}

func StoreData(){
	timenow:=time.Now().Unix()
	u := User{Username: "salomo07", Password: "fff"}
	log.Println(timenow)
	_, err := db.Put(context.TODO(),"kodok", u)
	if err != nil {
	    panic(err)
	}
}
type User struct {
    Username      string `json:"username"`
    Password     string    `json:"password"`
}