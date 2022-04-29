package api

import (
	"reflect"
	"testing"
)

func TestGetNetworks(t *testing.T) {
	ts := newTS("testdata/networkconf.http")
	defer ts.Close()

	c, err := New("User", "Pass", ts.URL, "default")
	if err != nil {
		t.Fatalf("can't crate API client %s", err)
	}

	clients, err := c.GetNetworks()
	if err != nil {
		t.Fatalf("Can't get clients: %s", err)
	}

	expected := []Network{
		{ID: "6228f3b2f7412a097be95b04", Name: "Default (WAN1)", DomainName: ""},
		{ID: "6228f3b2f7412a097be95b05", Name: "Backup (WAN2)", DomainName: ""},
		{ID: "6228f3b2f7412a097be95b06", Name: "Default", DomainName: "man-family.com"},
	}

	if !reflect.DeepEqual(clients, expected) {
		t.Fatalf("Got wrong result:\nExp: %+v\n\nGot: %+v", expected, clients)
	}

}
