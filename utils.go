package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func writeError(w http.ResponseWriter, code int, message string, err error) {
	fmt.Println(message, err)
	http.Error(w, message, code)
}

func marshalJson(i interface{}) ([]byte, error) {
	result, err := json.Marshal(i)
	if err != nil {
		return nil, err
	}
	if string(result) == "null" {
		result = []byte("[]")
	}
	return result, nil
}
