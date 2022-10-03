package models
import (
  kivik "github.com/go-kivik/kivik/v3"
)
var (db *kivik.DB)
func init(){
	// db=config.Connect()
}
type User struct {
    ID       string `json:"_id"`
    Rev      string `json:"_rev,omitempty"`
    Feet     int    `json:"feet"`
    Greeting string `json:"greeting"`
}