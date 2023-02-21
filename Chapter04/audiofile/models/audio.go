package models

import (
	"bytes"
	"encoding/json"
)

type Audio struct {
	Id       string   `json:"Id"`
	Path     string   `json:"Path"`
	Metadata Metadata `json:"Metadata"`
	Status   string   `json:"Status"`
	Error    []string `json:"Error"`
}

func (a *Audio) JSON() (string, error) {
	audioJSON, err := json.Marshal(a)
	if err != nil {
		return "", err
	}
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(audioJSON), "", "    "); err != nil {
		return "", err
	}
	return prettyJSON.String(), nil
}
