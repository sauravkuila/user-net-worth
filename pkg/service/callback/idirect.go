package callback

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sauravkuila/portfolio-worth/external"
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

	external.SetIDirectSessionKeyFromCallback(fmt.Sprintf("%d", request.ApiSession))

	c.JSON(http.StatusOK, &gin.H{
		"data": request.ApiSession,
	})
}
