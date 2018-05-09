package logentries

import (
	"encoding/json"
)

type LogSetCreateResponse struct {
	LogSet `json:"logset"`
}

type LogSetReadResponse struct {
	LogSet `json:"logset"`
}

type LogSetUpdateResponse struct {
	LogSet `json:"logset"`
}

type LogSetCreateRequest struct {
	LogSet LogSetFields `json:"logset"`
}

type LogSetReadRequest struct {
	ID string
}

type LogSetUpdateRequest struct {
	ID     string       `json:"id"`
	LogSet LogSetFields `json:"logset"`
}

type LogSetDeleteRequest struct {
	ID string
}

type LogSetUpdateRequestWrapper struct {
	LogSet LogSetFields `json:"logset"`
}

type LogSetFields struct {
	Name        string    `json:"name"`
	Description string    `json:"description,omitempty"`
	UserData    struct{}  `json:"user_data,omitempty"`
	LogsInfo    []LogInfo `json:"logs_info,omitempty"`
}

type LogSet struct {
	ID          string      `json:"id,omitempty"`
	Name        string      `json:"name"`
	Description interface{} `json:"description,omitempty"`
	UserData    UserData    `json:"user_data"`
	LogsInfo    []LogInfo   `json:"logs_info"`
}

type LogInfo struct {
	ID    string `json:"id"`
	Links []Link `json:"links"`
	Name  string `json:"name"`
}

type Link struct {
	Href string `json:"href"`
	Rel  string `json:"rel"`
}

type UserData struct {
	LeDistname string `json:"le_distname"`
	LeDistver  string `json:"le_distver"`
	LeNameintr string `json:"le_nameintr"`
}

func (l *LogSetClient) Create(createRequest *LogSetCreateRequest) (*LogSetCreateResponse, error) {
	url := "https://rest.logentries.com/management/logsets/"

	payload, err := json.Marshal(createRequest)
	if err != nil {
		return nil, err
	}

	resp, err := l.postLogentries(url, payload)
	if err != nil {
		return nil, err
	}

	var logset LogSetCreateResponse

	err = json.Unmarshal(resp, &logset)
	if err != nil {
		return nil, err
	}

	return &logset, nil
}

func (l *LogSetClient) Read(readRequest *LogSetReadRequest) (*LogSetReadResponse, error) {
	url := "https://rest.logentries.com/management/logsets/" + readRequest.ID

	resp, err := l.getLogentries(url)
	if err != nil {
		return nil, err
	}

	var logset LogSetReadResponse

	err = json.Unmarshal(resp, &logset)
	if err != nil {
		return nil, err
	}

	return &logset, nil
}

func (l *LogSetClient) Update(updateRequest *LogSetUpdateRequest) (*LogSetUpdateResponse, error) {
	url := "https://rest.logentries.com/management/logsets/" + updateRequest.ID

	payload, err := json.Marshal(&LogSetUpdateRequestWrapper{LogSet: updateRequest.LogSet})
	if err != nil {
		return nil, err
	}

	resp, err := l.putLogentries(url, payload)
	if err != nil {
		return nil, err
	}

	var logset LogSetUpdateResponse

	err = json.Unmarshal(resp, &logset)
	if err != nil {
		return nil, err
	}

	return &logset, nil

}
func (l *LogSetClient) Delete(deleteRequest *LogSetDeleteRequest) (bool, error) {
	url := "https://rest.logentries.com/management/logsets/" + deleteRequest.ID

	success, err := l.deleteLogentries(url)
	if err != nil {
		return false, err
	}

	return success, nil
}
