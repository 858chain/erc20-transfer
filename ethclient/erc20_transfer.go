package ethclient

import (
	"context"
	"math/big"
	"strings"

	"github.com/858chain/erc20-transfer/utils"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
)

func (c *Client) ERC20TokenTranser(contractAddress, toAddress string, amount float64) {
}
