package quote

import (
	"sync"

	"github.com/sauravkuila/portfolio-worth/pkg/db"
	"github.com/sauravkuila/portfolio-worth/pkg/utils"
)

type quoteSt struct {
	dbObj          db.DatabaseInterface
	symbolQuoteMap map[string]utils.SymbolLtp
	mfQuoteMap     map[string]utils.MutualFundNav
	tokenM         sync.Mutex
	navM           sync.Mutex
	syncLoop       bool
}

type QuoteInterface interface {
	AddSymbolToken(token string) error
	RemoveSymbolToken(token string) error
	GetSymbolLtp(token string) (*utils.SymbolLtp, error)

	AddMutualfundIsin(isin string) error
	RemoveMutualfundIsin(isin string) error
	GetMutualfundNav(isin string) (*utils.MutualFundNav, error)

	StartLtpNavRoutine(tokens []string, isin []string)
	StopLtpNavRoutine()
}

func NewQuoteInterface(db db.DatabaseInterface) QuoteInterface {
	return &quoteSt{
		dbObj:          db,
		symbolQuoteMap: make(map[string]utils.SymbolLtp),
		mfQuoteMap:     make(map[string]utils.MutualFundNav),
		tokenM:         sync.Mutex{},
		navM:           sync.Mutex{},
		syncLoop:       true,
	}
}
