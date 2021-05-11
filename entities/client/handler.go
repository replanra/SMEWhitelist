package client

import (
	"newWhitelist/config"
	"fmt"
	"encoding/json"
	"net/http"
)

func Uf(w http.ResponseWriter, r *http.Request) {
	config.TPL.ExecuteTemplate(w, "uf", nil)
}

func Login(w http.ResponseWriter, r *http.Request) {
	config.TPL.ExecuteTemplate(w, "login", nil)
}

func Exit(w http.ResponseWriter, r *http.Request) {
	config.TPL.ExecuteTemplate(w, "exit", nil)
}
func GetAllMember(w http.ResponseWriter, r *http.Request) {
	mls, err := AllMember()
	if err != nil {
		http.Error(w, http.StatusText(500)+err.Error(), http.StatusInternalServerError)
		return
	}
	config.TPL.ExecuteTemplate(w, "datatable", mls)
	/*w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Request-Reference-No", "`1e9ac446-8a62-4ae3-852d-c352ceda99b`")
	json.NewEncoder(w).Encode(mls)*/
	
}

func DeleteClientProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	err := DeleteClient(r)
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "allmember", http.StatusSeeOther)
}

func UpdateClientProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	_, err := Update(r)
	if err != nil {
		http.Error(w, http.StatusText(406), http.StatusBadRequest)
		return
	}
	http.Redirect(w, r, "allmember", http.StatusMovedPermanently)
	fmt.Println("Client Updated")
}


func InsertClientProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	_, err := Insert(r)
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "allmember", http.StatusSeeOther)
}

func CheckList(w http.ResponseWriter, r *http.Request) {
	cl, err := CheckClient(w, r)
	if err != nil {
		http.Error(w, http.StatusText(200)+err.Error(), http.StatusInternalServerError)
		return
	}
	/*	if cl == nil {
		http.Error(w, http.StatusText(400), http.StatusMethodNotAllowed)
	}*/
	//http.Error(w, http.StatusText(400), http.StatusMethodNotAllowed)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Request-Reference-No", "`1e9ac446-8a62-4ae3-852d-c352ceda99b`")
	json.NewEncoder(w).Encode(cl)
}

func CK(w http.ResponseWriter, r *http.Request) {
	config.TPL.ExecuteTemplate(w, "cl", nil)
}