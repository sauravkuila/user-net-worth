package external

import (
	"errors"
	"fmt"
	"math"
	"time"

	"github.com/pquerna/otp/hotp"

	SmartApi "github.com/angel-one/smartapigo"
)

var (
	abClient    *SmartApi.Client = nil
	userSession SmartApi.UserSession
)

func initiateConnection() error {
	// Create New Angel Broking Client
	// ABClient := SmartApi.New("S452329", "Angel@123", "QOna9C82")
	abClient = SmartApi.New("S452329", "7648", "QOna9C82")

	fmt.Println("Client :- ", abClient)

	totp := generatePassCode("DFVGOUJ4T2MW356CCP5ZR7RAGQ")

	// User Login and Generate User Session
	var err error
	userSession, err = abClient.GenerateSession(totp)

	if err != nil {
		fmt.Println("error creating session. Error: ", err.Error())
		return err
	}

	//Renew User Tokens using refresh token
	userSession.UserSessionTokens, err = abClient.RenewAccessToken(userSession.RefreshToken)

	if err != nil {
		fmt.Println("error in renewing access token. Error: ", err.Error())
		return err
	}

	// defer ABClient.Logout()

	fmt.Println("User Session Tokens :- ", userSession.UserSessionTokens)

	//Get User Profile
	userSession.UserProfile, err = abClient.GetUserProfile()

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("User Profile :- ", userSession.UserProfile)
	fmt.Println("User Session Object :- ", userSession)

	return nil
}

func GetMarginAngel() {

	// // Create New Angel Broking Client
	// // ABClient := SmartApi.New("S452329", "Angel@123", "QOna9C82")
	// ABClient := SmartApi.New("S452329", "7648", "QOna9C82")

	// fmt.Println("Client :- ", ABClient)

	// totp := generatePassCode("DFVGOUJ4T2MW356CCP5ZR7RAGQ")

	// // totp = ""

	// // User Login and Generate User Session
	// session, err := ABClient.GenerateSession(totp)

	// if err != nil {
	// 	fmt.Println("error creating session. Error: ", err.Error())
	// 	return
	// }

	// //Renew User Tokens using refresh token
	// session.UserSessionTokens, err = ABClient.RenewAccessToken(session.RefreshToken)

	// if err != nil {
	// 	fmt.Println("error in renewing access token. Error: ", err.Error())
	// 	return
	// }

	// // defer ABClient.Logout()

	// fmt.Println("User Session Tokens :- ", session.UserSessionTokens)

	// //Get User Profile
	// session.UserProfile, err = ABClient.GetUserProfile()

	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// fmt.Println("User Profile :- ", session.UserProfile)
	// fmt.Println("User Session Object :- ", session)

	// holdings, err := ABClient.GetHoldings()
	// if err != nil {
	// 	fmt.Println("error in fetching holdings", err.Error())
	// }

	// for _, holding := range holdings {
	// 	fmt.Println("isin: ", holding.ISIN, " quantity: ", holding.Quantity, "", " average price: ", holding.AveragePrice, "holding: ", holding)
	// }

	// //Place Order
	// order, err := ABClient.PlaceOrder(SmartApi.OrderParams{Variety: "NORMAL", TradingSymbol: "SBIN-EQ", SymbolToken: "3045", TransactionType: "BUY", Exchange: "NSE", OrderType: "LIMIT", ProductType: "INTRADAY", Duration: "DAY", Price: "19500", SquareOff: "0", StopLoss: "0", Quantity: "1"})

	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// fmt.Println("Placed Order ID and Script :- ", order)
}

func GetHoldingsForAngel() ([]HoldingsInfo, error) {
	if abClient == nil {
		initiateConnection()
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
