package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (api *ApiServer) Addresses(c *gin.Context) {
	c.JSON(http.StatusOK, R(api.client.Addresses()))

}
