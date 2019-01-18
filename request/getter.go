package request

import (
	"net/http"
)

// Getter is a requester doing GET
type Getter struct {
	*Requester
}

// Request do actual request
func (g *Getter) Request(url, body string) (*http.Response, error) {
	url = g.EscapeURL(url)

	r, e := http.NewRequest("GET", url, nil)
	if e != nil {
		return nil, e
	}
	for k, v := range g.Requester.headers {
		r.Header.Set(k, v)
	}
	for _, cookie := range g.Requester.cookies {
		r.AddCookie(cookie)
	}

	g.Requester.LogRequest(r)

	return g.Requester.client.Do(r)
}
