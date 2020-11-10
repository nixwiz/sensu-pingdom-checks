package pingdom

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	defaultBaseURL = "https://api.pingdom.com/api/3.1"
)

type Client struct {
	APIToken string
	BaseURL  *url.URL
	client   *http.Client
	Checks   *CheckService
	Credits  *CreditService
}

type ClientConfig struct {
	APIToken   string
	BaseURL    string
	HTTPClient *http.Client
}

func NewClientWithConfig(config ClientConfig) (*Client, error) {
	var baseURL *url.URL
	var err error
	if len(config.BaseURL) > 0 {
		baseURL, err = url.Parse(config.BaseURL)
	} else {
		baseURL, err = url.Parse(defaultBaseURL)
	}
	if err != nil {
		return nil, err
	}
	c := &Client{
		APIToken: config.APIToken,
		BaseURL:  baseURL,
	}

	if config.HTTPClient != nil {
		c.client = config.HTTPClient
	} else {
		c.client = http.DefaultClient
	}

	c.Checks = &CheckService{client: c}
	c.Credits = &CreditService{client: c}

	return c, nil
}

func (pc *Client) NewRequest(method string, rsc string) (*http.Request, error) {
	baseURL, err := url.Parse(pc.BaseURL.String() + rsc)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, baseURL.String(), nil)
	req.Header.Add("Authorization", "Bearer "+pc.APIToken)
	return req, err
}

func (pc *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := pc.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := validateResponse(resp); err != nil {
		return resp, err
	}

	err = decodeResponse(resp, v)
	return resp, err
}

func decodeResponse(r *http.Response, v interface{}) error {
	if v == nil {
		return fmt.Errorf("nil interface provided to decodeResponse")
	}

	bodyBytes, _ := ioutil.ReadAll(r.Body)
	bodyString := string(bodyBytes)
	err := json.Unmarshal([]byte(bodyString), &v)
	return err
}

// Takes an HTTP response and determines whether it was successful.
// Returns nil if the HTTP status code is within the 2xx range.  Returns
// an error otherwise.
func validateResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	bodyBytes, _ := ioutil.ReadAll(r.Body)
	bodyString := string(bodyBytes)
	m := &errorJSONResponse{}
	err := json.Unmarshal([]byte(bodyString), &m)
	if err != nil {
		return err
	}

	return m.Error
}
