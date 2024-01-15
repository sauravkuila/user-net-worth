package creds

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (obj *credSt) GetSupportedSources(c *gin.Context) {

	sources, err := obj.dbObj.GetSupportedSources()
	if err != nil {
		log.Println("unable to fetch supported brokers", err.Error())
		c.JSON(http.StatusInternalServerError, &gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &gin.H{
		"data": sources,
	})
}
