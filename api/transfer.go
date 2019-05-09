package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const MINIMUM_AMOUNT = 0

func (api *ApiServer) Transfer(c *gin.Context) {
	toAddress, found := c.GetQuery("to")
	if !found {
		c.JSON(http.StatusBadRequest, R("no to address specified"))
		return
	}
	fmt.Println(toAddress)

	amount, found := c.GetQuery("amount")
	if !found {
		c.JSON(http.StatusBadRequest, R("no amount specified"))
		return
	}

	converedAmount, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, R("amount must be a valid float64"))
		return
	}

	if converedAmount <= MINIMUM_AMOUNT {
		c.JSON(http.StatusBadRequest, R(fmt.Sprintf("amount should bigger than %f", MINIMUM_AMOUNT)))
		return
	}

	c.JSON(200, gin.H{
		"message": "healthy",
	})
}
