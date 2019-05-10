package ethclient

import (
	"context"
	"errors"
	"os"

	"github.com/858chain/erc20-transfer/utils"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/rpc"
)

type Client struct {
	// client config
	config *Config

	// keystore - used to sign a tx
	store *keystore.KeyStore

	// rpc client
	rpcClient *rpc.Client
}

func New(config *Config) (*Client, error) {
	// keystore directory check
	stat, err := os.Stat(config.EthWalletDir)
	if err != nil {
		return nil, err
	}

	if !stat.IsDir() {
		return nil, errors.New("eth-wallet-dir is not directory")
	}
	store := keystore.NewKeyStore(config.EthWalletDir,
		keystore.StandardScryptN, keystore.StandardScryptP)

	// initialize client
	client := &Client{
		config: config,
		store:  store,
	}
	err = client.connect()
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (c *Client) Ping() error {
	var hexHeight string
	err := c.rpcClient.CallContext(context.Background(), &hexHeight, "eth_blockNumber")
	if err != nil {
		return err
	}

	return nil
}

// connect to rpc endpoint
func (c *Client) connect() (err error) {
	utils.L.Debugf("ethClient connect to %s", c.config.RpcAddr)
	c.rpcClient, err = rpc.Dial(c.config.RpcAddr)
	if err != nil {
		return err
	}

	return nil
}

// valid contractAddress
// - format valid
// - in whitelite
func (c *Client) ContractValid(contractAddress string) bool {
	return true
}
