package sharesession

type Cookie struct {
	Name     string `json:"name"`
	Value    string `json:"value"`
	Path     string `json:"path"`
	Domain   string `json:"domain"`
	Secure   bool   `json:"secure"`
	HttpOnly bool   `json:"httpOnly"`
}

type LocalStorage string
type SessionStorage string

type Session struct {
	Id             string                    `json:"id"`
	Cookies        map[string]Cookie         `json:"cookies"`
	LocalStorage   map[string]LocalStorage   `json:"localStorage"`
	SessionStorage map[string]SessionStorage `json:"sessionStorage"`
	HTML           string                    `json:"html"`
}
