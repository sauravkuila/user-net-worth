package api

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/sauravkuila/portfolio-worth/external"
	"github.com/sauravkuila/portfolio-worth/pkg/db"
	"github.com/sauravkuila/portfolio-worth/pkg/service"
)

var (
	serverObj *http.Server = nil
)

func StartServer() error {
	fmt.Println("setup connections and start server")

	//init db
	dbObj, err := db.InitDb()
	if err != nil {
		log.Fatal("unable to initialize db", err.Error())
		return err
	}

	serviceObj := service.InitService(dbObj)

	serverObj = &http.Server{
		Addr:    ":8080",
		Handler: getRouter(serviceObj),
	}

	go func() {
		// service connections
		if err := serverObj.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	return nil
}

func ShutDownServer(ctx context.Context) error {
	var err error
	err = db.CloseDb()
	if err != nil {
		return err
	}
	external.LogoutAngel()
	return serverObj.Shutdown(ctx)
}
