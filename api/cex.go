package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/harsh-m-patil/cryptomasters/datatypes"
)

const apiUrl = "https://cex.io/api/ticker/%s/%s"

func GetRate(cryptoCurrency, currency string) (*datatypes.Rate, error) {
	res, err := http.Get(fmt.Sprintf(
		apiUrl,
		strings.ToUpper(cryptoCurrency),
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

	rate := datatypes.Rate{Currency: cryptoCurrency, Price: response.Bid}

	return &rate, nil
}
