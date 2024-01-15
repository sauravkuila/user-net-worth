package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

var (
	syminfoMap map[string]SymbolInfo = make(map[string]SymbolInfo)
)

func PopulateSymbolTokenMap() error {

	// Open our jsonFile
	jsonFile, err := os.Open("../external/OpenAPIScripMaster.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Println(err.Error())
		return err
	}
	log.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	data, err := io.ReadAll(jsonFile)
	if err != nil {
		return err
	}
	// var symbolMap []map[string]interface{}
	var symbolMap []SymbolInfo
	if err := json.Unmarshal(data, &symbolMap); err != nil {
		log.Println("failed to marshal", err.Error())
		return err
	}
	for _, symInfo := range symbolMap {
		tradingSym := strings.Split(symInfo.TradingSymbol, "-")[0]
		symInfo.TradingSymbol = tradingSym
		syminfoMap[tradingSym] = symInfo
		// log.Println(symInfo["token"])
		// fmt.Printf("exchg segment type %T\n", symInfo["exch_seg"])
		// fmt.Printf("expiry type %T\n", symInfo["expiry"])
		// fmt.Printf("instrumenttype type %T\n", symInfo["instrumenttype"])
		// fmt.Printf("lotsize type %T\n", symInfo["lotsize"])
		// fmt.Printf("name type %T\n", symInfo["name"])
		// fmt.Printf("lotsize type %T\n", symInfo["lotsize"])
		// fmt.Printf("strike type %T\n", symInfo["strike"])
		// fmt.Printf("symbol type %T\n", symInfo["symbol"])
		// fmt.Printf("tick_size type %T\n", symInfo["tick_size"])
		// fmt.Printf("token type %T\n", symInfo["token"])
	}
	return nil
}

func GetSymbolInfoFromTradingSymbol(tradingSymbol string) (*SymbolInfo, error) {
	tradingSym := strings.Split(tradingSymbol, "-")[0]
	data, found := syminfoMap[tradingSym]
	if !found {
		log.Println("data not found")
		return nil, fmt.Errorf("data not found")
	}
	return &data, nil
}
