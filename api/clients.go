package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ClientResp struct {
	Data []NetworkClient `json:"data"`
}

type NetworkClient struct {
	ID         string `json:"_id"`
	Name       string `json:"name"`
	Hostname   string `json:"hostname"`
	IP         string `json:"ip"`
	UseFixedIP bool   `json:"use_fixedip"`
	NetworkId  string `json:"network_id"`
}

func (c *Client) GetClients() (clients []NetworkClient, err error) {
	url := c.URL + fmt.Sprintf("/proxy/network/api/s/%s/stat/sta", c.Site) // TODO use net/url.JoinPath() after go 1.19 released https://github.com/golang/go/issues/47005

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return clients, fmt.Errorf("can't create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return clients, fmt.Errorf("can't send request: %w", err)
	}
	defer resp.Body.Close()

	// dump, err := httputil.DumpResponse(resp, true)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%s", dump)
	// f, err := os.Create("data.txt")
	// f.Write(dump)
	// f.Close()

	if resp.StatusCode != http.StatusOK {
		return clients, fmt.Errorf("error got %d status code from server", resp.StatusCode)
	}

	decoded := new(ClientResp)
	err = json.NewDecoder(resp.Body).Decode(decoded)
	if err != nil {
		return clients, fmt.Errorf("can't parse JSON during: %w", err)
	}

	return decoded.Data, err

}
