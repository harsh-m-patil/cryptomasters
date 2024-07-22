package api_test

import (
	"testing"

	"github.com/harsh-m-patil/cryptomasters/api"
)

func TestAPICall(t *testing.T) {
	_, err := api.GetRate("")

	if err == nil {
		t.Errorf("Error was not found")
	}
}
