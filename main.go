package sharesession

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/http"
)

func Create(session *Session) (*Session, error) {
	client := resty.New().SetBaseURL("https://functions.share-session.com/api")

	var response Session

	_, err := client.
		R().
		EnableTrace().
		SetBody(session).
		SetResult(&response).
		Post("s")

	if err != nil {
		return nil, err
	}

	return &response, nil
}

func New(session *Session) (*Session, error) {
	client := resty.New().SetBaseURL("https://functions.share-session.com/api")

	var response Session

	_, err := client.
		R().
		EnableTrace().
		SetBody(session).
		SetResult(&response).
		Post("s")

	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (session *Session) SetCookie(name string, value Cookie) *Session {
	if session.Cookies == nil {
		session.Cookies = map[string]Cookie{}
	}

	session.Cookies[name] = value

	return session
}

func (session *Session) SetCookies(cookies map[string]Cookie) *Session {
	session.Cookies = cookies

	return session
}

func (session *Session) SetHttpCookie(cookie *http.Cookie) *Session {
	session.SetCookie(cookie.Name, Cookie{
		Name:     cookie.Name,
		Value:    cookie.Value,
		Domain:   cookie.Domain,
		Path:     cookie.Path,
		Secure:   cookie.Secure,
		HttpOnly: cookie.HttpOnly,
	})

	return session
}

func (session *Session) SetHttpCookies(cookies []*http.Cookie) *Session {
	for _, cookie := range cookies {
		session.SetHttpCookie(cookie)
	}

	return session
}

func (session *Session) SetHTML(html string) *Session {
	session.HTML = html

	return session
}

func Get(id string) (*Session, error) {
	client := resty.New().SetBaseURL("https://functions.share-session.com/api")

	var response Session

	_, err := client.
		R().
		EnableTrace().
		SetResult(&response).
		Get(fmt.Sprintf("s/%s", id))

	if err != nil {
		return nil, err
	}

	return &response, nil
}
