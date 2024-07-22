package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/harsh-m-patil/cryptomasters/datatypes"
)

const apiUrl = "https://cex.io/api/ticker/%s/USD"

func GetRate(currency string) (*datatypes.Rate, error) {
	if len(currency) != 3 {
		return nil, fmt.Errorf(
			"3 characters required; %d recieved",
			len(currency),
		)
	}
	res, err := http.Get(fmt.Sprintf(
		apiUrl,
		strings.ToUpper(currency),
	))
	if err != nil {
		return nil, err
	}

	var response CEXResponse
	if res.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(bodyBytes, &response)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("status code recieved %v", res.StatusCode)
	}

	rate := datatypes.Rate{Currency: currency, Price: response.Bid}

	return &rate, nil
}
