package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (api *ApiServer) GetBalance(c *gin.Context) {
	contractAddress, found := c.GetQuery("contract")
	if !found {
		c.JSON(http.StatusBadRequest, R("contract must specified"))
		return
	}

	valid := api.client.ContractValid(contractAddress)
	if !valid {
		c.JSON(http.StatusBadRequest, R("not a valid contractAddress"))
		return
	}

	address, found := c.GetQuery("address")
	if !found {
		c.JSON(http.StatusBadRequest, R("address must specified"))
		return
	}

	balanceWrapper, err := api.client.GetBalance(contractAddress, address)
	if err != nil {
		c.JSON(http.StatusInternalServerError, R(fmt.Sprintf("%+v", err)))
	} else {
		c.JSON(http.StatusOK, R(gin.H{
			"balance":      balanceWrapper.Balance,
			"balanceFloat": balanceWrapper.BalanceFloat,
			"decimal":      balanceWrapper.Decimals,
			"name":         balanceWrapper.Name,
			"symbol":       balanceWrapper.Symbol,
		}))
	}
}
