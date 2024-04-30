package cve

import (
	"awesomeProject/database"
	"awesomeProject/http"
	"fmt"
)

//var payload string
//var err error
//var request string

func CVE_2017_8917(url string) (string, error) {
	payload, err := database.QueryPoc(database.Db, "CVE_2017_8917")
	fmt.Println(payload)
	if err != nil {
		return "", err
	}
	request, err := http.SendGetRequest(url + payload)
	if err != nil {
		return "", err
	}
	return request, nil
}
