package cve

import (
	"awesomeProject/database"
	"awesomeProject/http"
	"fmt"
)

// CVE_2017_8917 检查给定 URL 是否存在 CVE_2017_8917 漏洞
func CVE_2017_8917(url string) (string, error) {
	payload, err := database.QueryPoc("CVE_2017_8917")
	if err != nil {
		return "", fmt.Errorf("查询payload失败: %v", err)
	}

	rest, statusCode, err := http.SendGetRequest(url + payload)
	if err != nil {
		return "", fmt.Errorf("发送payload失败: %v", err)
	}

	message := "不存在漏洞!"
	if statusCode == 200 {
		message = "存在漏洞!"
	}

	fmt.Println(url, message) // 打印 URL 和消息

	return rest, nil
}
