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
		return GlobalClient.Post(url, body, headers)
	case "PUT":
		return GlobalClient.Request("PUT", url, body, headers)
	case "PATCH":
		return GlobalClient.Request("PATCH", url, body, headers)
	case "DELETE":
		return GlobalClient.Request("DELETE", url, "", headers)
	default:
		return nil, fmt.Errorf("unsupported method: %s", resType)
	}
}
