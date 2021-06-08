package psychoclient

import (
	"errors"
	"io/ioutil"
	"net/http"
	"sicko-aio-auth/models"

	tls "github.com/refraction-networking/utls"
)

type Client struct {
	client http.Client
	err    error
}

func NewClient(s *SessionBuilder) *Client {

	var client http.Client
	var err error

	if s.UseDefaultClient {
		client, err = newDefaultClient(&clientConfig{FollowRedirects: s.FollowRedirects}, s.Proxy)
	} else {
		config := &clientConfig{ // custom chrome tls fingerprint config
			Certificate:     s.Certificate,
			Privatekey:      s.Privatekey,
			FollowRedirects: s.FollowRedirects,
			ClientHelloID:   tls.HelloCustom,
			ClientHelloSpec: getChromeClientHelloSpecs(), // may change according as user desire
		}
		client, err = newClient(config, s.Proxy)
	}

	return &Client{client: client, err: err}
}

func (c *Client) DoNewRequest(b *RequestBuilder) (res *http.Response, respBody []byte, err *models.Error) {
	defer c.close()

	if c.err != nil {
		return nil, nil, &models.Error{Error: c.err, Code: 500, Message: "Error Building Request Client"}
	}

	// setup a new http request, setup request method & endpoint & payload'
	req, _ := http.NewRequest(b.Method, b.Endpoint, b.Payload)

	if len(b.Host) > 0 { // set the request host
		req.Host = b.Host
	}
	if len(b.Headers) > 0 { // set the request headers
		for key, value := range b.Headers {
			req.Header.Add(key, value)
		}
	}

	// starts a new http request
	channel := make(chan *response, 1)
	defer close(channel)

	go func() {
		res, err := c.client.Do(req)
		if err != nil || res == nil { // checking if empty response or err occured during the request
			channel <- &response{
				httpResponse: nil,
				err:          &models.Error{Error: err, Code: 500, Message: "Internal Server Error"},
			}
			return
		}

		channel <- &response{ // puts the response back to channel
			httpResponse: res,
			err:          nil,
		}
	}()

	response := <-channel // waiting for channel to receive response

	res = response.httpResponse
	err = response.err

	if res != nil {
		// set response body
		if res.Body != nil {
			defer response.httpResponse.Body.Close()
			body, e := ioutil.ReadAll(res.Body)
			respBody = body
			if e != nil {
				err = &models.Error{Error: errors.New("ERROR_PARSING_RESPONSE_BODY"), Code: 500, Message: "Error Parsing Response Body"}
			}
		}
	}

	return res, respBody, err
}

func (c *Client) RoundTripNewRequest(b *RequestBuilder) (res *http.Response, respBody []byte, err *models.Error) {
	defer c.close()

	if c.err != nil {
		return nil, nil, &models.Error{Error: c.err, Code: 500, Message: "Error Building Request Client"}
	}

	// setup a new http request, setup request method & endpoint & payload'
	req, _ := http.NewRequest(b.Method, b.Endpoint, b.Payload)

	if len(b.Host) > 0 { // set the request host
		req.Host = b.Host
	}
	if len(b.Headers) > 0 { // set the request headers
		for key, value := range b.Headers {
			req.Header.Add(key, value)
		}
	}

	// starts a new http request
	channel := make(chan *response, 1)
	defer close(channel)

	go func() {
		res, err := c.client.Transport.RoundTrip(req)
		if err != nil || res == nil { // checking if empty response or err occured during the request
			channel <- &response{
				httpResponse: nil,
				err:          &models.Error{Error: err, Code: 500, Message: "Internal Server Error"},
			}
			return
		}

		channel <- &response{ // puts the response back to channel
			httpResponse: res,
			err:          nil,
		}
	}()

	response := <-channel // waiting for channel to receive response

	res = response.httpResponse
	err = response.err

	if res != nil {
		// set response body
		if res.Body != nil {
			defer response.httpResponse.Body.Close()
			body, e := ioutil.ReadAll(res.Body)
			respBody = body
			if e != nil {
				err = &models.Error{Error: errors.New("ERROR_PARSING_RESPONSE_BODY"), Code: 500, Message: "Error Parsing Response Body"}
			}
		}
	}

	return res, respBody, err
}

func (c *Client) close() {
	c.client.CloseIdleConnections()
}
