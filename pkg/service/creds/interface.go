package creds

import (
	"github.com/gin-gonic/gin"
	"github.com/sauravkuila/portfolio-worth/pkg/db"
)

type credSt struct {
	dbObj db.DatabaseInterface
}

type CredsInterface interface {
	GetSupportedSources(c *gin.Context)
	UpdateBrokerCred(c *gin.Context)
	UpdateMutualFundCred(c *gin.Context)
}

func NewCredsInterfaceObject(db db.DatabaseInterface) CredsInterface {
	return &credSt{
		dbObj: db,
	}
}
