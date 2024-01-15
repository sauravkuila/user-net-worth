package callback

import (
	"github.com/gin-gonic/gin"
	"github.com/sauravkuila/portfolio-worth/pkg/db"
)

type callbackSt struct {
	dbObj db.DatabaseInterface
}

type CallbackInterface interface {
	CallbackForIdirectApiSessionKey(c *gin.Context)
}

func NewCallbackInterface(db db.DatabaseInterface) CallbackInterface {
	return &callbackSt{
		dbObj: db,
	}
}
