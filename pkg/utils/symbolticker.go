package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/sauravkuila/portfolio-worth/external"
)

// returns a map of token vs symbol ltp
func GetSymbolLtpData(tokens []string) (map[string]SymbolLtp, error) {

	baseUrl := "https://apiconnect.angelbroking.com/rest/secure/angelbroking/market/v1/quote/"

	var (
		request    LtpTickerRequest
		response   LtpTickerResponse
		tokenSlice [][]string           = make([][]string, 0)
		retResp    map[string]SymbolLtp = make(map[string]SymbolLtp)
	)
	request.Mode = "LTP"

	//api limits 50 tokens only. since nse and bse are searched together for token, we parse 25 at a time
	len := len(tokens)
	i := 0
	for len > 0 {
		end := i + 25
		if (i + 25) > cap(tokens) {
			end = cap(tokens)
		}
		tokenSlice = append(tokenSlice, tokens[i:end])
		len = len - 25
		i = i + 25
	}

	//headers
	headers := map[string]string{
		// "X-PrivateKey":  key,
		// "Authorization": "Bearer "+auth,
		"Content-Type":  "application/json",
		"X-PrivateKey":  external.GetAngelAppKey(),
		"Authorization": "Bearer " + external.GetAngelAccessToken(),
	}

	// Make request using resty
	client := resty.New().
		SetRetryCount(2).
		SetRetryWaitTime(200 * time.Millisecond).
		SetRetryAfter(nil).
		SetTimeout(2000 * time.Millisecond).
		SetDebug(true).
		SetContentLength(true)

	for _, t := range tokenSlice {
		request.ExchangeTokens.Nse = t
		request.ExchangeTokens.Bse = t

		payload, err := json.Marshal(request)
		if err != nil {
			log.Println("failed to marshal", err.Error())
			return nil, err
		}

		req := client.R().
			SetHeaders(headers).
			SetBody(payload)
		resp, err := req.Execute(http.MethodPost, baseUrl)
		if err != nil {
			log.Println("Error making request:", err)
			return nil, err
		}

		if resp.StatusCode() != http.StatusOK {
			log.Println("unsuccessful response. body: ", string(resp.Body()))
			return nil, fmt.Errorf("failed to fetch tick data")
		}

		if err := json.Unmarshal(resp.Body(), &response); err != nil {
			log.Println("failed to marshal data. ", err.Error())
			return nil, err
		}
		for _, info := range response.Data.Fetched {
			retResp[info.Token] = SymbolLtp{
				Exchange:      info.Exchange,
				TradingSymbol: info.TradingSymbol,
				Token:         info.Token,
				Ltp:           info.Ltp,
				UpdatedOn:     time.Now(),
			}
		}
		//to manage rps
		time.Sleep(1 * time.Second)
	}
	return retResp, nil
}
