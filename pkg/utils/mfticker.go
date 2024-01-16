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

// var (
// 	isinNavMap map[string]MutualFundNav = make(map[string]MutualFundNav)
// )

func GetNavValueFromIsin(isin string) (*MutualFundNav, error) {
	// if nav, found := isinNavMap[isin]; found {
	// 	return &nav, nil
	// }

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
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		log.Println("unsuccessful response. body: ", string(resp.Body()))
		return nil, fmt.Errorf("failed to fetch tick data")
	}

	var data PaytmMoneyMfTickerResponse
	if err := json.Unmarshal(resp.Body(), &data); err != nil {
		log.Println("failed to marshal data. ", err.Error())
		return nil, err
	}
	nav := MutualFundNav{
		Nav:        data.Data.PageLoad.NavVal,
		Isin:       data.Data.PageLoad.Isin,
		SchemeName: data.Data.PageLoad.Name,
		UpdatedOn:  time.Now(),
	}
	// isinNavMap[isin] = nav
	return &nav, nil
}
