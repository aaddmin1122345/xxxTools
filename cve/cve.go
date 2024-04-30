// cve.go

package cve

import (
	"awesomeProject/database"
	"awesomeProject/http"
	"fmt"
)

func CVE_2017_8917(url string) (string, error) {
	payload, err := database.QueryPoc("CVE_2017_8917")
	if err != nil {
		return "查询payload失败!", err
	}
	// 发送 payload
	rest, statusCode, err := http.SendGetRequest(url + payload)
	//fmt.Println(statusCode)
	if err != nil {
		return "发送payload失败!", err
	}
	//var message  string
	if statusCode == 200 {
		message := "存在漏洞!"
		fmt.Println(url + " " + message)

	} else {
		message := "不存在漏洞!"
		fmt.Println(url + " " + message)
	}
	//if err = fyneGui.TextUpdater.UpdateText(message); err != nil {
	//	return "更新 GUI 文本失败!", err
	//}
	return rest, nil
}
