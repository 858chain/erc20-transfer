package ethclient

import (
	"math/big"
)

func (c *Client) GetBalance(contractAddress, address string) (*big.Int, int, error) {
	return new(big.Int), 18, nil
}
