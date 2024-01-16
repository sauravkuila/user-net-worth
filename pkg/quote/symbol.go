package quote

import (
	"fmt"

	"github.com/sauravkuila/portfolio-worth/pkg/utils"
)

func (obj *quoteSt) AddSymbolToken(token string) error {
	_, found := obj.symbolQuoteMap[token]
	if found {
		return nil
	}
	obj.tokenM.Lock()
	defer obj.tokenM.Unlock()
	obj.symbolQuoteMap[token] = utils.SymbolLtp{}
	return nil
}

func (obj *quoteSt) RemoveSymbolToken(token string) error {
	_, found := obj.symbolQuoteMap[token]
	if !found {
		return fmt.Errorf("token not subscribed previously")
	}
	obj.tokenM.Lock()
	defer obj.tokenM.Unlock()
	delete(obj.symbolQuoteMap, token)
	return nil
}

func (obj *quoteSt) GetSymbolLtp(token string) (*utils.SymbolLtp, error) {
	data, found := obj.symbolQuoteMap[token]
	if !found {
		return nil, fmt.Errorf("token not subscribed")
	}
	if (utils.SymbolLtp{} == data) {
		return nil, fmt.Errorf("token data not fetched yet")
	}
	return &data, nil
}
