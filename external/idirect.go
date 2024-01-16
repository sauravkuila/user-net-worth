package external

import (
	"bytes"
	"crypto/sha256"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os/exec"
	"strconv"
	"time"

	"github.com/go-resty/resty/v2"
)

const (
	CallbackUrl         = "http://127.0.0.1:8080/callback/idirect"
	IdirectLoginBaseUrl = "https://api.icicidirect.com/apiuser/login?api_key=AppKey"
)

var (
	iDirectSessionKey   = ""
	iDirectSessionToken = ""
	AppKey              = ""
	SecretKey           = ""
)

func LoginAndSyncIDirect(appKey string, secretKey string) {
	AppKey = appKey
	SecretKey = secretKey
}

func GetHoldingsForIdirect() ([]HoldingsInfo, error) {
	if err := generateSessionToken(); err != nil {
		log.Println("error in creating session token")
		return nil, err
	}
	log.Println("session token: ", iDirectSessionToken)
	// if false {
	// 	appKey := url.QueryEscape(AppKey)
	// 	log.Println(appKey)
	// }

	// API endpoint
	url := "https://api.icicidirect.com/breezeapi/api/v1/portfolioholdings"

	// Payload
	payloadData := map[string]interface{}{
		"exchange_code":  "NSE",
		"from_date":      "",
		"to_date":        "",
		"stock_code":     "",
		"portfolio_type": "",
	}
	payload, err := json.Marshal(payloadData)
	if err != nil {
		log.Println("Error marshaling JSON payload:", err)
		return nil, err
	}

	timestamp, checksum, _ := prepareRequest(payloadData)

	// Request headers
	headers := map[string]string{
		"Content-Type":   "application/json",
		"X-Checksum":     "token " + checksum,
		"X-Timestamp":    timestamp,
		"X-AppKey":       AppKey,
		"X-SessionToken": iDirectSessionToken,
	}

	// Create request
	req, err := http.NewRequest("GET", url, bytes.NewBuffer(payload))
	if err != nil {
		log.Println("Error creating request:", err)
		return nil, err
	}

	// Set headers
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// Make request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error making request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	// Read and print response data
	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return nil, err
	}
	log.Println("Response Body:", string(responseData))

	var response IDirectHoldingsResponse
	err = json.Unmarshal(responseData, &response)
	if err != nil {
		log.Println("Error reading response body:", err)
		return nil, err
	}

	if response.Status != 200 {
		log.Println("failed to fetch holdings")
		return nil, errors.New("vendor api returned non 200 status")
	}

	var holdings []HoldingsInfo
	for _, stock := range response.Success {
		qty, err := strconv.ParseInt(stock.Quantity, 10, 64)
		if err != nil {
			return nil, err
		}
		price, err := strconv.ParseFloat(stock.AveragePrice, 64)
		if err != nil {
			return nil, err
		}
		holdings = append(holdings, HoldingsInfo{
			Symbol:    stock.StockCode,
			Quantity:  qty,
			Isin:      "",
			AvgPrice:  price,
			UpdatedOn: time.Now(),
		})
	}

	return holdings, nil
}

func generateSessionToken() error {
	url := "https://api.icicidirect.com/breezeapi/api/v1/customerdetails"
	var (
		request  IDirectCustomerDetailRequest
		response IDirectCustomerDetailResponse
	)
	request.SessionToken = iDirectSessionKey
	request.AppKey = AppKey
	payload, err := json.Marshal(request)
	if err != nil {
		log.Println("failed to marshal obj", err.Error())
		return err
	}

	req, err := http.NewRequest("GET", url, bytes.NewBuffer(payload))
	if err != nil {
		log.Println("Error creating request:", err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error making request:", err)
		return err
	}
	defer resp.Body.Close()

	log.Println("Response Status:", resp.Status)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return err
	}
	log.Println("Response Body:", string(body))

	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Println("Error unmarshalling response:", err)
		return err
	}

	if response.Status != 200 {
		return errors.New("failed to fetch session token")
	}

	iDirectSessionToken = response.Success.SessionToken

	return nil
}

func GetDematHoldingsForIDirect() ([]HoldingsInfo, error) {
	if iDirectSessionToken == "" && iDirectSessionKey != "" {
		if err := generateSessionToken(); err != nil {
			return nil, err
		}
	} else if iDirectSessionKey == "" {
		return nil, fmt.Errorf("no login done yet")
	}

	// Command to run the Python script with arguments
	cmd := exec.Command("python", "../external/idirect_demat.py", AppKey, SecretKey, iDirectSessionToken)

	// Run the command and capture the output
	responseData, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println("Error executing command:", err)
		return nil, err
	}

	// Print the script output
	log.Println("Script Output:", string(responseData))

	var response IDirectDematHoldingsResponse
	err = json.Unmarshal(responseData, &response)
	if err != nil {
		log.Println("Error reading response body:", err)
		return nil, err
	}

	if response.Status != 200 {
		log.Println("failed to fetch holdings")
		return nil, errors.New("vendor api returned non 200 status")
	}

	var holdings []HoldingsInfo
	for _, stock := range response.Success {
		qty, err := strconv.ParseInt(stock.Quantity, 10, 64)
		if err != nil {
			return nil, err
		}
		holdings = append(holdings, HoldingsInfo{
			Symbol:    stock.StockCode,
			Quantity:  qty,
			Isin:      stock.StockISIN,
			AvgPrice:  0.0,
			UpdatedOn: time.Now(),
		})
	}

	return holdings, nil
}

