package request

import (
	"testing"
)

func TestNew(t *testing.T) {
	SetGlobalOption(
		WithBasicAuth("test", "123"),
	)

	req := New(WithBasicAuth("test1", "p123"))

	if globalRequest.username == req.username {
		t.Errorf("global request username was %s, but %s", globalRequest.username, req.username)
	}

	if globalRequest.password == req.password {
		t.Errorf("global request password was %s, but %s", globalRequest.password, req.password)
	}
}
