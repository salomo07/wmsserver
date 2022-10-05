package models
import (
	"context"
	"wms/config"
)
var (db=config.Connect())

func Find()(Animal){
	find:=map[string]interface{}{}
	row := db.Get(context.TODO(), "xxx",find)


	var u Animal
	err := row.ScanDoc(&u)
	if err != nil {
	    panic(err)
	}
	return u
}

func StoreData(){
	cow := Animal{ID: "cow", Feet: 4, Greeting: "moo"}
	rev, err := db.Put(context.TODO(), "cow", cow)
	if err != nil {
	    panic(err)
	}
	cow.Rev = rev
}
type Animal struct {
    ID       string `json:"_id"`
    Rev      string `json:"_rev,omitempty"`
    Feet     int    `json:"feet"`
    Greeting string `json:"greeting"`
}
type User struct {
    ID       string `json:"_id"`
    Rev      string `json:"_rev,omitempty"`
    Feet     int    `json:"feet"`
    Greeting string `json:"greeting"`
}