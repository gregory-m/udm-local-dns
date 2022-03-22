package dnsmasq

import (
	"fmt"
	"sort"
	"strings"
	"unicode"

	"github.com/gregory-m/udm-se-dns/api"
)

type Record struct {
	Hostname   string
	DomainName string
	IP         string
}

func Gen(clients []api.NetworkClient, networks []api.Network, onlyFixedIPs bool) (records []Record, err error) {
	records = []Record{}
	for _, c := range clients {
		if onlyFixedIPs && c.UseFixedIP != true {
			continue
		}
		n, err := networkByID(networks, c.NetworkId)
		if err != nil {
			return records, err
		}

		hostname := ""
		if c.Name != "" {
			hostname = dnsEscape(c.Name)
		} else if c.Hostname != "" {
			hostname = c.Hostname
		} else {
			continue
		}

		records = append(records, Record{Hostname: hostname, DomainName: n.DomainName, IP: c.IP})
	}

	sort.Slice(records, func(i, j int) bool {
		return records[i].Hostname < records[j].Hostname
	})

	return records, err
}

func networkByID(networks []api.Network, id string) (network api.Network, err error) {
	for _, n := range networks {
		if n.ID == id {
			return n, nil
		}
	}

	return network, fmt.Errorf("can't find networks with id %s", id)
}

func dnsEscape(s string) string {
	res := strings.Map(func(r rune) rune {
		if unicode.Is(unicode.Space, r) {
			return '-'
		}

		if ('a' <= r && r <= 'z') || ('A' <= r && r <= 'Z') {
			return r
		}

		if '0' <= r && r <= '9' {
			return r
		}

		if r == '-' {
			return r
		}

		return -1
	}, s)

	return res
}
