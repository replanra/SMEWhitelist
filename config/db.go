package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

//Local DB Connection
const (
	host     = "ec2-54-147-126-202.compute-1.amazonaws.com"
	port     = 5432
	user     = "jkrumqelyzvjko"
	password = "7cb7730b4fd79d4fd42303874880a8b0d9900a8308bb4bf4786498ae2645d4b7"
	dbname   = "dncm7pakkq5vl"
)

var DB *sql.DB

//sslmode=require kapag sa heroku
//sslmode=disable kapag sa local
func init() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=require",
		host, port, user, password, dbname)
	var err error
	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	if err = DB.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("You connected to your database.")
}
