package ethclient

import (
	"context"

	"github.com/858chain/erc20-transfer/utils"

	"github.com/ethereum/go-ethereum/rpc"
)

type Client struct {
	config *Config

	// rpc client
	rpcClient *rpc.Client
}

func New(config *Config) (*Client, error) {
	client := &Client{
		config: config,
	}

	err := client.connect()
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (c *Client) Start() error {
	errCh := make(chan error, 1)
	//ctx := context.Background()

	return <-errCh
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
