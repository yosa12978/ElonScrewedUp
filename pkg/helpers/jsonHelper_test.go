package helpers_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/yosa12978/ElonScrewedUp/pkg/helpers"
)

func TestDump(t *testing.T) {
	payload := make(map[string]interface{})
	payload["string"] = "string"
	payload["int"] = 15
	res, err := helpers.JsonDump(os.Stdout, payload)
	if err != nil {
		t.Fatal(err.Error())
	}
	act, _ := json.Marshal(payload)
	if res != string(act) {
		t.Fail()
	}
}

func TestLaod(t *testing.T) {

}
