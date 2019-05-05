package ethclient

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"sync"

	"github.com/858chain/token-shout/notifier"
	"github.com/858chain/token-shout/utils"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/google/uuid"
)

type Client struct {
	config *Config

	// rpc client
	rpcClient *rpc.Client
}

func New(config *Config) (*Client, error) {
	client := &Client{
		config: config,
		noti:   notifier.New(),

		lock:         sync.Mutex{},
		balanceCache: make(map[string]float64),
	}

	err := client.connect()
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (c *Client) Start() error {
	errCh := make(chan error, 1)
	ctx := context.Background()

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