func GetDematHoldingsForIDirect_Retired() ([]HoldingsInfo, error) {
	if err := generateSessionToken(); err != nil {
		log.Println("error in creating session token")
		return nil, err
	}

	log.Println("session token: ", iDirectSessionToken)

	path := "https://api.icicidirect.com/breezeapi/api/v1/dematholdings"

	// Payload
	payloadData := map[string]interface{}{}

	timestamp, checksum, payload := prepareRequest(payloadData)
	if payload == nil {
		log.Println("failed to prepare request")
		return nil, fmt.Errorf("failed to prepare payload")
	}
	log.Println("payload", payload)

	headers := map[string]string{
		"X-Checksum":     "token " + checksum,
		"X-Timestamp":    timestamp,
		"X-AppKey":       AppKey,
		"X-SessionToken": iDirectSessionToken,
		"Content-Type":   "application/json",
		"host":           "api.icicidirect.com",
	}

	// Make request using resty
	client := resty.New().
		SetRetryCount(2).
		SetRetryWaitTime(200 * time.Millisecond).
		SetRetryAfter(nil).
		SetTimeout(2000 * time.Millisecond).
		SetDebug(true).
		SetAllowGetMethodPayload(true).
		SetContentLength(true).
		SetTLSClientConfig(&tls.Config{RootCAs: nil})
	// client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	req := client.R().
		SetHeaders(headers).
		SetBody(payload)
	resp, err := req.Execute(http.MethodGet, path)
	if err != nil {
		log.Println("Error making request:", err)
		return nil, err
	}

	// Access response data
	responseData := resp.Body()

	// Print response data
	log.Println("Response Body:", string(responseData))
	log.Println("Response Status:", resp.StatusCode())

	var response IDirectDematHoldingsResponse
	err = json.Unmarshal(responseData, &response)
	if err != nil {
		log.Println("Error reading response body:", err)
		return nil, err
	}

	if response.Status != 200 {
		log.Println("failed to fetch holdings")
		return nil, errors.New("vendor api returned non 200 status")
	}

	var holdings []HoldingsInfo
	for _, stock := range response.Success {
		qty, err := strconv.ParseInt(stock.Quantity, 10, 64)
		if err != nil {
			return nil, err
		}
		holdings = append(holdings, HoldingsInfo{
			Symbol:    stock.StockCode,
			Quantity:  qty,
			Isin:      stock.StockISIN,
			AvgPrice:  0.0,
			UpdatedOn: time.Now(),
		})
	}

	return holdings, nil
}

func prepareRequest(body map[string]interface{}) (string, string, []byte) {
	// App related Secret Key
	secretKey := SecretKey

	// 'body' is the request-body of your current request
	payload, err := json.Marshal(body)
	if err != nil {
		log.Println("Error marshaling JSON:", err)
		return "", "", nil
	}
	log.Println("payload: ", string(payload))

	// Time_stamp & checksum generation for request-headers
	timeStamp := time.Now().UTC().Format("2006-01-02T15:04:05") + ".000Z"
	// dataToHash := timeStamp + string(payload) + secretKey
	dataToHash := timeStamp + string(payload) + secretKey
	checksum := sha256.Sum256([]byte(dataToHash))
	checksumHex := fmt.Sprintf("%x", checksum)

	log.Println("Time Stamp:", timeStamp)
	log.Println("Checksum:", string(checksumHex))
	return timeStamp, string(checksumHex), payload
}

func SetIDirectSessionKeyFromCallback(key string) {
	log.Println("sessionkey: ", iDirectSessionKey)
	iDirectSessionKey = key
	log.Println("sessionkey: ", iDirectSessionKey)
	generateSessionToken()
}

func GetIDirectLoginUrl() string {
	appKey := url.QueryEscape(AppKey)
	log.Println(appKey)
	return "https://api.icicidirect.com/apiuser/login?api_key=" + appKey
}

//for checksum generation
// import json
// from datetime import datetime
// import hashlib

// payload = json.dumps({})

// appkey = "23vU6735q98H876r&01Ia88V1=i3Xx+8"
// secret_key = "l&h8_69E86s9V23024W1r273fD0Q024j"
// time_stamp = datetime.utcnow().isoformat()[:19] + '.000Z'
// checksum = hashlib.sha256((time_stamp+payload+secret_key).encode("utf-8")).hexdigest()

// print(time_stamp)
// print(checksum)
