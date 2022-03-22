package dnsmasq

import "testing"

func TestMarshal(t *testing.T) {

	tests := []struct {
		name    string
		records []Record
		want    string
	}{
		{
			"Single record",
			[]Record{{"myhost", "mydomain.beer", "192.168.1.1"}},
			"host-record=myhost.mydomain.beer,myhost,192.168.1.1",
		},
		{
			"multiple records",
			[]Record{{"myhost", "mydomain.beer", "192.168.1.1"}, {"myhost2", "mydomain.beer", "192.168.1.2"}},
			"host-record=myhost.mydomain.beer,myhost,192.168.1.1\nhost-record=myhost2.mydomain.beer,myhost2,192.168.1.2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Marshal(tt.records); got != tt.want {
				t.Errorf("Marshal() = %v, want %v", got, tt.want)
			}
		})
	}
}
