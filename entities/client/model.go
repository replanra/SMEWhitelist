package client

import (
	"NewSMEWhitelist/config"
	"errors"
	"fmt"
	"log"
	"net/http"
	
)

type Member struct {
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

func AllMember() ([]Member, error) {
	//	rows, err := config.DB.Query("SELECT tblLoan.LoanID, tblLoan.firstname, tblLoan.lastname, tblLoan.Contact, tblLoan.amount, tblLoan.Date, tblLoan.Status, tblLoan.AppID, tblRelease.releaseID, tblRelease.appamount, tblRelease.monthly, tblrelease.amortization, tblRelease.branch, tblrelease.appid from tblLoan INNER JOIN tblRelease on tblLoan.LoanID = tblRelease.loanid;")
	rows, err := config.DB.Query("Select area, birthday, cid, centername, contact, flag, lengthofmembership, membername, newbranchcode, newcid, recognizeddate, sn, unit from tblclient;")
	if err != nil {

		log.Println(err)
		return nil, err

	}
	defer rows.Close()
	m := make([]Member, 0)
	for rows.Next() {
		ml := Member{}
		err := rows.Scan(&ml.Area, &ml.Birthday, &ml.CID, &ml.CenterName, &ml.Contact, &ml.Flag, &ml.LengthOfMembership, &ml.MemberName, &ml.NewBranchCode, &ml.NewCID, &ml.RecognizedDate, &ml.SN, &ml.Unit)
		if err != nil {
			return nil, err
		}
		m = append(m, ml)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return m, nil
}


func Update(r *http.Request) (Member, error) {
	m := Member{}
	m.Area = r.FormValue("area")
	m.Birthday = r.FormValue("birthday")
	m.CID = r.FormValue("cid")
	m.CenterName = r.FormValue("centername")
	m.Contact = r.FormValue("contact")
	m.Flag = r.FormValue("flag")
	m.LengthOfMembership = r.FormValue("lengthofmembership")
	m.MemberName = r.FormValue("membername")
	m.NewBranchCode = r.FormValue("newbranchcode")
	m.NewCID = r.FormValue("newcid")
	m.RecognizedDate = r.FormValue("recognizeddate")
	m.SN = r.FormValue("sn")
	m.Unit = r.FormValue("unit")

	var err error
	_, err = config.DB.Exec("UPDATE tblclient SET cid=$1,  area = $2, birthday = $3,  centername=$4, contact=$5, flag=$6, lengthofmembership=$7, membername=$8, newbranchcode=$9, newcid=$10, recognizeddate=$11, sn=$12, unit=$13 WHERE cid=$1;", m.CID, m.Area, m.Birthday, m.CenterName, m.Contact, m.Flag, m.LengthOfMembership, m.MemberName, m.NewBranchCode, m.NewCID, m.RecognizedDate, m.SN, m.Unit )

	if err != nil {
		//return ln, err
		log.Println(err)

		return m, errors.New("500. Internal Server Error." + err.Error())
	}
	log.Println("Client Updated")
	return m, nil
}

func DeleteClient(r *http.Request) error {
	cid := r.FormValue("cid")
	if cid == "" {
		return errors.New("400. Bad Request.")
	}
	_, err := config.DB.Exec("DELETE FROM tblclient WHERE cid=$1;", cid)
	if err != nil {
		return errors.New("500. Internal Server Error")
	}
	return nil
}

func Insert(r *http.Request) (Member, error) {
	m := Member{}
	m.Area = r.FormValue("area")
	m.Birthday = r.FormValue("birthday")
	m.CID = r.FormValue("cid")
	m.CenterName = r.FormValue("centername")
	m.Contact = r.FormValue("contact")
	m.Flag = r.FormValue("flag")
	m.LengthOfMembership = r.FormValue("lengthofmembership")
	m.MemberName = r.FormValue("membername")
	m.NewBranchCode = r.FormValue("newbranchcode")
	m.NewCID = r.FormValue("newcid")
	m.RecognizedDate = r.FormValue("recognizeddate")
	m.SN = r.FormValue("sn")
	m.Unit = r.FormValue("unit")
	

	fmt.Println(m)
	var err error
	_, err = config.DB.Exec("INSERT INTO tblclient(area, birthday, cid, centername, contact, flag, lengthofmembership, membername, newbranchcode, newcid, recognizeddate, sn, unit) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)", m.Area, m.Birthday, m.CID, m.CenterName, m.Contact, m.Flag, m.LengthOfMembership, m.MemberName, m.NewBranchCode, m.NewCID, m.RecognizedDate, m.SN, m.Unit)
	if err != nil {
		log.Println(err)
		return m, errors.New("500. Internal Server Error." + err.Error())
	}
	log.Println("New Client Added")
	return m, nil
}


func CheckClient(w http.ResponseWriter, r *http.Request) ([]Member, error) {
	newcid := r.FormValue("newcid")
	rows, err := config.DB.Query("Select area, birthday, cid, centername, contact, flag, lengthofmembership, membername, newbranchcode, newcid, recognizeddate, sn, unit from tblclient where newcid = '" + newcid + "' ;")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()
	cls := make([]Member, 0)
	for rows.Next() {
		cl := Member{}
		err := rows.Scan(&cl.Area, &cl.Birthday, &cl.CID, &cl.CenterName, &cl.Contact, &cl.Flag, &cl.LengthOfMembership, &cl.MemberName, &cl.NewBranchCode, &cl.NewCID, &cl.RecognizedDate, &cl.SN, &cl.Unit)
		if err != nil {
			return nil, err
		}
		
		cls = append(cls, cl)
       
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	
	fmt.Println(cls)
	return cls, nil
}























