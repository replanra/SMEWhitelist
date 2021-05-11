package main

import (
	"newWhitelist/entities/client"
	"newWhitelist/entities/upload"
	"newWhitelist/entities/user"
   

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

	log.Println("Server Started...")
	log.Println("Current Port: " + port)
	mux.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public"))))
	mux.Handle("/templates/", http.StripPrefix("/templates", http.FileServer(http.Dir("templates"))))

	//end
	//html
	mux.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir("css"))))
	mux.Handle("/js/", http.StripPrefix("/js", http.FileServer(http.Dir("js"))))
	mux.Handle("/json-browse/", http.StripPrefix("/json-browse", http.FileServer(http.Dir("json-browse"))))


	mux.HandleFunc("/exit", client.Exit)




    mux.HandleFunc("/update", client.UpdateClientProcess)
	mux.HandleFunc("/insert", client.InsertClientProcess)
    mux.HandleFunc("/delete", client.DeleteClientProcess)
	mux.HandleFunc("/uf", client.Uf)
	mux.HandleFunc("/login", client.Login)
	mux.HandleFunc("/allmember", client.GetAllMember)
	mux.HandleFunc("/checkclient", client.CheckList)
	mux.HandleFunc("/cl", client.CK)




	mux.HandleFunc("/uploadfile", upload.UploadFile)
	mux.HandleFunc("/upload", upload.UploadForm)
	mux.HandleFunc("/log", user.SamLogin)
	mux.HandleFunc("/logg", user.SamLogg)
	



	
    mux.HandleFunc("/login/user/process", user.LoginUserProcess)
	mux.HandleFunc("/password", user.Password)
	mux.HandleFunc("/register/user", user.RegisterUserForm)
	mux.HandleFunc("/registertest", user.RegisterUserProcessTest)
	
	

	
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		panic(err)
	}
}
 