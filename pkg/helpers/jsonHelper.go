package helpers

import (
	"encoding/json"
	"io"
)

func JsonDump(w io.Writer, payload interface{}) (string, error) {
	dump, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}
	_, err = w.Write(dump)
	return string(dump), err
}

func JsonLoad(dump string) (interface{}, error) {
	var res interface{}
	err := json.Unmarshal([]byte(dump), &res)
	return res, err
}
