package upload

import (
	"newWhitelist/config"
	"newWhitelist/entities/client"

	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	//"text/template"
	"log"
	//"time"
)

type MemberList struct {
	Area               string `json:"area"`
	Birthday           string `json:"birthday"`
	CID                string `json:"cid"`
	CenterName         string `json:"centername"`
	Contact            string `json:"contact"`
	Flag               string `json:"flag"`
	LengthOfMembership string `json:"lengthofmembership"`
	MemberName         string `json:"membername"`
	NewBranchCode      string `json:"newbranchcode"`
	NewCID             string `json:"newcid"`
	RecognizedDate     string `json:"recognizeddate"`
	SN                 string `json:"sn"`
	Unit               string `json:"unit"`
}
type UploadData struct {
	NilFile, InvalidFile bool
	Clients              map[string]MemberList
}

/*func Cookies(w http.ResponseWriter, req *http.Request) {

	session, err := user.Store.Get(req, "cookie-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user := user.GetUser(session)
	if auth := user.Authenticated; !auth {
		session.AddFlash("You don't have access!")
		err = session.Save(req, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, req, "/login/user", http.StatusFound)

		return
	}
	fmt.Println(user)
	config.TPL.ExecuteTemplate(w, "uploadfile.html", user)
}*/

func UploadFile(w http.ResponseWriter, req *http.Request) {

	if req.Method == "GET" {
		config.TPL.ExecuteTemplate(w, "uploadfile", nil)
	} else if req.Method == "POST" {
		file, handler, err := req.FormFile("uploadfile")
		if err != nil {

			
			log.Println("Null File: ", err)

			config.TPL.ExecuteTemplate(w, "uploadfile", nil)
			return

		} else {
			fmt.Println("error throws in else statement")
			fmt.Println("handler.Filename", handler.Filename)
			fmt.Printf("Type of handler.Filename:%T\n", handler.Filename)
			fmt.Println("Length:", len(handler.Filename))

			f, err := os.OpenFile("./data/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				log.Println("Open File Error: ", err)
				config.TPL.ExecuteTemplate(w, "uploadfile", nil)
				return
			}
			defer f.Close()
			io.Copy(f, file)
			blobPath := "./data/" + handler.Filename
			var extension = filepath.Ext(blobPath)
			parsedData := ExcelCsvParser(blobPath, extension)
			parsedJson, _ := json.Marshal(parsedData)
			//fmt.Println(string(parsedJson))

			err = os.Remove(blobPath)
			keysBody := []byte(parsedJson)
			keys := make([]MemberList, 0)
			json.Unmarshal(keysBody, &keys)
			_, err = config.DB.Exec("INSERT INTO tbltest SELECT area, birthday, cid, centername, contact, flag, lengthofmembership, membername, newbranchcode, newcid, recognizeddate, sn, unit FROM json_populate_recordset(NULL::tbltest,  '" + string(keysBody) + "') ON CONFLICT (cid) DO NOTHING ")
			if extension != ".csv" {
				//w.Write([]byte("It's not CSV file"))
				config.TPL.ExecuteTemplate(w, "invalidfile", nil)
			} else if err != nil {
				//http.Error(w, http.StatusText(500), http.StatusInternalServerError)
				log.Println("Invalid File: ", err)
				panic(err)

			} else {
				ac, err := client.AllMember()
				if err != nil {
					http.Error(w, http.StatusText(500)+err.Error(), http.StatusInternalServerError)
				}
				config.TPL.ExecuteTemplate(w, "uploadfile", ac)
			}
		}

	}
}

func ExcelCsvParser(blobPath string, blobExtension string) (parsedData []map[string]interface{}) {
	fmt.Println("---------------> We are in product.go")
	if blobExtension == ".csv" {
		fmt.Println("-------We are parsing an csv file.-------------")
		parsedData := ReadCsvFile(blobPath)
		fmt.Printf("Type:%T\n", parsedData)
		return parsedData

	} /*else if blobExtension == ".xlsx" {
		fmt.Println("----------------We are parsing an xlsx file.---------------")
		parsedData := parser.ReadXlsxFile(blobPath)
		return parsedData
	} else if blobExtension == ".xls" {
		fmt.Println("----------------We are parsing an xls file.---------------")
		parsedData := parser.ReadXlsFile(blobPath)
		return parsedData
	}*/
	return parsedData
}

func init() {
	path := "./data"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0777)
		fmt.Println("Created data directory")
	} else {
		fmt.Println("Data directory already exists")
	}
}
