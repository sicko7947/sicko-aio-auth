package psychoclient

import (
	"net/http"

	utls "github.com/refraction-networking/utls"
	"golang.org/x/net/proxy"
)

// clientConfig : Client hello config
type clientConfig struct {
	Host            string
	Certificate     string
	Privatekey      string
	FollowRedirects bool
	ClientHelloID   utls.ClientHelloID
	ClientHelloSpec utls.ClientHelloSpec
}

// newClient : creates a new base psycho http client
func newClient(config *clientConfig, proxyURL ...string) (http.Client, error) {
	var dialer proxy.ContextDialer = proxy.Direct
	if len(proxyURL) > 0 && len(proxyURL[0]) > 0 {
		d, err := newConnectDialer(proxyURL[0])
		if err != nil {
			return http.Client{}, err
		}
		dialer = d
	}

	redirectCallback := noRedirects
	if config.FollowRedirects {
		// Reset to default policy (redirect up to 10 times)
		redirectCallback = nil
	}

	return http.Client{
		Transport: newRoundTripper(&RoundTripper{
			Host:            config.Host,
			Certificate:     config.Certificate,
			Privatekey:      config.Privatekey,
			ClientHelloID:   config.ClientHelloID,
			ClientHelloSpec: config.ClientHelloSpec,
		}, dialer),
		CheckRedirect: redirectCallback,
	}, nil
}

func noRedirects(req *http.Request, via []*http.Request) error {
	return http.ErrUseLastResponse
}
