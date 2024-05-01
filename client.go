package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

type ClientException struct {
	Endpoint        string
	StatusCode      int
	ResponseContent []byte
}

func NewClientException(endpoint string, statusCode int, responseContent []byte) *ClientException {
	return &ClientException{
		Endpoint:        endpoint,
		StatusCode:      statusCode,
		ResponseContent: responseContent,
	}
}

func (c *ClientException) Error() string {
	return fmt.Sprintf("Request to %s unsuccessful. Status code: %d. Response content: %s", c.Endpoint, c.StatusCode, c.ResponseContent)
}

type Client struct {
	client  *http.Client
	apiKey  string
	baseURL string
}

func NewClient(apiKey string) (*Client, error) {
	if apiKey == "" {
		apiKey = os.Getenv("STATSNET_API_KEY")
	}
	if apiKey == "" {
		return nil, errors.New("no api key provided. Set via argument or environment variable STATSNET_API_KEY")
	}
	return &Client{
		http.DefaultClient,
		apiKey,
		"https://statsnet.co/api/v2",
	}, nil
}

func (c *Client) request(method, endpoint string, params map[string]string, body io.Reader) ([]byte, error) {
	fullUrl, err := url.JoinPath(c.baseURL, endpoint)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(method, fullUrl, body)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	for key, value := range params {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	if resp.StatusCode > 299 {
		return nil, NewClientException(endpoint, resp.StatusCode, buf.Bytes())
	}

	return buf.Bytes(), nil
}

func (c *Client) get(endpoint string, params map[string]string, body io.Reader) ([]byte, error) {
	return c.request("GET", endpoint, params, body)
}

func (c *Client) post(endpoint string, params map[string]string, body io.Reader) ([]byte, error) {
	return c.request("POST", endpoint, params, body)
}
