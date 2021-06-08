package psychoclient

import (
	"net/http"
	"net/url"
)

// newDefaultClient : creates a new default Client
func newDefaultClient(config *clientConfig, proxyURL ...string) (http.Client, error) {
	var proxy *url.URL
	if len(proxyURL) > 0 && len(proxyURL[0]) > 0 {
		proxy, _ = url.Parse(proxyURL[0])
	} else {
		return http.Client{
			Transport: &http.Transport{},
		}, nil
	}

	redirectCallback := noRedirects
	if config.FollowRedirects {
		// Reset to default policy (redirect up to 10 times)
		redirectCallback = nil
	}

	return http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxy),
		},
		CheckRedirect: redirectCallback,
	}, nil
}
