package mal

import (
	"net/http"

	"github.com/minkezhang/truffle/client"
	"github.com/minkezhang/truffle/util"
)

type O struct {
	transport http.RoundTripper
	auth      client.AuthType
	config    util.Config
}

func FromConfig(c util.Config) O {
	return O{
		transport: t{
			cid: c.MAL.ClientID,
		},
		auth:   client.AuthTypePublic,
		config: c,
	}
}

type t struct {
	cid string
}

func (t t) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("X-MAL-CLIENT-ID", t.cid)
	return http.DefaultTransport.RoundTrip(req)
}
