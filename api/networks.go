package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type NetworkResp struct {
	Data []Network `json:"data"`
}

type Network struct {
	ID         string `json:"_id"`
	Name       string `json:"name"`
	DomainName string `json:"domain_name"`
}

func (c *Client) GetNetworks() (networks []Network, err error) {
	url := c.URL + fmt.Sprintf("/proxy/network/api/s/%s/rest/networkconf", c.Site) // TODO use net/url.JoinPath() after go 1.19 released https://github.com/golang/go/issues/47005

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return networks, fmt.Errorf("can't create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return networks, fmt.Errorf("can't send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return networks, fmt.Errorf("error got %d status code from server", resp.StatusCode)
	}

	decoded := new(NetworkResp)
	err = json.NewDecoder(resp.Body).Decode(decoded)
	if err != nil {
		return networks, fmt.Errorf("can't parse JSON during: %w", err)
	}

	return decoded.Data, err

}
