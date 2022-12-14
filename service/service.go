package service

import (
	"comm-api/connector"
	"comm-api/models"
	"encoding/json"
)

func Downloadallmsg() (string, error) {

	var msgs []models.Message

	msgs, err := connector.GetAll()

	if err != nil {
		return "", err
	}

	content, err := json.Marshal(msgs)
	if err != nil {
		return "", err
	}

	return string(content), nil

}
