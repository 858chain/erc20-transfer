package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (api *ApiServer) GetBalance(c *gin.Context) {
	address, found := c.GetQuery("address")
	if !found {
		c.JSON(http.StatusBadRequest, R("address must specified"))
		return
	}

	fmt.Println(address)

	c.JSON(200, gin.H{
		"message": "healthy",
	})
}
