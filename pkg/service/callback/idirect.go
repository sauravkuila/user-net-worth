package callback

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (obj *callbackSt) CallbackForIdirectApiSessionKey(c *gin.Context) {
	var (
		request GetIdirectSessionKeyRequest
	)

	if err := c.BindQuery(&request); err != nil {
		log.Println("request params unavailable", err.Error(), c.Request.Header)
		c.JSON(http.StatusBadRequest, &gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &gin.H{
		"data": request.ApiSession,
	})
}
