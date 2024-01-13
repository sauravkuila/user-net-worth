package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

var (
	isinNavMap map[string]float64 = make(map[string]float64)
)

func GetNavValueFromIsin(isin string) (float64, error) {
	if nav, found := isinNavMap[isin]; found {
		return nav, nil
	}

	baseUrl := "https://www.paytmmoney.com/api/mf/isin/" + isin

	headers := map[string]string{
		"x-request-id": uuid.New().String(),
	}
	// Make request using resty
	client := resty.New().
		SetRetryCount(2).
		SetRetryWaitTime(200 * time.Millisecond).
		SetRetryAfter(nil).
		SetTimeout(2000 * time.Millisecond).
		// SetDebug(true).
		SetContentLength(true)
	req := client.R().
		SetHeaders(headers)
	resp, err := req.Execute(http.MethodGet, baseUrl)
	if err != nil {
		log.Println("Error making request:", err)
		return 0, err
	}

	if resp.StatusCode() != http.StatusOK {
		log.Println("unsuccessful response. body: ", string(resp.Body()))
		return 0, fmt.Errorf("failed to fetch tick data")
	}

	var data PaytmMoneyMfTickerResponse
	if err := json.Unmarshal(resp.Body(), &data); err != nil {
		log.Println("failed to marshal data. ", err.Error())
		return 0, err
	}
	isinNavMap[isin] = data.Data.PageLoad.NavVal
	return data.Data.PageLoad.NavVal, nil
}
