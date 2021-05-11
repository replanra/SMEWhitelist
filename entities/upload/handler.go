package upload

import (
	"NewSMEWhitelist/config"
	"net/http"
)

func UploadForm(w http.ResponseWriter, r *http.Request) {
	config.TPL.ExecuteTemplate(w, "allmember", nil)
	
}

