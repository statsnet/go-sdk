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
		"https://dev.statsnet.co/api/v2",
	}, nil
}

func (c *Client) request(ctx context.Context, method, endpoint string, params map[string]string, body io.Reader, isJsonBody bool) ([]byte, error) {
	fullUrl, err := url.JoinPath(c.baseURL, endpoint)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, method, fullUrl, body)
	req.Header.Set("X-API-KEY", c.apiKey)
	if isJsonBody {
		req.Header.Set("Content-Type", "application/json")
	}
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

func (c *Client) get(ctx context.Context, endpoint string, params map[string]string) ([]byte, error) {
	return c.request(ctx, "GET", endpoint, params, nil, false)
}

func (c *Client) post(ctx context.Context, endpoint string, params map[string]string, body io.Reader, isJsonBody bool) ([]byte, error) {
	return c.request(ctx, "POST", endpoint, params, body, isJsonBody)
}

func (c *Client) Me(ctx context.Context) (*MeResponse, error) {
	r, err := c.get(ctx, "/user/me", nil)
	if err != nil {
		return nil, err
	}

	var model MeResponse
	err = json.Unmarshal(r, &model)
	if err != nil {
		return nil, err
	}
	return &model, nil
}

func (c *Client) Search(ctx context.Context, query string, jurisdiction *string, limitRef *int) (*SearchCompanyResponse, error) {
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
	r := bytes.NewBuffer(b)
	response, err := c.post(ctx, "/business/search", params, r, true)
	if err != nil {
		return nil, err
	}
	var model SearchCompanyResponse
	err = json.Unmarshal(response, &model)
	if err != nil {
		return nil, err
	}
	return &model, nil
}

func (c *Client) GetCompany(ctx context.Context, jurisdiction string, id int) (*GetCompanyResponse, error) {
	if err := validateJurisdiction(false, &jurisdiction); err != nil {
		return nil, err
	}
	response, err := c.get(ctx, fmt.Sprintf("/business/%s/%d/paid", jurisdiction, id), nil)
	if err != nil {
		return nil, err
	}
	var model GetCompanyResponse
	err = json.Unmarshal(response, &model)
	if err != nil {
		return nil, err
	}
	return &model, nil
}
func (c *Client) GetCompanyMeta(ctx context.Context, id int) (*GetCompanyMetaResponse, error) {
	response, err := c.get(ctx, fmt.Sprintf("/business/%d/data/view/meta", id), nil)
	if err != nil {
		return nil, err
	}
	var model GetCompanyMetaResponse
	err = json.Unmarshal(response, &model)
	if err != nil {
		return nil, err
	}
	return &model, err
}
func (c *Client) GetCompanyCourtCases(ctx context.Context, id int, limitRef *int) (*GetCompanyCourtCasesResponse, error) {
	limit := normalizeLimit(limitRef)
	if err := validateLimit(limit); err != nil {
		return nil, err
	}
	params := map[string]string{
		"limit": fmt.Sprintf("%d", limit),
	}
	response, err := c.get(ctx, fmt.Sprintf("/business/%d/data/view/meta", id), params)
	if err != nil {
		return nil, err
	}
	var model GetCompanyCourtCasesResponse
	err = json.Unmarshal(response, &model)
	if err != nil {
		return nil, err
	}
	return &model, err
}
func (c *Client) GetCompanyDepartments(ctx context.Context, id int, limitRef *int) (*GetCompanyDepartmentsResponse, error) {
	limit := normalizeLimit(limitRef)
	if err := validateLimit(limit); err != nil {
		return nil, err
	}
	params := map[string]string{
		"limit": fmt.Sprintf("%d", limit),
	}
	response, err := c.get(ctx, fmt.Sprintf("/business/%d/department", id), params)
	if err != nil {
		return nil, err
	}
	var model GetCompanyDepartmentsResponse
	err = json.Unmarshal(response, &model)
	if err != nil {
		return nil, err
	}
	return &model, nil
}
func (c *Client) GetCompanyGovContracts(ctx context.Context, id int, limitRef *int) (*GetCompanyGovContractsResponse, error) {
	limit := normalizeLimit(limitRef)
	if err := validateLimit(limit); err != nil {
		return nil, err
	}
	params := map[string]string{
		"limit": fmt.Sprintf("%d", limit),
	}
	response, err := c.get(ctx, fmt.Sprintf("/business/%d/gov_contracts", id), params)
	if err != nil {
		return nil, err
	}
	var model GetCompanyGovContractsResponse
	err = json.Unmarshal(response, &model)
	if err != nil {
		return nil, err
	}
	return &model, nil
}
func (c *Client) GetCompanyEvents(ctx context.Context, id int, limitRef *int) (*GetCompanyEventsResponse, error) {
	limit := normalizeLimit(limitRef)
	if err := validateLimit(limit); err != nil {
		return nil, err
	}
	params := map[string]string{
		"limit": fmt.Sprintf("%d", limit),
	}
	response, err := c.get(ctx, fmt.Sprintf("/business/%d/events", id), params)
	if err != nil {
		return nil, err
	}
	var model GetCompanyEventsResponse
	err = json.Unmarshal(response, &model)
	if err != nil {
		return nil, err
	}
	return &model, err
}
func (c *Client) GetCompanyRelations(ctx context.Context, id int, limitRef *int) (*GetCompanyRelationsResponse, error) {
	limit := normalizeLimit(limitRef)
	if err := validateLimit(limit); err != nil {
		return nil, err
	}
	params := map[string]string{
		"limit": fmt.Sprintf("%d", limit),
	}
	response, err := c.get(ctx, fmt.Sprintf("/business/%d/relations/table", id), params)
	if err != nil {
		return nil, err
	}
	var model GetCompanyRelationsResponse
	err = json.Unmarshal(response, &model)
	if err != nil {
		return nil, err
	}
	return &model, err
}

func (c *Client) GetCompanyByIdentifier(ctx context.Context, identifier string) (*GetCompanyByIdentifierResponse, error) {
	response, err := c.get(ctx, fmt.Sprintf("/business/%s/paid", identifier), nil)
	if err != nil {
		return nil, err
	}
	var model GetCompanyByIdentifierResponse
	err = json.Unmarshal(response, &model)
	if err != nil {
		return nil, err
	}
	return &model, nil
}
