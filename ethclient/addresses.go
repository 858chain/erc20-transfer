package ethclient

import (
	"github.com/ethereum/go-ethereum/accounts"
)

func (c *Client) Addresses() []string {
	accounts := make([]accounts.Account, 0)
	addresses := make([]string, 0)

	for _, wallet := range c.store.Wallets() {
		accounts = append(accounts, wallet.Accounts()...)
	}

	for _, account := range accounts {
		addresses = append(addresses, account.Address.Hex())
	}

	return addresses
}
