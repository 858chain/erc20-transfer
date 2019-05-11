package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/858chain/erc20-transfer/ethclient"
	"github.com/gin-gonic/gin"
)

const MINIMUM_AMOUNT = 0

func (api *ApiServer) Transfer(c *gin.Context) {
	contractAddress, found := c.GetQuery("contract")
	if !found {
		c.JSON(http.StatusBadRequest, R("no contract address specified"))
		return
	}

	valid := api.client.ContractValid(contractAddress)
	if !valid {
		c.JSON(http.StatusBadRequest, R("not a valid contract"))
		return
	}

	fromAddress, found := c.GetQuery("from")
	if !found {
		c.JSON(http.StatusBadRequest, R("no from address specified"))
		return
	}

	toAddress, found := c.GetQuery("to")
	if !found {
		c.JSON(http.StatusBadRequest, R("no to address specified"))
		return
	}

	if fromAddress == toAddress {
		c.JSON(http.StatusBadRequest, R("can not transfer to original address"))
		return
	}

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
		c.JSON(http.StatusBadRequest, R(fmt.Sprintf("amount should bigger than %d", MINIMUM_AMOUNT)))
		return
	}

	hash, err := api.client.TokenTranser(&ethclient.TransferRequest{
		ContractAddress: contractAddress,
		FromAddress:     fromAddress,
		ToAddress:       toAddress,
		Amount:          converedAmount,
		//TODO
		Decimals: 18,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, R(fmt.Sprintf("%+v", err)))
	} else {
		c.JSON(http.StatusOK, R(gin.H{"txid": hash}))
	}
}
