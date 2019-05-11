package ethclient

import (
	"errors"
	"math/big"
	"strings"

	"github.com/858chain/erc20-transfer/utils"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type BalanceWrapper struct {
	Balance      *big.Int
	Decimals     *big.Int
	BalanceFloat float64
	Name         string
	Symbol       string
}

func (c *Client) GetBalance(contractAddress, address string) (*BalanceWrapper, error) {
	bw := new(BalanceWrapper)

	cc, ok := c.config.ContractConfigForAddress(contractAddress)
	if !ok {
		return bw, errors.New("contract config not found")
	}

	parsedAbi, err := abi.JSON(strings.NewReader(string(cc.Abi)))
	if err != nil {
		utils.L.Error(err)
		return bw, err
	}

	boundContract := bind.NewBoundContract(common.HexToAddress(cc.Address),
		parsedAbi, c, nil, nil)

	err = boundContract.Call(nil, &bw.Decimals, "decimals")
	if err != nil {
		utils.L.Error(err)
		return bw, err
	}

	err = boundContract.Call(nil, &bw.Balance, "balanceOf", common.HexToAddress(address))
	if err != nil {
		utils.L.Error(err)
		return bw, err
	}
	bw.BalanceFloat, _ = bigIntToBigFloat(bw.Balance, int(bw.Decimals.Int64())).Float64()

	err = boundContract.Call(nil, &bw.Name, "name")
	if err != nil {
		utils.L.Error(err)
		return bw, err
	}

	err = boundContract.Call(nil, &bw.Symbol, "symbol")
	if err != nil {
		utils.L.Error(err)
		return bw, err
	}

	return bw, nil
}
