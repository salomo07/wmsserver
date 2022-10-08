package models
import (
	kivik "github.com/go-kivik/kivik/v3"
	"strconv"
	"log"
	"time"
	"context"
	"wms/config"
)
var (db=config.Connect())
func Find(s map[string]interface{})([] User){
	sss:=map[string]interface{}{"selector":s}
	rows,err := db.Find(context.TODO(),sss,map[string]interface{}{"include_docs":true})
	if err != nil {
	    panic(err)
	}
	var result [] User
	for rows.Next() {
	    var u User
	    if err := rows.ScanDoc(&u); err != nil {
			panic(err)
		}
		
		result=append(result, u)
	    /* do something with doc */
	}
	log.Println(result,"ddd")
	return result
}
func Query()(){
	rows, err := db.Query(context.TODO(), "_design/user", "_view/bar", kivik.Options{
		"startkey": `"foo"`,                           // Quotes are necessary so the
		"endkey":   `"foo` + kivik.EndKeySuffix + `"`, // key is a valid JSON object
	})
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var doc interface{}
		if err := rows.ScanDoc(&doc); err != nil {
			panic(err)
		}
		/* do something with doc */
	}
	if rows.Err() != nil {
		panic(rows.Err())
	}
		// return rows
	}
// func Query()(){
// 	rows,err := db.Query(context.TODO(),"_design/user", "_view/user",kivik.Options{"include_docs": true})

// 	if err != nil {
// 	    panic(err)
// 	}
// 	for rows.Next() {
// 	    var u User
// 	    if err := rows.ScanDoc(&u); err != nil {
// 	        panic(err)
// 	    }
// 	    log.Println(u,"ddd")
// 	    /* do something with doc */
// 	}
// 	if rows.Err() != nil {
// 	    panic(rows.Err())
// 	}
// 	// log.Println(rows)
// 	// return rows
// }

func StoreData(){
	timenow:=time.Now().Unix()
	u := User{Username: "salomo07", Password: "fff"}
	log.Println(timenow)
	_, err := db.Put(context.TODO(),strconv.FormatInt(timenow, 10), u)
	if err != nil {
	    panic(err)
	}
}
type User struct {
    Username      string `json:"username"`
    Password     string    `json:"password"`
    Contact     Contact    `json:"contact"`
}
type Selector struct {
    Selector      User `json:"selector"`
    Limit		int `json:"limit"`
}
type Contact struct {
    Mobile      string `json:"mobile"`
    Email		string `json:"email"`
}
// type Selector struct {
//     Selector      User `json:"selector"`
//     Limit		int `json:"limit"`
//     Skip int  `json:"skip"`
//     Short []Short `json:"skip"`
//     Field []string `json:"fields"`
// }
// type Short struct {

// }
// type Field struct {
	
// }