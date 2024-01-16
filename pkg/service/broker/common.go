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
		var (
			currVal float64 = 0
		)
		if holdings, investedVal, err := obj.getHoldingsByBroker(broker); err != nil {
			return nil, err
		} else {

			for i, holding := range holdings {
				info, err := utils.GetSymbolInfoFromTradingSymbol(holding.Symbol)
				if err != nil {
					currVal = 0
					break
				}
				if d, err := obj.quoteObj.GetSymbolLtp(info.Token); err != nil {
					currVal = 0
					break
				} else {
					currVal += d.Ltp * float64(holding.Quantity)
					holdings[i].Ltp = d.Ltp
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

func (obj *brokerSt) GetHoldingTokenAcrossAllBrokers() ([]string, error) {
	var (
		tokens []string = make([]string, 0)
	)
	sources, err := obj.dbObj.GetSupportedSources()
	if err != nil {
		log.Println("unable to identify broker sources", err.Error())
		return nil, err
	}

	for _, broker := range sources["broker"] {
		if holdings, _, err := obj.getHoldingsByBroker(broker); err != nil {
			return nil, err
		} else {
			//fetch tokens from holdings
			for _, holding := range holdings {
				info, err := utils.GetSymbolInfoFromTradingSymbol(holding.Symbol)
				if err != nil {
					log.Printf("no data for trading symbol. symbol: %s, broker: %s\n", holding.Symbol, broker)
					return nil, err
				}
				tokens = append(tokens, info.Token)
			}
		}
	}
	return tokens, nil
}
