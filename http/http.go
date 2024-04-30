package http

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

// SendGetRequest 发送 GET 请求并返回响应体、状态码和可能的错误
func SendGetRequest(url string) (string, int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", http.StatusInternalServerError, fmt.Errorf("发送请求失败: %v", err)
	}
	defer resp.Body.Close()

	var body bytes.Buffer
	_, err = io.Copy(&body, resp.Body)
	if err != nil {
		return "", resp.StatusCode, fmt.Errorf("读取响应体失败: %v", err)
	}

	return body.String(), resp.StatusCode, nil
}
