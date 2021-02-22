package robot

import (
	"fmt"
	"net/http"
	"strings"
)

type Message interface {
	Send(webHook string) SendResult
}

type SendResult struct {
	StatusCode int
}

func send(url string, content string) SendResult {
	result := SendResult{}
	resp, err := http.Post(url, "application/json; charset=utf-8", strings.NewReader(content))
	if err != nil {
		fmt.Println(err)
		return result
	}
	_ = resp.Body.Close()
	result.StatusCode = resp.StatusCode
	return result
}
