package upload

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
)

func ReadCsvFile(filePath string) []map[string]interface{} {
	// Load a csv file.
	f, _ := os.Open(filePath)
	// Create a new reader.
	r := csv.NewReader(bufio.NewReader(f))
	result, _ := r.ReadAll()
	parsedData := make([]map[string]interface{}, 0, 0)
	header_name := result[0]

	for row_counter, row := range result {

		if row_counter != 0 {
			var singleMap = make(map[string]interface{})
			for col_counter, col := range row {
				singleMap[header_name[col_counter]] = col
			}
			if len(singleMap) > 0 {

				parsedData = append(parsedData, singleMap)
			}
		}
	}

	fmt.Println("Length of parsedData:", len(parsedData))
	return parsedData

}

/*
func InsertClient(r *http.Request) (ClientList, error) {
	cl := ClientList{}
	fmt.Println(cl)
	var err error
	_, err = config.DB.Exec("INSERT INTO tblclient(flag, sn, area, unit, cid, centername, membername, rdate, bday, mlength, bcode, newcid, contact) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)", cl.Flag, cl.SN, cl.Area, cl.Unit, cl.Cid, cl.Centername, cl.Membername, cl.RegisteredDate, cl.Bday, cl.LengthMembers, cl.BranchCode, cl.Newcid, cl.Contact)
	if err != nil {
		log.Println(err)
		return cl, errors.New("500. Internal Server Error." + err.Error())
	}
	log.Println("New Client Added")
	return cl, nil
}
*/
