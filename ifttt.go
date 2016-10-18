package ifttt

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type IFTTT struct {
	Value1 string `json:"value1"`
	Value2 string `json:"value2"`
	Value3 string `json:"value3"`
}

func Trigger(key, event, value1, value2, value3 string) (*http.Response, error) {
	msg := IFTTT{value1, value2, value3}
	postUrl := "https://maker.ifttt.com/trigger/" + event + "/with/key/" + key

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(msg)

	return http.Post(postUrl, "application/json; charset=utf-8", b)
}
