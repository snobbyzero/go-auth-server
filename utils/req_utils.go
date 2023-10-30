package utils

import (
	"encoding/json"
	"io"
	"log"
)

// TODO doesnt work
func GetObjectFromJson(obj interface{}, body io.ReadCloser) error {
	decoder := json.NewDecoder(body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&obj)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
