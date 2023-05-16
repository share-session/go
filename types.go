package sharesession

type Cookies map[string]interface{}
type LocalStorage map[string]interface{}
type SessionStorage map[string]interface{}

type Session struct {
	Id             string          `json:"id"`
	Cookies        *Cookies        `json:"cookies"`
	LocalStorage   *LocalStorage   `json:"localStorage"`
	SessionStorage *SessionStorage `json:"sessionStorage"`
}
