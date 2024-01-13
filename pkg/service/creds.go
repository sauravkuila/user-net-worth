package service

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (obj *serviceStruct) UpdateBrokerCred(c *gin.Context) {
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

	// if request.Broker == "angelone" {
	// 	data, err := obj.dbObj.GetBrokerCred("angelone")
	// 	if err != nil {
	// 		log.Println("error fetching cred in db.", err.Error())
	// 		log.Println("error in sync: ", err)
	// 		response.Error = err.Error()
	// 		c.JSON(http.StatusInternalServerError, response)
	// 		return
	// 	}
	// 	err = external.LoginAndSyncAngelOne(data["user_key"].(string), data["pass_key"].(string), data["totp_secret"].(string), data["app_api_key"].(string))
	// 	if err != nil {
	// 		log.Println("error in sync: ", err)
	// 		response.Error = err.Error()
	// 		c.JSON(http.StatusInternalServerError, response)
	// 		return
	// 	}
	// 	response.Data = "update successful"
	// } else if request.Broker == "zerodha" {
	// 	data, err := obj.dbObj.GetBrokerCred("zerodha")
	// 	if err != nil {
	// 		log.Println("error fetching cred in db.", err.Error())
	// 		log.Println("error in sync: ", err)
	// 		c.JSON(http.StatusInternalServerError, response)
	// 		return
	// 	}
	// 	err = external.LoginAndSyncZerodha(data["user_key"].(string), data["pass_key"].(string), data["totp_secret"].(string))
	// 	if err != nil {
	// 		log.Println("error in sync: ", err)
	// 		c.JSON(http.StatusInternalServerError, response)
	// 		return
	// 	}
	// 	response.Data = "update successful"
	// } else if request.Broker == "idirect" {
	// 	response.Error = "implementation pending"
	// 	c.JSON(http.StatusBadRequest, response)
	// 	return
	// } else if request.Broker == "test" {
	// 	data := make(map[string]interface{})
	// 	data["account"] = request.Broker
	// 	data["user_key"] = request.UserKey
	// 	data["pass_key"] = request.PassKey
	// 	data["totp_secret"] = request.TOTPSecret
	// 	data["app_api_key"] = request.AppCode
	// 	data["secret_key"] = request.SecretKey
	// 	err := obj.dbObj.UpdateBrokerCred(data)
	// 	if err != nil {
	// 		log.Println("error fetching cred in db.", err.Error())
	// 		log.Println("error in sync: ", err)
	// 		response.Error = err.Error()
	// 		c.JSON(http.StatusInternalServerError, response)
	// 		return
	// 	}
	// 	response.Data = "update successful"
	// } else {
	// 	response.Error = "invalid broker received"
	// 	c.JSON(http.StatusBadRequest, response)
	// 	return
	// }

	c.JSON(http.StatusOK, response)
}
