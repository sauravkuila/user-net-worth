package creds

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (obj *credSt) UpdateBrokerCred(c *gin.Context) {
	var (
		request  UpdateBrokerCredRequest
		response UpdateBrokerCredResponse
	)
	if err := c.BindJSON(&request); err != nil {
		log.Println("bad request", err.Error())
		response.Error = err.Error()
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := make(map[string]interface{})
	data["account"] = request.Broker
	data["user_key"] = request.UserKey
	data["pass_key"] = request.PassKey
	data["totp_secret"] = request.TOTPSecret
	data["app_api_key"] = request.AppCode
	data["secret_key"] = request.SecretKey
	err := obj.dbObj.UpdateBrokerCred(data)
	if err != nil {
		log.Println("error fetching cred in db.", err.Error())
		log.Println("error in sync: ", err)
		response.Error = err.Error()
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response.Data = "update successful"

	c.JSON(http.StatusOK, response)
}
