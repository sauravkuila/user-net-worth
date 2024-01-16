package quote

import (
	"fmt"

	"github.com/sauravkuila/portfolio-worth/pkg/utils"
)

func (obj *quoteSt) AddMutualfundIsin(isin string) error {
	_, found := obj.mfQuoteMap[isin]
	if found {
		return nil
	}
	obj.navM.Lock()
	defer obj.navM.Unlock()
	obj.mfQuoteMap[isin] = utils.MutualFundNav{}
	return nil
}
func (obj *quoteSt) RemoveMutualfundIsin(isin string) error {
	_, found := obj.mfQuoteMap[isin]
	if !found {
		return fmt.Errorf("isin not subscribed previously")
	}
	obj.navM.Lock()
	defer obj.navM.Unlock()
	delete(obj.mfQuoteMap, isin)
	return nil
}
func (obj *quoteSt) GetMutualfundNav(isin string) (*utils.MutualFundNav, error) {
	data, found := obj.mfQuoteMap[isin]
	if !found {
		return nil, fmt.Errorf("isin not subscribed")
	}
	if (utils.MutualFundNav{} == data) {
		return nil, fmt.Errorf("token data not fetched yet")
	}
	return &data, nil
}
