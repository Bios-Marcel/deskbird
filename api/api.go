package api

import (
	"net/http"
)

var (
	BaseURL    string = "https://app.deskbird.com/api"
	APIVersion        = "v1.1"
)

type API struct {
	bearerToken string
}

func New(bearerToken string) *API {
	return &API{bearerToken: bearerToken}
}

func injectDefaultHeaders(request *http.Request) {
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Accept", "application/json, text/plain, */*")
	request.Header.Add("Accept-Language", "en-US,en;q=0.5")
	// FIXME We ain't uncompressing rn:
	// request.Header.Add("Accept-Encoding", "gzip, deflate, br")

	// User-Agent doesn't seem to relevant for their backend, but might be smart to specify anyway.
	// request.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:104.0) Gecko/20100101 Firefox/104.0")

	// Stuff i am using to try'n get the request through:
	request.Header.Add("Host", "app.deskbird.com")
	request.Header.Add("Origin", "https://app.deskbird.com")
	request.Header.Add("DNT", "1")
	request.Header.Add("Sec-Fetch-Dest", "empty")
	request.Header.Add("Sec-Fetch-Mode", "cors")
	request.Header.Add("Sec-Fetch-Site", "same-origin")
	request.Header.Add("TE", "trailers")

}
