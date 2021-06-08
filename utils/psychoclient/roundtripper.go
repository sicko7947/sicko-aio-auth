package psychoclient

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"net"
	"net/http"
	"strings"
	"sync"

	"golang.org/x/net/http2"
	"golang.org/x/net/proxy"

	"github.com/gogf/gf/container/gmap"
	utls "github.com/refraction-networking/utls"
)

var errProtocolNegotiated = errors.New("protocol negotiated")

type RoundTripper roundTripper

type roundTripper struct {
	sync.Mutex

	Host            string
	Certificate     string
	Privatekey      string
	ClientHelloID   utls.ClientHelloID
	ClientHelloSpec utls.ClientHelloSpec

	cachedConnections *gmap.StrAnyMap
	cachedTransports  *gmap.StrAnyMap

	dialer proxy.ContextDialer

	forceCertPinning bool
	pinnedCerts      map[string]bool

	// This is for debugging only, should not be enabled in production
	ignoreCertIssues bool
}

func (rt *roundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	addr := rt.getDialTLSAddr(req)

	roundTripperObj := rt.cachedTransports.Get(addr)
	if roundTripperObj == nil {
		if err := rt.getTransport(req, addr); err != nil {
			return nil, err
		}
		return rt.RoundTrip(req)
	}
	return roundTripperObj.(http.RoundTripper).RoundTrip(req)
}

func (rt *roundTripper) getTransport(req *http.Request, addr string) error {
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM([]byte(rt.Certificate))

	switch strings.ToLower(req.URL.Scheme) {
	case "http":
		rt.cachedTransports.Set(addr, &http.Transport{
			DialContext: rt.dialer.DialContext,
			TLSClientConfig: &tls.Config{
				RootCAs:            caCertPool,
				InsecureSkipVerify: rt.ignoreCertIssues,
			},
		})
		return nil
	case "https":
		rt.cachedTransports.Set(addr, &http2.Transport{
			DialTLS: rt.dialTLSHTTP2,
			TLSClientConfig: &tls.Config{
				RootCAs:            caCertPool,
				InsecureSkipVerify: rt.ignoreCertIssues,
			},
		})
		return nil
	default:
		return fmt.Errorf("invalid URL scheme: [%v]", req.URL.Scheme)
	}

	// _, err := rt.dialTLS(context.Background(), "tcp", addr)
	// switch err {
	// case errProtocolNegotiated:
	// case nil:
	// 	// Should never happen.
	// 	panic("dialTLS returned no error when determining cachedTransports")
	// default:
	// 	return err
	// }

	// return nil
}

func (rt *roundTripper) dialTLS(ctx context.Context, network, addr string) (net.Conn, error) {
	rt.Lock()
	defer rt.Unlock()

	// If we have the connection from when we determined the HTTPS
	// cachedTransports to use, return that.

	if connObj := rt.cachedConnections.Get(addr); connObj != nil {
		rt.cachedConnections.Remove(addr)
		return connObj.(net.Conn), nil
	}

	rawConn, err := rt.dialer.DialContext(ctx, network, addr)
	if err != nil {
		return nil, err
	}

	var host string
	switch {
	case len(rt.Host) > 0:
		host = rt.Host
	default:
		if host, _, err = net.SplitHostPort(addr); err != nil {
			host = addr
		}
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM([]byte(rt.Certificate))
	certs, _ := utls.X509KeyPair([]byte(rt.Certificate), []byte(rt.Privatekey))

	var tlsConn *utls.UConn
	switch rt.ClientHelloID {
	case utls.HelloCustom:
		tlsConn = utls.UClient(rawConn, &utls.Config{NextProtos: []string{"h2", "http/1.1"}, ServerName: host, Certificates: []utls.Certificate{certs}, InsecureSkipVerify: rt.ignoreCertIssues}, utls.HelloCustom)
		if err = tlsConn.ApplyPreset(&rt.ClientHelloSpec); err != nil {
			_ = tlsConn.Close()
			return nil, err
		}
	default:
		tlsConn = utls.UClient(rawConn, &utls.Config{ServerName: host, InsecureSkipVerify: rt.ignoreCertIssues}, rt.ClientHelloID)
	}

	tlsConn.SetSNI(host)
	if err = tlsConn.Handshake(); err != nil {
		_ = tlsConn.Close()
		return nil, err
	}

	if rt.cachedTransports.Get(addr) != nil {
		return tlsConn, nil
	}

	// No http.Transport constructed yet, create one based on the results
	// of ALPN.
	switch tlsConn.ConnectionState().NegotiatedProtocol {
	case http2.NextProtoTLS:
		// The remote peer is speaking HTTP 2 + TLS.
		rt.cachedTransports.Set(addr, &http2.Transport{
			DialTLS: rt.dialTLSHTTP2,
			TLSClientConfig: &tls.Config{
				RootCAs:            caCertPool,
				InsecureSkipVerify: rt.ignoreCertIssues,
			},
		})
	default:
		// Assume the remote peer is speaking HTTP 1.x + TLS.
		rt.cachedTransports.Set(addr, &http.Transport{
			DialTLSContext: rt.dialTLS,
			TLSClientConfig: &tls.Config{
				RootCAs: caCertPool,
			},
		})
	}

	// Stash the connection just established for use servicing the
	// actual request (should be near-immediate).
	rt.cachedConnections.Set(addr, tlsConn)
	return nil, errProtocolNegotiated
}

func (rt *roundTripper) dialTLSHTTP2(network, addr string, _ *tls.Config) (net.Conn, error) {
	return rt.dialTLS(context.Background(), network, addr)
}

func (rt *roundTripper) getDialTLSAddr(req *http.Request) string {
	host, port, err := net.SplitHostPort(req.URL.Host)
	if err == nil {
		return net.JoinHostPort(host, port)
	}
	return net.JoinHostPort(req.URL.Host, "443") // we can assume port is 443 at this point
}

func newRoundTripper(rt *RoundTripper, dialer ...proxy.ContextDialer) http.RoundTripper {
	cachedTransports := gmap.NewStrAnyMap(true)
	cachedConnections := gmap.NewStrAnyMap(true)

	if len(dialer) > 0 {
		return &roundTripper{
			dialer: dialer[0],

			Host:            rt.Host,
			ClientHelloID:   rt.ClientHelloID,
			Certificate:     rt.Certificate,
			Privatekey:      rt.Privatekey,
			ClientHelloSpec: rt.ClientHelloSpec,

			ignoreCertIssues: false,

			cachedTransports:  cachedTransports,
			cachedConnections: cachedConnections,
		}
	} else {
		return &roundTripper{
			dialer: proxy.Direct,

			Host:            rt.Host,
			ClientHelloID:   rt.ClientHelloID,
			Certificate:     rt.Certificate,
			Privatekey:      rt.Privatekey,
			ClientHelloSpec: rt.ClientHelloSpec,

			ignoreCertIssues: false,

			cachedTransports:  cachedTransports,
			cachedConnections: cachedConnections,
		}
	}
}
