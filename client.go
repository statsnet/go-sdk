package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

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

func (c *Client) request(ctx context.Context, method, endpoint string, params map[string]string, body io.Reader) ([]byte, error) {
	fullUrl, err := url.JoinPath(c.baseURL, endpoint)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, method, fullUrl, body)
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
		return nil, NewClientError(endpoint, resp.StatusCode, buf.Bytes())
	}

	return buf.Bytes(), nil
}

func (c *Client) get(ctx context.Context, endpoint string, params map[string]string, body io.Reader) ([]byte, error) {
	return c.request(ctx, "GET", endpoint, params, body)
}

func (c *Client) post(ctx context.Context, endpoint string, params map[string]string, body io.Reader) ([]byte, error) {
	return c.request(ctx, "POST", endpoint, params, body)
}

func (c *Client) Me(ctx context.Context) ([]byte, error) {
	return c.get(ctx, "/user/me", nil, nil)
}

func (c *Client) Search(ctx context.Context, query string, jurisdiction *string, limitRef *int) ([]byte, error) {
	var limit int
	if limitRef == nil {
		limit = 5
	} else {
		limit = *limitRef
	}
	if err := validateLimit(limit); err != nil {
		return nil, err
	}

	if err := validateJurisdiction(true, jurisdiction); err != nil {
		return nil, err
	}

	params := map[string]string{"limit": fmt.Sprintf("%d", limit)}

	type body struct {
		Filters struct {
			Jurisdiction []string `json:"jurisdiction"`
		} `json:"filters"`
		Query string `json:"query"`
	}
	var filters struct {
		Jurisdiction []string `json:"jurisdiction"`
	}
	if jurisdiction != nil {
		filters = struct {
			Jurisdiction []string `json:"jurisdiction"`
		}{
			Jurisdiction: []string{*jurisdiction},
		}
	}

	b, _ := json.Marshal(body{
		Filters: filters,
		Query:   query,
	})
	r := bytes.NewReader(b)
	return c.post(ctx, "/business/search", params, r)
}

func (c *Client) GetCompany(ctx context.Context, jurisdiction string, id int) ([]byte, error) {
	validateJurisdiction(false, &jurisdiction)
	return c.get(ctx, fmt.Sprintf("/business/%s/%d/paid", jurisdiction, id), nil, nil)
}
func (c *Client) GetCompanyMeta(ctx context.Context, id int) ([]byte, error) {
	return c.get(ctx, fmt.Sprintf("/business/%d/data/view/meta", id), nil, nil)
}
func (c *Client) GetCompanyCourtCases(ctx context.Context, id, limitRef *int) ([]byte, error) {
	limit := normalizeLimit(limitRef)
	validateLimit(limit)
	params := map[string]string{
		"limit": fmt.Sprintf("%d", limit),
	}
	return c.get(ctx, fmt.Sprintf("/business/%d/data/view/meta", id), params, nil)
}
func (c *Client) GetCompanyDepartments(ctx context.Context, id, limitRef *int) ([]byte, error) {
	limit := normalizeLimit(limitRef)
	validateLimit(limit)
	params := map[string]string{
		"limit": fmt.Sprintf("%d", limit),
	}
	return c.get(ctx, fmt.Sprintf("/business/%d/department", id), params, nil)
}
func (c *Client) GetCompanyGovContracts(ctx context.Context, id, limitRef *int) ([]byte, error) {
	limit := normalizeLimit(limitRef)
	validateLimit(limit)
	params := map[string]string{
		"limit": fmt.Sprintf("%d", limit),
	}
	return c.get(ctx, fmt.Sprintf("/business/%d/gov_contracts", id), params, nil)
}
func (c *Client) GetCompanyEvents(ctx context.Context, id, limitRef *int) ([]byte, error) {
	limit := normalizeLimit(limitRef)
	validateLimit(limit)
	params := map[string]string{
		"limit": fmt.Sprintf("%d", limit),
	}
	return c.get(ctx, fmt.Sprintf("/business/%d/events", id), params, nil)
}
