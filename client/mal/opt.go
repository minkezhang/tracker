package mal

import (
	"net/http"

	"github.com/minkezhang/truffle/client"
)

type O struct {
	transport http.RoundTripper
	auth      client.AuthType
}

func WithPublicAPIKey(client_id string) O {
	return O{
		transport: t{
			cid: client_id,
		},
		auth: client.AuthTypePublic,
	}
}

type t struct {
	cid string
}

func (t t) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("X-MAL-CLIENT-ID", t.cid)
	return http.DefaultTransport.RoundTrip(req)
}
