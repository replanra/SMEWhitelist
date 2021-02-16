package main

import (
	"RA/entities/client"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	mux := http.NewServeMux()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	//log.Fatal(http.ListenAndServe(port, mux))
	//Used port := os.Getenv("PORT")
	//"PORT" variable name na ansa .env file
	//kapag nasa local sya nag run, yung 8080 na port ang kukunin nya. yun yung nasa .env na file
	//then kapag sa heroku or other platform naman, kung anong ibinigay nilang port
	//ay maadopt na ni application, so there's no need to manually update the port kapag inilagay mo sa ibang hosting

	log.Println("Server Started...")
	log.Println("Current Port: " + port)

	mux.HandleFunc("/burn", client.GetAllClients)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		panic(err)
	}
}
