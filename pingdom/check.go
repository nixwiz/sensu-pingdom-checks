package pingdom

import (
	"encoding/json"
	"io/ioutil"
)

type CheckService struct {
	client *Client
}

func (cs *CheckService) List() ([]ListChecksDetails, error) {

	req, err := cs.client.NewRequest("GET", "/checks")
	if err != nil {
		return nil, err
	}

	resp, err := cs.client.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := validateResponse(resp); err != nil {
		return nil, err
	}

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)
	m := &ListChecksResponse{}
	err = json.Unmarshal([]byte(bodyString), &m)

	return m.Checks, err
}
