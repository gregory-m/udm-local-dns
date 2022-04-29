package dnsmasq

import (
	"reflect"
	"testing"

	"github.com/gregory-m/udm-local-dns/api"
)

func TestGen(t *testing.T) {
	type args struct {
		clients  []api.NetworkClient
		networks []api.Network
	}
	tests := []struct {
		name string

		clients      []api.NetworkClient
		networks     []api.Network
		onlyFixedIPs bool

		wantRecords []Record
		wantErr     bool
	}{
		{
			"Generate records",

			[]api.NetworkClient{{Name: "myhost10", NetworkId: "123", IP: "192.168.1.1"}},
			[]api.Network{{ID: "123", DomainName: "mydomain.beer"}},
			false,

			[]Record{{"myhost10", "mydomain.beer", "192.168.1.1"}},
			false,
		},
		{
			"Generate multiple records in same network",

			[]api.NetworkClient{{Name: "myhost", NetworkId: "123", IP: "192.168.1.1"}, {Name: "myhost2", NetworkId: "123", IP: "192.168.1.2"}},
			[]api.Network{{ID: "123", DomainName: "mydomain.beer"}},
			false,

			[]Record{{"myhost", "mydomain.beer", "192.168.1.1"}, {"myhost2", "mydomain.beer", "192.168.1.2"}},
			false,
		},
		{
			"Generate multiple records in different networks",

			[]api.NetworkClient{{Name: "myhost", NetworkId: "123", IP: "192.168.1.1"}, {Name: "myhost2", NetworkId: "124", IP: "192.168.2.1"}},
			[]api.Network{{ID: "123", DomainName: "mydomain.beer"}, {ID: "124", DomainName: "mydomain.com"}},
			false,

			[]Record{{"myhost", "mydomain.beer", "192.168.1.1"}, {"myhost2", "mydomain.com", "192.168.2.1"}},
			false,
		},
		{
			"Prefer name to hostname",

			[]api.NetworkClient{{Name: "myhost-renamed", Hostname: "myhost", NetworkId: "123", IP: "192.168.1.1"}},
			[]api.Network{{ID: "123", DomainName: "mydomain.beer"}},
			false,

			[]Record{{"myhost-renamed", "mydomain.beer", "192.168.1.1"}},
			false,
		},
		{
			"Use hostname if no name",

			[]api.NetworkClient{{Hostname: "myhost", NetworkId: "123", IP: "192.168.1.1"}},
			[]api.Network{{ID: "123", DomainName: "mydomain.beer"}},
			false,

			[]Record{{"myhost", "mydomain.beer", "192.168.1.1"}},
			false,
		},
		{
			"Filter clients without both hostname and name",

			[]api.NetworkClient{{Hostname: "", Name: "", NetworkId: "123", IP: "192.168.1.11"}, {Name: "myhost", NetworkId: "123", IP: "192.168.1.1"}},
			[]api.Network{{ID: "123", DomainName: "mydomain.beer"}},
			false,

			[]Record{{"myhost", "mydomain.beer", "192.168.1.1"}},
			false,
		},
		{
			"Escape name",

			[]api.NetworkClient{{Name: "Gregory’s iPhone", NetworkId: "123", IP: "192.168.1.1"}},
			[]api.Network{{ID: "123", DomainName: "mydomain.beer"}},
			false,

			[]Record{{"Gregorys-iPhone", "mydomain.beer", "192.168.1.1"}},
			false,
		},
		{
			"Escape name 2",

			[]api.NetworkClient{{Name: "Gregory!$#!", NetworkId: "123", IP: "192.168.1.1"}},
			[]api.Network{{ID: "123", DomainName: "mydomain.beer"}},
			false,

			[]Record{{"Gregory", "mydomain.beer", "192.168.1.1"}},
			false,
		},
		{
			"Only fixed ips",

			[]api.NetworkClient{{Name: "myhost", NetworkId: "123", UseFixedIP: true, IP: "192.168.1.1"}, {Name: "myhost2", NetworkId: "123", UseFixedIP: false, IP: "192.168.1.2"}},
			[]api.Network{{ID: "123", DomainName: "mydomain.beer"}},
			true,

			[]Record{{"myhost", "mydomain.beer", "192.168.1.1"}},
			false,
		},
		{
			"Return error if notwork don't exists",

			[]api.NetworkClient{{Name: "myhost", NetworkId: "bad-id", IP: "192.168.1.1"}},
			[]api.Network{{ID: "123", DomainName: "mydomain.beer"}},
			false,

			[]Record{},
			true,
		},
		{
			"Records in alpha numeric order",

			[]api.NetworkClient{{Name: "myhost", NetworkId: "123", IP: "192.168.1.2"}, {Name: "1myhost", NetworkId: "123", IP: "192.168.1.1"}},
			[]api.Network{{ID: "123", DomainName: "mydomain.beer"}},
			false,

			[]Record{{"1myhost", "mydomain.beer", "192.168.1.1"}, {"myhost", "mydomain.beer", "192.168.1.2"}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRecords, err := Gen(tt.clients, tt.networks, tt.onlyFixedIPs)
			if (err != nil) != tt.wantErr {
				t.Errorf("Gen() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRecords, tt.wantRecords) {
				t.Errorf("Gen() = %v, want %v", gotRecords, tt.wantRecords)
			}
		})
	}
}
