package mutualfund

import (
	"log"

	"github.com/sauravkuila/portfolio-worth/external"
)

func (obj *mutualfundSt) GetMutualFundHoldingData() ([]external.MfHoldingsInfo, float64, float64, error) {

	//fetch data from db
	holdings, investedVal, err := obj.dbObj.GetMfCentralHoldings()
	if err != nil {
		log.Println("failed to fetch holdings from db", err.Error())
		return nil, 0, 0, err
	}

	var currentVal float64
	for i, holding := range holdings {
		holdingCurrentVal := holding.CurrentVal
		data, err := obj.quoteObj.GetMutualfundNav(holding.Isin)
		if err == nil {
			holdingCurrentVal = data.Nav * holding.Quantity
			holdings[i].Nav = data.Nav
		}
		currentVal += holdingCurrentVal
	}

	return holdings, investedVal, currentVal, nil
}

func (obj *mutualfundSt) GetIsinOfFundHoldings() ([]string, error) {
	//fetch data from db
	holdings, _, err := obj.dbObj.GetMfCentralHoldings()
	if err != nil {
		log.Println("failed to fetch holdings from db", err.Error())
		return nil, err
	}

	var isins []string = make([]string, 0)
	for _, holding := range holdings {
		isins = append(isins, holding.Isin)
	}

	return isins, nil
}
