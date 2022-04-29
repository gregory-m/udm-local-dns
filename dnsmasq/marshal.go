package dnsmasq

import (
	"fmt"
	"strings"
)

func Marshal(records []Record) string {
	out := []string{}

	for _, r := range records {
		line := fmt.Sprintf("host-record=%s.%s,%s,%s", r.Hostname, r.DomainName, r.Hostname, r.IP)
		out = append(out, line)
	}

	return strings.Join(out, "\n")
}
