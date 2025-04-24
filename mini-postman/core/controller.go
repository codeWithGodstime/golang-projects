package core

import (
	"fmt"
	"net/http"
)

func MakeRequestController(resType, url string, headers map[string]string, body string) (*http.Response, error) {
	switch resType {
	case "GET":
		return GlobalClient.Get(url, headers)
	case "POST":
		return GlobalClient.Post(url, "", headers)
	case "DELETE":
		// return GlobalClient.Request("DELETE", url, "", headers)
		return nil, nil
	default:
		return nil, fmt.Errorf("unsupported method: %s", resType)
	}
}
