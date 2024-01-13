package service

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sauravkuila/portfolio-worth/external"
)

func (obj *serviceStruct) UpdateBrokerCred(c *gin.Context) {
	var (
		request  UpdateBrokerCredRequest
		response UpdateBrokerCredResponse
	)
	if err := c.BindUri(&request); err != nil {
		log.Println("bad request", err.Error())
		response.Error = err.Error()
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if request.Broker == "angelone" {
		response.Error = "implementation pending"
		c.JSON(http.StatusBadRequest, response)
		return
	} else if request.Broker == "zerodha" {
		data, err := obj.dbObj.GetBrokerCred()
		if err != nil {
			log.Println("error fetching cred in db.", err.Error())
			log.Println("error in sync: ", err)
			c.JSON(http.StatusInternalServerError, response)
			return
		}
		err = external.LoginAndSyncZerodha(data["user_key"].(string), data["pass_key"].(string), data["totp_secret"].(string))
		if err != nil {
			log.Println("error in sync: ", err)
			c.JSON(http.StatusInternalServerError, response)
			return
		}
		response.Data = "update successful"
	} else if request.Broker == "idirect" {
		response.Error = "implementation pending"
		c.JSON(http.StatusBadRequest, response)
		return
	} else {
		response.Error = "invalid broker received"
		c.JSON(http.StatusBadRequest, response)
		return
	}

	c.JSON(http.StatusOK, response)
}
