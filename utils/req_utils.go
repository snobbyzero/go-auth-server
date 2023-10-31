package utils

import (
	"encoding/json"
	"io"
	"log"
)

// TODO doesnt work
func Unmarshal(obj interface{}, body io.ReadCloser) error {
	decoder := json.NewDecoder(body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(obj); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
