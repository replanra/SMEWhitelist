package user

import (
	"newWhitelist/config"
	"log"
	
	"net/http"
)

func SamLogin(w http.ResponseWriter, r *http.Request) {
	config.TPL.ExecuteTemplate(w, "log", nil)
}

func SamLogg(w http.ResponseWriter, r *http.Request) {
	config.TPL.ExecuteTemplate(w, "logg", nil)
}


func SampeLogin(w http.ResponseWriter, r *http.Request) {
	config.TPL.ExecuteTemplate(w, "login", nil)
}
func SampleRegister(w http.ResponseWriter, r *http.Request) {
	config.TPL.ExecuteTemplate(w, "sampleregister", nil)
}
func LoginUserForm(w http.ResponseWriter, r *http.Request) {
	config.TPL.ExecuteTemplate(w, "loginuser", nil)
}
func LoginUserProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	_, err := SignIn(w, r)
	if err != nil {
		log.Println("Unable to procees: ", err)

	}

}
func RegisterUserForm(w http.ResponseWriter, r *http.Request) {
	config.TPL.ExecuteTemplate(w, "registeruser", nil)
}
func Password(w http.ResponseWriter, r *http.Request) {
	config.TPL.ExecuteTemplate(w, "password", nil)
}

func RegisterUserProcessTest(w http.ResponseWriter, r *http.Request,) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	_, err := InsertUserTest(w, r)
	if err != nil {
		panic(err)
		//return
	}


	http.Redirect(w, r, "/login", http.StatusMovedPermanently)
	

}

