package http

import (
	"bytes"
	"io"
	"net/http"
)

func SendGetRequest(url string) (string, int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "获取body失败!", resp.StatusCode, nil
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
		return "打印body失败!", '_', nil
	}
	//fmt.Println(resp.StatusCode)

	return body.String(), resp.StatusCode, nil

}
