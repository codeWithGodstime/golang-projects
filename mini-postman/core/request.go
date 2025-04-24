package core

import (
	"bytes"
	"io"
	"net/http"
	"time"
)

type Client struct {
	client *http.Client
}


func (r *Client) GET(data: http.Request) (r  http.Response) {

}

// singleton instance
var GlobalClient = &Client{
	client: &http.Client{
		Timeout: 30 * time.Second,
	},
}

func (c *Client) Request(method, url, body string, header map[string]string) {
	var requestBody io.Reader

	if body != "" {
		requestBody = bytes.NewBufferString(body)
	}

	req, err := http.NewRequest(method, url, requestBody)
	if err != nil {
		return nil, err
	}

	// Set headers
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	return c.client.Do(req)
}

func (c *Client) Get(url string, headers map[string]string) (*http.Response, error) {
	return c.Request("GET", url, "", headers)
}

// Convenience POST (with string body)
func (c *Client) Post(url string, body string, headers map[string]string) (*http.Response, error) {
	return c.Request("POST", url, body, headers)
}

// JSON POST (accepts any struct or map)
func (c *Client) PostJSON(url string, payload interface{}, headers map[string]string) (*http.Response, error) {
	jsonBody, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	if headers == nil {
		headers = map[string]string{}
	}
	headers["Content-Type"] = "application/json"
	return c.Request("POST", url, string(jsonBody), headers)
}