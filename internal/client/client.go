package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	server string
	apiKey string
	tenant string
	http   *http.Client
}

type Response struct {
	Success bool            `json:"success"`
	Data    json.RawMessage `json:"data,omitempty"`
	Error   string          `json:"error,omitempty"`
}

func New(server, apiKey, tenant string) *Client {
	return &Client{
		server: server,
		apiKey: apiKey,
		tenant: tenant,
		http:   &http.Client{Timeout: 30 * time.Second},
	}
}

func (c *Client) Call(tool string, args map[string]any) (json.RawMessage, error) {
	body, _ := json.Marshal(args)
	req, _ := http.NewRequest("POST", c.server+"/api/v1/cli/"+tool, bytes.NewReader(body))
	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("X-Tenant", c.tenant)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("connection error: %w", err)
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)

	if resp.StatusCode >= 400 {
		var apiErr Response
		json.Unmarshal(respBody, &apiErr)
		return nil, fmt.Errorf("API error (%d): %s", resp.StatusCode, apiErr.Error)
	}

	var apiResp Response
	json.Unmarshal(respBody, &apiResp)
	return apiResp.Data, nil
}
