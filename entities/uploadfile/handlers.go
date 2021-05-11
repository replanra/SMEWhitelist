package uploadfile

import (
	"newWhitelist/config"
	"net/http"
)

func Uploadlists(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	up, err := ListofUploadFile()
	if err != nil {
		http.Error(w, http.StatusText(500)+err.Error(), http.StatusInternalServerError)
		return
	}
	config.TPL.ExecuteTemplate(w, "listofuploadedfile", up)
}
