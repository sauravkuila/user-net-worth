package external

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-resty/resty/v2"
)

const (
	MfCentralPortfolioUrl = "https://services.mfcentral.com/user/getportfolio"
)

var (
	mfCentralEncToken = ""
	mfCentralPan      = ""
	mfCentralPekrn    = ""
	mfCentralMobile   = ""
	mfCentralEmail    = ""
)

func GetMutualFundHoldingsFromMfCentral() ([]MfHoldingsInfo, error) {

	mfCentralEncToken = "eyJhbGciOiJIUzI1NiJ9.eyJzdWIiOiJmYTllZGQ4ZC03OTFmLTQxZGMtYTM1MC0wMjlkOTg2OGNkOWMiLCJpYXQiOjE3MDUxODIwMTEsImp0aSI6IjJmZTNjODczLTMwYmItNDcxNi04NzcwLTIyNjlhNTQyYWJkNiJ9.c9lO58Mv1wSnIx-hfDV657B95IyVo0ZDEoKg-NZK69I"
	mfCentralPan = "ABCDE1234F"
	mfCentralPekrn = ""
	mfCentralMobile = "+919999999999"
	mfCentralEmail = ""

	request := MFCentralPortfolioRequest{
		Email:  mfCentralEmail,
		Mobile: mfCentralMobile,
		Pan:    mfCentralPan,
		Pekrn:  mfCentralPekrn,
	}
	payload, err := json.Marshal(request)
	if err != nil {
		log.Println("failed to marshal payload", err.Error())
		return nil, err
	}
	headers := map[string]string{
		"authorization": "Bearer " + mfCentralEncToken,
		"origin":        "https://app.mfcentral.com",
		"content-type":  "application/json",
	}

	// Make request using resty
	client := resty.New().
		SetRetryCount(2).
		SetRetryWaitTime(200 * time.Millisecond).
		SetRetryAfter(nil).
		SetTimeout(2000 * time.Millisecond).
		SetDebug(true).
		SetContentLength(true)
	req := client.R().
		SetHeaders(headers).
		SetBody(payload)
	resp, err := req.Execute(http.MethodPost, MfCentralPortfolioUrl)
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

	var response MfCentralPortfolioResponse
	err = json.Unmarshal(responseData, &response)
	if err != nil {
		log.Println("Error reading response body:", err)
		return nil, err
	}

	var holdings []MfHoldingsInfo
	currentWorth := 0.0
	for _, itr := range response.Data {

		for _, scheme := range itr.Schemes {
			if scheme.IsDemat != "Y" {
				var (
					units        float64
					costPrice    float64
					currentPrice float64
					price        float64
					err          error
				)
				// units
				switch scheme.AvailableUnits.(type) {
				case int, int64:
					units = float64(scheme.AvailableUnits.(int64))
				case float64:
					units = scheme.AvailableUnits.(float64)
				case string:
					units, err = strconv.ParseFloat(scheme.AvailableUnits.(string), 64)
					if err != nil {
						return nil, err
					}
				default:
					log.Println("units type is unknown!", scheme.AvailableUnits)
				}
				//price
				switch scheme.CostValue.(type) {
				case int, int64:
					costPrice = float64(scheme.CostValue.(int64))
				case float64:
					costPrice = scheme.CostValue.(float64)
				case string:
					costPrice, err = strconv.ParseFloat(scheme.CostValue.(string), 64)
					if err != nil {
						return nil, err
					}
				default:
					log.Println("units type is unknown!", scheme.CostValue)
				}
				switch scheme.CurrentMktValue.(type) {
				case int, int64:
					currentPrice = float64(scheme.CurrentMktValue.(int64))
				case float64:
					currentPrice = scheme.CurrentMktValue.(float64)
				case string:
					currentPrice, err = strconv.ParseFloat(scheme.CurrentMktValue.(string), 64)
					if err != nil {
						return nil, err
					}
				default:
					log.Println("units type is unknown!", scheme.CurrentMktValue)
				}
				log.Println("mapping holding scheme", scheme.SchemeName)
				if units > 0.0 {
					price = costPrice / units
				}
				currentWorth += currentPrice
				holdings = append(holdings, MfHoldingsInfo{
					Name:        scheme.SchemeName,
					Folio:       scheme.Folio,
					Isin:        scheme.Isin,
					Quantity:    units,
					AvgPrice:    price,
					InvestedVal: costPrice,
					CurrentVal:  currentPrice,
					UpdatedOn:   time.Now(),
				})
			}
		}
	}
	log.Println("mf current price: ", currentWorth)

	return holdings, nil
}

//for hash generation for password/otp/secret answer fields
// import base64

// def Qi(e):
//     # Inner Function
//     def process_input(input_str=""):
//         t = input_str[:-2]  # Exclude the last two characters
//         n = ''.join([char for char in input_str[-2:] if char != '='])
//         return n

//     # Apply the inner function and concatenate the result
//     t = process_input(base64.b64encode(e.encode()).decode())

//     # Construct a new string and apply base64 encoding
//     result = base64.b64encode(f"X2lMZWFwUH{t}BTVRJel9pTGVhcF".encode()).decode()

//     return result

// # Example usage:
// input_string = "otp/secret"
// output_result = Qi(input_string)
// print(output_result)
