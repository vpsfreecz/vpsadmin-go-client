package client

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// DoQueryStringRequests makes a HTTP requests in which input parameters are
// sent as query parameters.
func (client *Client) DoQueryStringRequest(path string, queryParams map[string]string, output interface{}) error {
	httpClient := &http.Client{}
	url := client.Url + path

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return err
	}

	if client.Authentication != nil {
		client.Authentication.Authenticate(req)
	}

	q := req.URL.Query()

	for k, v := range queryParams {
		q.Add(k, v)
	}

	req.URL.RawQuery = q.Encode()

	resp, err := httpClient.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, output); err != nil {
		return err
	}

	return nil
}

// DoBodyRequest makes a HTTP requests in which the input parameters are sent
// within the request body, encoded in JSON.
func (client *Client) DoBodyRequest(method string, path string, params interface{}, output interface{}) error {
	httpClient := &http.Client{}
	url := client.Url + path

	jsonData, err := json.Marshal(params)

	if err != nil {
		return err
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonData))

	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	if client.Authentication != nil {
		client.Authentication.Authenticate(req)
	}

	resp, err := httpClient.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, output); err != nil {
		return err
	}

	return nil
}
