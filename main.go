package sharesession

import (
	"fmt"
	"github.com/go-resty/resty/v2"
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
