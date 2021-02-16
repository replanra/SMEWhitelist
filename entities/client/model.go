package client

import (
	"RA/config"
	"log"
)

// field for data base
type ClientList struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// data type client
func AllClients() ([]ClientList, error) {
	rows, err := config.DB.Query("Select id, name, username, password from replan;")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()
	cls := make([]ClientList, 0)
	for rows.Next() {
		cl := ClientList{}
		err := rows.Scan(&cl.Id, &cl.Name, &cl.Username, &cl.Password)
		if err != nil {
			return nil, err
		}
		cls = append(cls, cl)
		//fmt.Println(cl)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return cls, nil
}
