package api

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/sauravkuila/portfolio-worth/external"
	"github.com/sauravkuila/portfolio-worth/pkg/db"
	"github.com/sauravkuila/portfolio-worth/pkg/quote"
	"github.com/sauravkuila/portfolio-worth/pkg/service"
	"github.com/sauravkuila/portfolio-worth/pkg/utils"
)

var (
	serverObj  *http.Server = nil
	serviceObj service.ServiceInterface
	quoteItf   quote.QuoteInterface
)

func StartServer() error {
	fmt.Println("setup connections and start server")

	//init db
	dbObj, err := db.InitDb()
	if err != nil {
		log.Fatal("unable to initialize db", err.Error())
		return err
	}

	//populate symbol map
	if err := utils.PopulateSymbolTokenMap(); err != nil {
		log.Println("failed to parse symbol map", err.Error())
		return err
	}

	//init quote package
	quoteItf = quote.NewQuoteInterface(dbObj)

	//set up service interface
	serviceObj = service.InitService(dbObj, quoteItf)

	serverObj = &http.Server{
		Addr:    ":8080",
		Handler: getRouter(serviceObj),
	}

	//start quotes
	tokens, err := serviceObj.GetBrokerObject().GetHoldingTokenAcrossAllBrokers()
	if err != nil {
		log.Println("unable to start ltp sync due to failed token list fetch")
		return err
	}
	isins, err := serviceObj.GetMutualFundObject().GetIsinOfFundHoldings()
	if err != nil {
		log.Println("unable to start ltp sync due to failed isin list fetch")
		return err
	}
	quoteItf.StartLtpNavRoutine(tokens, isins)

	go func() {
		// service connections
		if err := serverObj.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	return nil
}

func ShutDownServer(ctx context.Context) error {
	err := db.CloseDb()
	if err != nil {
		return err
	}
	external.LogoutAngel()
	quoteItf.StopLtpNavRoutine()
	return serverObj.Shutdown(ctx)
}
