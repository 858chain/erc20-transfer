package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (api *ApiServer) Addresses(c *gin.Context) {
	addresses, err := api.client.Addresses()
	if err != nil {
		c.JSON(http.StatusInternalServerError, R(fmt.Sprintf("%+v", err)))
	} else {
		c.JSON(http.StatusOK, R(addresses))
	}
}
