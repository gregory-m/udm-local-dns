package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type UserResp struct {
	Data []User `json:"data"`
}

type User struct {
	Hostname   string `json:"hostname"`
	FixedIP    string `json:"fixed_ip"`
	UseFixedIP bool   `json:"use_fixedip"`
}

func (c *Client) GetUsers() (users []User, err error) {
	url := c.URL + fmt.Sprintf("/proxy/network/api/s/%s/list/user", c.Site) // TODO use net/url.JoinPath() after go 1.19 released https://github.com/golang/go/issues/47005

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return users, fmt.Errorf("can't create request during getUsers: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return users, fmt.Errorf("can't send request during getUsers: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return users, fmt.Errorf("can't getUsers got %d status code from server", resp.StatusCode)
	}

	decoded := new(UserResp)
	err = json.NewDecoder(resp.Body).Decode(decoded)
	if err != nil {
		return users, fmt.Errorf("can't parse JSON during getUsers: %w", err)
	}

	return decoded.Data, err
}
