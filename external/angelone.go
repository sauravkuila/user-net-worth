package external

import (
	"errors"
	"fmt"
	"log"
	"math"
	"time"

	"github.com/pquerna/otp/hotp"

	SmartApi "github.com/angel-one/smartapigo"
)

var (
	abClient         *SmartApi.Client = nil
	angelUserSession SmartApi.UserSession
)

func LoginAndSyncAngelOne(userId string, pin string, totpkey string, appKey string) error {
	// Create New Angel Broking Client
	abClient = SmartApi.New(userId, pin, appKey)

	totp := generatePassCode(totpkey)

	// User Login and Generate User Session
	var err error
	angelUserSession, err = abClient.GenerateSession(totp)

	if err != nil {
		fmt.Println("error creating session. Error: ", err.Error())
		return err
	}

	//Renew User Tokens using refresh token
	angelUserSession.UserSessionTokens, err = abClient.RenewAccessToken(angelUserSession.RefreshToken)

	if err != nil {
		fmt.Println("error in renewing access token. Error: ", err.Error())
		return err
	}

	// defer ABClient.Logout()

	fmt.Println("User Session Tokens :- ", angelUserSession.UserSessionTokens)
	fmt.Println("User Access Token :- ", angelUserSession.UserSessionTokens.AccessToken)
	fmt.Println("User Refresh Token :- ", angelUserSession.UserSessionTokens.RefreshToken)
	fmt.Println("User Feed Token :- ", angelUserSession.UserSessionTokens.FeedToken)

	//Get User Profile
	angelUserSession.UserProfile, err = abClient.GetUserProfile()

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("User Profile :- ", angelUserSession.UserProfile)
	fmt.Println("User Session Object :- ", angelUserSession)

	return nil
}

func GetHoldingsForAngel() ([]HoldingsInfo, error) {
	if abClient == nil {
		log.Println("missing client object")
		return nil, fmt.Errorf("please sync broker to initiate holdings")
	}

	if abClient != nil {
		portfolioHoldings := make([]HoldingsInfo, 0)

		//fetch holdings
		holdings, err := abClient.GetHoldings()
		if err != nil {
			fmt.Println("error in fetching holdings", err.Error())
			return nil, err
		}

		for _, holding := range holdings {
			fmt.Println("isin: ", holding.ISIN, " quantity: ", holding.Quantity, "", " average price: ", holding.AveragePrice, "holding: ", holding)
			portfolioHoldings = append(portfolioHoldings, HoldingsInfo{
				Symbol:    holding.Tradingsymbol,
				Quantity:  holding.Quantity,
				Isin:      holding.ISIN,
				AvgPrice:  holding.AveragePrice,
				UpdatedOn: time.Now(),
			})
		}
		return portfolioHoldings, nil
	}

	return nil, errors.New("unable to connect to angel apis")
}

func generatePassCode(utf8string string) string {
	// secret := base32.StdEncoding.EncodeToString([]byte(utf8string))
	// passcode, err := totp.GenerateCodeCustom(secret, time.Now(), totp.ValidateOpts{
	// 	Period:    30,
	// 	Skew:      0,
	// 	Digits:    otp.DigitsSix,
	// 	Algorithm: otp.AlgorithmSHA256,
	// })
	// if err != nil {
	// 	fmt.Println("custom passcode error")
	// 	panic(err)
	// }
	// fmt.Println("custom passcode: ", passcode)

	passcode, err := hotp.GenerateCode(utf8string, uint64(math.Floor(float64(time.Now().Unix())/float64(30))))
	if err != nil {
		fmt.Println("generate passcode error")
		panic(err)
	}
	fmt.Println("generate passcode: ", passcode)
	return passcode
}

func LogoutAngel() {
	if abClient != nil {
		abClient.Logout()
	}
}

func GetAngelAccessToken() string {
	return angelUserSession.UserSessionTokens.AccessToken
}

func GetAngelAppKey() string {
	if abClient != nil {
		return abClient.GetClientApiKey()
	}
	return ""
}
