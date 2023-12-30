package main

import (
	"encoding/base32"
	"fmt"
	"time"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"github.com/sauravkuila/portfolio-worth/api"
)

func main() {
	fmt.Println("Portfolio Valuation Project")
	totp := generatePassCode("DFVGOUJ4T2MW356CCP5ZR7RAGQ")
	fmt.Println(totp)
	api.ReachAngel()
	// api.ReachZerodha()
}

func generatePassCode(utf8string string) string {
	secret := base32.StdEncoding.EncodeToString([]byte(utf8string))
	passcode, err := totp.GenerateCodeCustom(secret, time.Now(), totp.ValidateOpts{
		Period:    30,
		Skew:      1,
		Digits:    otp.DigitsSix,
		Algorithm: otp.AlgorithmSHA256,
	})
	if err != nil {
		panic(err)
	}
	return passcode
}
