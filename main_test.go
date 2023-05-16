package sharesession

import (
	"testing"
)

func TestCreate(t *testing.T) {
	create := &Session{
		Cookies: &Cookies{
			"cookie-name": "cookie-value",
		},
	}

	session, err := Create(create)

	if err != nil {
		t.Error(err)
	}

	if session == nil {
		t.Errorf("Failed to create new session")
	}
}

func TestGet(t *testing.T) {
	session, err := Get("123456")

	if err != nil {
		t.Error(err)
	}

	t.Log(session)

	if session == nil {
		t.Errorf("Failed to create new session")
	}
}
