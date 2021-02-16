package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

//Local DB Connection
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "rareplan"
)

var DB *sql.DB

//sslmode=require kapag sa heroku
//sslmode=disable kapag sa local
func init() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
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
