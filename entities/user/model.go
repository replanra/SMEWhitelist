package user

import (
	"NewSMEWhitelist/config"
	"database/sql"
	"encoding/gob"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	
	"golang.org/x/crypto/bcrypt"
)

type userCreds struct {
	Uid           string `json:"uid"`
	Fname         string `json:"fname"`
	Lname         string `json:"lname"`
	Username      string `json:"username"`
	Password      string `json:"password"`
	Branch        string `json:"branch"`
	Insti         string `json:"insti"`
	Authenticated bool
}
type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Register struct {
	Uid      string `json:"uid"`
	Fname    string `json:"fname"`
	Lname    string `json:"lname"`
	Username string `json:"username"`
	Password string `json:"password"`
	Branch   string `json:"branch"`
	Insti    string `json:"insti"`
}

var Store *sessions.CookieStore

func init() {
	authKeyOne := []byte("::AUTH-ONE::")
	encryptionKeyOne := []byte("::AK&-L@&@Q-SKS-LSQ-aI::")

	Store = sessions.NewCookieStore(
		authKeyOne,
		encryptionKeyOne,
	)

	Store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   60 * 15,
		HttpOnly: true,
		SameSite: 2,
	}

	gob.Register(userCreds{})
}

var granted bool

var act string

func GetUser(s *sessions.Session) userCreds {
	val := s.Values["user"]
	var user = userCreds{}
	user, ok := val.(userCreds)
	if !ok {
		return userCreds{Authenticated: false}
	}

	return user
}

func Chechrecord(r *http.Request) ([]Register, error) {

	//	dot, err := dotsql.LoadFromFile("queries.sql")
	//rows, err := dot.Query(config.DB, "search-loan")
	fname := r.FormValue("fname")
	lname := r.FormValue("lname")
	rows, err := config.DB.Query("Select * from tbluser where fname = '" + fname + "' AND lname= '" + lname + "';")

	//rows, err := config.DB.Query("SELECT * from tblLoan WHERE dateapplied BETWEEN '" + from + "' AND '" + to + "' and status='Assigned';")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	rgr := make([]Register, 0)
	for rows.Next() {
		rg := Register{}
		err := rows.Scan(&rg.Uid, &rg.Fname, &rg.Lname, &rg.Username, &rg.Password, &rg.Branch, &rg.Insti)
		//	err := rows.Scan(&lns.Lid, &lns.Amount, &lns.PaymentNo, &lns.Purpose, &lns.PN, &lns.DateApplied, &lns.Encoded, &lns.Status, &lns.Category, &lns.Pgroup, &lns.Product, &lns.Frequency, &lns.Guarantor, &lns.Cid)
		if err != nil {
			return nil, err
		}
		rgr = append(rgr, rg)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	fmt.Println(rgr)
	return rgr, nil
}



func InsertUserTest(w http.ResponseWriter, r *http.Request) (Register, error) {

	rg := Register{}
	rg.Fname = r.FormValue("fname")
	rg.Lname = r.FormValue("lname")
	rg.Username = r.FormValue("username")
	rg.Password = r.FormValue("textareaID")
	rg.Branch = r.FormValue("branch")
	rg.Insti = r.FormValue("insti")
	fmt.Println(rg)
	if rg.Fname == "" || rg.Lname == "" || rg.Username == "" || rg.Password == "" {
		//return rg, errors.New("400. Bad Request. Fields can't be empty.")
		//w.Write([]byte("Fields can't be empty"))

	} else {
		userSQLStatement := `
		SELECT username
		FROM tbluser
		WHERE username=$1;`

		//var username string
		row := config.DB.QueryRow(userSQLStatement, &rg.Username)
		switch err := row.Scan(&rg.Username); err {
		case sql.ErrNoRows:
			//Password Validation
			password := rg.Password
			hash, err := HashPassword(password)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				fmt.Println("Something went wrong in generating hash password. Please try again.")
			}

			// userHash, err := HashPassword(username)

			// if err != nil {
			// 	http.Error(w, err.Error(), http.StatusInternalServerError)
			// 	fmt.Println("Something went wrong in generating hash password. Please try again.")
			// }

			match := CheckPasswordHash(password, hash)

			switch match {
			case true:

				if _, err = config.DB.Exec("INSERT INTO tbluser (fname, lname, username, password, branch, insti) VALUES ($1, $2, $3, $4, $5, $6)", rg.Fname, rg.Lname, rg.Username, hash, rg.Branch, rg.Insti); err != nil {
					w.Write([]byte("userError"))
					panic(err)

				}

			case false:
				w.Write([]byte("notMatch"))
				fmt.Println("Password Doesn't Match")
			}

		case nil:
			http.Redirect(w, r, "/logg", http.StatusMovedPermanently)
			w.Write([]byte("Username exist"))
			
	
		default:
			panic(err)
		}

	}

	return rg, nil

	/*var err error
	_, err = config.DB.Exec("INSERT INTO tbluser(fname, lname, username, password, branch, insti) VALUES ($1, $2, $3, $4, $5, $6)", rg.Fname, rg.Lname, rg.Username, rg.Password, rg.Branch, rg.Insti)
	if err != nil {
		log.Println(err)
		return rg, errors.New("500. Internal Server Error." + err.Error())
	}
	log.Println("New User Added")
	return rg, nil*/
}




func SignIn(w http.ResponseWriter, r *http.Request) ([]Login, error) {
	if r.Method == "POST" {
		log.Println("model")
		session, err := Store.Get(r, "cookie-name")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			//	return
		}

		username := r.FormValue("username")
		password := r.FormValue("password")

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			fmt.Println("Something went wrong in generating hass password. Please try again.")
		}

		result := config.DB.QueryRow("SELECT uid, fname, lname, username, password, branch, insti FROM tbluser WHERE username=$1", username)
		if err != nil {
			// If there is an issue with the database, return a 500 error
			panic(err)

		}

		storedCreds := &userCreds{}
		err = result.Scan(
			&storedCreds.Uid,
			&storedCreds.Fname,
			&storedCreds.Lname,
			&storedCreds.Username,
			&storedCreds.Password,
			&storedCreds.Branch,
			&storedCreds.Insti,
		)

		if err != nil {
			// If an entry with the username does not exist, send an "Unauthorized"(401) status
			if err == sql.ErrNoRows {
				http.Redirect(w, r, "/log", http.StatusMovedPermanently)
				return nil, err
			}
   
			http.Redirect(w, r, "/log", http.StatusMovedPermanently)
			return nil, err

		}

		if err = bcrypt.CompareHashAndPassword([]byte(storedCreds.Password), []byte(password)); err != nil {

			granted = false

		} else {
			granted = true
		}

		log.Println("Granted: ", granted)
		log.Println("Input Password: ", password)
		log.Println("Stored Password: ", storedCreds.Password)

		switch granted {
		case true:
			creds := &userCreds{
				Fname:         storedCreds.Fname,
				Lname:         storedCreds.Lname,
				Username:      storedCreds.Username,
				Branch:        storedCreds.Branch,
				Insti:         storedCreds.Insti,
				Uid:           storedCreds.Uid,
				Authenticated: true,
			}

			session.Values["user"] = creds

			err = session.Save(r, w)

			if err != nil {
				log.Println(err)
			}
			http.Redirect(w, r, "/uploadfile", http.StatusMovedPermanently)

		case false:
			http.Redirect(w, r, "/log", http.StatusMovedPermanently)
			log.Println("invalid Use")
			break
		}

	}
	var err error
	return nil, err
}




