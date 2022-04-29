package api

import (
	"reflect"
	"testing"
)

func TestGetClients(t *testing.T) {
	ts := newTS("testdata/sta.http")
	defer ts.Close()

	c, err := New("User", "Pass", ts.URL, "default")
	if err != nil {
		t.Fatalf("can't crate API client %s", err)
	}

	clients, err := c.GetClients()
	if err != nil {
		t.Fatalf("Can't get clients: %s", err)
	}

	expected := []NetworkClient{
		{ID: "624c4e12ece3e2089f8c7de2", Name: "", Hostname: "", IP: "192.168.1.159", UseFixedIP: false, NetworkId: "6228f3b2f7412a097be95b06"},
		{ID: "624f1ebfece3e2089f8cae32", Name: "", Hostname: "Gregorys-iPhone", IP: "192.168.1.218", UseFixedIP: false, NetworkId: "6228f3b2f7412a097be95b06"},
		{ID: "6235e82cc01e440a5c5ac076", Name: "", Hostname: "brewery", IP: "192.168.1.109", UseFixedIP: true, NetworkId: "6228f3b2f7412a097be95b06"},
		{ID: "6228f8cdf7412a69e0e31b80", Name: "Gregorys-MBP-Renamed", Hostname: "Gregorys-MBP", IP: "192.168.1.54", UseFixedIP: false, NetworkId: "6228f3b2f7412a097be95b06"},
	}

	if !reflect.DeepEqual(clients, expected) {
		t.Fatalf("Got wrong result:\nExp: %+v\n\nGot: %+v", expected, clients)
	}

}
