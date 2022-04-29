package api

import (
	"testing"
)

func TestGetUsers(t *testing.T) {
	ts := newTS("testdata/list_user.http")
	defer ts.Close()

	c, err := New("User", "Pass", ts.URL, "default")
	if err != nil {
		t.Fatalf("can't crate API client %s", err)
	}

	users, err := c.GetUsers()
	if err != nil {
		t.Fatalf("Can't get users: %s", err)
	}

	if len(users) != 8 {
		t.Errorf("Wrong len, expected 8, got %d", len(users))
	}

}
