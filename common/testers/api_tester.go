package testers

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

// ApiTester contains Server and Client
type ApiTester struct {
	Server *httptest.Server
	Client *http.Client
}

// ApiTesterResponse contains the response in convenience
// objects to be evaluated
type ApiTesterResponse struct {
	RawResponse *http.Response
	StatusCode  int
	BodyStr     string
}

// NewApiTester returns an ApiTester object, initializing
// its Server and Client components
func NewApiTester(router http.Handler) *ApiTester {
	return &ApiTester{
		Server: httptest.NewServer(router),
		Client: &http.Client{},
	}
}

// Close shut downs the HTTP Server of the tester
func (t *ApiTester) Close() {
	t.Server.Close()
}

// Get performs an HTTP GET call
func (t *ApiTester) Get(route string) (*ApiTesterResponse, error) {
	// prepare the request here
	request, err := http.NewRequest("GET", "http://"+t.Server.Listener.Addr().String()+route, nil)
	if err != nil {
		return nil, err
	}

	// do the actual request here
	response, err := t.Client.Do(request)
	if err != nil {
		return nil, err
	}

	bodyBuffer, _ := ioutil.ReadAll(response.Body)

	return &ApiTesterResponse{
		RawResponse: response,
		StatusCode:  response.StatusCode,
		BodyStr:     string(bodyBuffer),
	}, nil
}
