package broker

import (
	"fmt"
	"log"

	"github.com/sauravkuila/portfolio-worth/external"
	"github.com/sauravkuila/portfolio-worth/pkg/utils"
)

func (obj *brokerSt) getHoldingsByBroker(broker string) ([]external.HoldingsInfo, float64, error) {
	var (
		holdings    []external.HoldingsInfo
		investedVal float64
		err         error
	)
	if broker == "angelone" {
		holdings, investedVal, err = obj.dbObj.GetAngelOneHoldings()
		if err != nil {
			log.Println("failed to fetch holdings from db", err.Error())
			return nil, 0, err
		}
	} else if broker == "zerodha" {
		holdings, investedVal, err = obj.dbObj.GetZerodhaHoldings()
		if err != nil {
			log.Println("failed to fetch holdings from db", err.Error())
			return nil, 0, err
		}
	} else if broker == "idirect" {
		holdings, investedVal, err = obj.dbObj.GetIDirectHoldings()
		if err != nil {
			log.Println("failed to fetch holdings from db", err.Error())
			return nil, 0, err
		}
	} else {
		return nil, 0, fmt.Errorf("invalid broker received")
	}
	return holdings, investedVal, nil
}

func (obj *brokerSt) GetAllBrokerHoldings() (map[string]GetSpecificBrokerHoldings, error) {
	var (
		response map[string]GetSpecificBrokerHoldings = make(map[string]GetSpecificBrokerHoldings)
	)

	sources, err := obj.dbObj.GetSupportedSources()
	if err != nil {
		log.Println("unable to identify broker sources", err.Error())
		return nil, err
	}

	for _, broker := range sources["broker"] {
		if holdings, investedVal, err := obj.getHoldingsByBroker(broker); err != nil {
			return nil, err
		} else {
			//fetch tokens from holdings
			var (
				tokens  []string
				currVal float64
			)
			for _, holding := range holdings {
				info, err := utils.GetSymbolInfoFromTradingSymbol(holding.Symbol)
				if err != nil {
					break
				}
				tokens = append(tokens, info.Token)
			}

			if len(tokens) == len(holdings) {
				ltpMap, err := utils.GetSymbolLtpData(tokens)
				if err == nil {
					//analyze ltp and decide currval for holdings
					for _, holding := range holdings {
						info, _ := utils.GetSymbolInfoFromTradingSymbol(holding.Symbol)
						log.Println(ltpMap[info.Token])
						currVal += float64(holding.Quantity) * ltpMap[info.Token].Ltp
						log.Println(currVal)
					}
				}
			}
			response[broker] = GetSpecificBrokerHoldings{
				BrokerName:    broker,
				InvestedValue: investedVal,
				CurrentValue:  currVal,
				Holdings:      holdings,
			}
		}
	}

	return response, nil
}
