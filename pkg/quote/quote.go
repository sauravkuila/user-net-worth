package quote

import (
	"time"

	"github.com/sauravkuila/portfolio-worth/pkg/utils"
)

func (obj *quoteSt) StartLtpNavRoutine(tokens []string, isins []string) {
	//add  to map
	for _, token := range tokens {
		obj.symbolQuoteMap[token] = utils.SymbolLtp{}
	}
	for _, isin := range isins {
		obj.mfQuoteMap[isin] = utils.MutualFundNav{}
	}
	//waiting 2s until brokers have logged in
	time.Sleep(2 * time.Second)

	//start go routines for fetching ltp asynchronously
	go func() {
		for obj.syncLoop {
			keys := make([]string, 0, len(obj.symbolQuoteMap))
			for k := range obj.symbolQuoteMap {
				keys = append(keys, k)
			}

			data, err := utils.GetSymbolLtpData(keys)
			if err == nil {
				obj.tokenM.Lock()
				obj.symbolQuoteMap = data
				obj.tokenM.Unlock()
			}
			time.Sleep(30 * time.Second)
		}
	}()
	go func() {
		for obj.syncLoop {
			for k := range obj.mfQuoteMap {
				data, err := utils.GetNavValueFromIsin(k)
				if err == nil {
					obj.navM.Lock()
					obj.mfQuoteMap[k] = *data
					obj.navM.Unlock()
				}
			}
			time.Sleep(10 * time.Minute)
		}
	}()
}

func (obj *quoteSt) StopLtpNavRoutine() {
	obj.syncLoop = false
}
