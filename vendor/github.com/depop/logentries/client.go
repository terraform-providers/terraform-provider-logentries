package logentries

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Request struct {
	ApiKey string
}
type LogSetClient struct {
	Request
}
type LogClient struct {
	Request
}

type Client struct {
	Log    *LogClient
	LogSet *LogSetClient
}

func New(apikey string) *Client {
	return &Client{
		LogSet: &LogSetClient{
			Request{
				ApiKey: apikey,
			},
		},
		Log: &LogClient{
			Request{
				ApiKey: apikey,
			},
		},
	}
}

func (r *Request) getLogentries(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	return r.performrequest(req)
}

func (r *Request) postLogentries(url string, payload []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	return r.performrequest(req)
}

func (r *Request) putLogentries(url string, payload []byte) ([]byte, error) {
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	return r.performrequest(req)
}
func (r *Request) deleteLogentries(url string) (bool, error) {
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return false, err
	}

	client := &http.Client{}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", r.ApiKey)

	res, err := client.Do(req)
	if err != nil {
		return false, err
	}

	if res.StatusCode != 204 {
		return false, fmt.Errorf("Logset was not deleted, response was %v", res.StatusCode)
	}

	return true, nil
}

func (r *Request) performrequest(req *http.Request) ([]byte, error) {
	client := &http.Client{}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", r.ApiKey)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, err
}
