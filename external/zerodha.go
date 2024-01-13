package external

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
)

const (
	KiteLoginUrl     = "https://kite.zerodha.com/api/login"
	KitePortfolioUrl = "https://api.kite.trade/portfolio/holdings"
	Kite2FAUrl       = "https://kite.zerodha.com/api/twofa"
	KiteOrdersUrl    = "https://api.kite.trade/orders"
)

var (
	encToken = ""
)

func LoginAndSyncZerodha(userId string, password string, totpkey string) error {
	//call login api for request id
	client := resty.New().
		SetRetryCount(2).
		SetRetryWaitTime(200 * time.Millisecond).
		SetRetryAfter(nil).
		SetTimeout(2000 * time.Millisecond)
	client.SetDebug(true)
	client.SetContentLength(true)
	loginReq := client.R().
		SetFormData(map[string]string{
			"user_id":  userId,
			"password": password,
		}).
		SetHeader("Cache-Control", "no-cache")

	loginResp, err := loginReq.Execute(http.MethodPost, KiteLoginUrl)
	if err != nil {
		log.Println("Error making zerodha login request:", err)
		return err
	}

	if loginResp.StatusCode() != http.StatusOK {
		log.Println("return data: ", string(loginResp.Body()))
		return fmt.Errorf("failed to return 200 OK")
	}

	// Access response data
	responseData := loginResp.Body()

	var loginResponse ZerodhaLoginResponse
	if err := json.Unmarshal(responseData, &loginResponse); err != nil {
		log.Println("Error unmarshalling zerodha login response:", err)
		return err
	}

	reqId := loginResponse.Data.RequestID
	userId = loginResponse.Data.UserID
	totp := generatePassCode(totpkey)

	twoFAReq := client.R().
		SetFormData(map[string]string{
			"request_id":  reqId,
			"twofa_value": totp,
			"user_id":     userId,
		}).
		SetHeader("Cache-Control", "no-cache")

	twoFAResp, err := twoFAReq.Execute(http.MethodPost, Kite2FAUrl)
	if err != nil {
		log.Println("Error making zerodha login request:", err)
		return err
	}

	if twoFAResp.StatusCode() != http.StatusOK {
		log.Println("return data: ", string(twoFAResp.Body()))
		return fmt.Errorf("failed to return 200 OK")
	}

	for _, cookie := range twoFAResp.Cookies() {
		if cookie.Name == "enctoken" {
			encToken = cookie.Value
			break
		}
	}

	orderReq := client.R().
		SetHeader("Cache-Control", "no-cache").
		SetHeader("Authorization", "enctoken "+encToken)
	resp, err := orderReq.Get(KiteOrdersUrl)
	if err != nil {
		log.Println("Error making zerodha login request:", err)
		return err
	}
	if resp.StatusCode() != http.StatusOK {
		log.Println("return data: ", string(resp.Body()))
		return fmt.Errorf("failed to return 200 OK")
	}
	log.Println("orders: ", string(resp.Body()))
	return nil
}

func GetHoldingsForZerodha() ([]HoldingsInfo, error) {

	headers := map[string]string{
		"Authorization": "enctoken " + encToken,
	}

	// Make request using resty
	client := resty.New().
		SetRetryCount(2).
		SetRetryWaitTime(200 * time.Millisecond).
		SetRetryAfter(nil).
		SetTimeout(2000 * time.Millisecond)
	// client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	req := client.R().
		SetHeaders(headers)
	resp, err := req.Execute(http.MethodGet, KitePortfolioUrl)
	if err != nil {
		log.Println("Error making request:", err)
		return nil, err
	}

	// Access response data
	responseData := resp.Body()

	// Print response data
	log.Println("Response Body:", string(responseData))

	if resp.StatusCode() != 200 {
		log.Println("failed to fetch holdings")
		return nil, errors.New("vendor api returned non 200 status")
	}

	var response ZerodhaPortfolioResponse
	err = json.Unmarshal(responseData, &response)
	if err != nil {
		log.Println("Error reading response body:", err)
		return nil, err
	}

	var holdings []HoldingsInfo
	for _, stock := range response.Data {

		if err != nil {
			return nil, err
		}
		holdings = append(holdings, HoldingsInfo{
			Symbol:    stock.Tradingsymbol,
			Quantity:  stock.Quantity,
			Isin:      stock.Isin,
			AvgPrice:  stock.AveragePrice,
			UpdatedOn: time.Now(),
		})
	}

	return holdings, nil
}
