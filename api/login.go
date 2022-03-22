package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) Login() error {
	values := map[string]string{"username": c.Username, "password": c.Password}
	jsonValue, err := json.Marshal(values)
	if err != nil {
		return fmt.Errorf("can't marshal json during login: %w", err)
	}

	url := c.URL + "/api/auth/login" // TODO use net/url.JoinPath() after go 1.19 released https://github.com/golang/go/issues/47005

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
	if err != nil {
		return fmt.Errorf("can't create request during login: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("can't send request during login: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("can't login got %d status code from server", resp.StatusCode)
	}

	return nil
}
