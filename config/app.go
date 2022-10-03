package config

import(
	"os"
	"context"
	"github.com/joho/godotenv"
	kivik "github.com/go-kivik/kivik/v3"
	_"github.com/go-kivik/couchdb/v3"
	"log"
)
func Connect(){
	// ctx := context.Background()
	er := godotenv.Load()
	if er !=nil{
		panic("Fail to load .env file")
	}
	// log.Println(ctx)
	DB_URL:=os.Getenv("DB_URL")
	client, err := kivik.New("couch", DB_URL)
    if err != nil {
        panic(err)
    }
    db := client.DB(context.TODO(), "wms")
    log.Println(db)
}