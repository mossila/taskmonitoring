package util

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

// Binding : http.Request to json
func Binding(r *http.Request, out interface{}) error {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		return err
	}
	if err := r.Body.Close(); err != nil {
		return err
	}
	if err := json.Unmarshal(body, &out); err != nil {
		return err

	}
	return nil
}
