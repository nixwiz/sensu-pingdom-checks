package pingdom

import (
	"encoding/json"
	"io/ioutil"
)

type CreditService struct {
	client *Client
}

func (cs *CreditService) Info() (CreditResponseDetails, error) {

	req, err := cs.client.NewRequest("GET", "/credits")
	if err != nil {
		return CreditResponseDetails{}, err
	}

	resp, err := cs.client.client.Do(req)
	if err != nil {
		return CreditResponseDetails{}, err
	}
	defer resp.Body.Close()

	if err := validateResponse(resp); err != nil {
		return CreditResponseDetails{}, err
	}

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)
	m := &CreditResponse{}
	err = json.Unmarshal([]byte(bodyString), &m)

	return m.Credits, err
}
