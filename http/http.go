package http

import (
	"bytes"
	"io"
	"net/http"
)

//var resp *http.Response
//var err error

func SendGetRequest(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "请求http时发生错误!", err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {

		}
	}(resp.Body)

	// 使用 bytes.Buffer 作为目标，将响应体拷贝到其中
	var body bytes.Buffer
	_, err = io.Copy(&body, resp.Body)
	if err != nil {
		return "", err
	}

	return body.String(), nil
}
