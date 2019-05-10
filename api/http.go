package api

import (
	"fmt"
	"strings"

	"github.com/858chain/erc20-transfer/ethclient"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

// http api list
var METHODS_SUPPORTED = map[string]string{
	// misc
	"/ping":   "check if api service valid and backend bitcoin service healthy",
	"/health": "check system status",
	"/help":   "display this message",
}

type ApiServer struct {
	httpListenAddr string
	// http gin engine instance
	engine *gin.Engine

	// eth rpc client
	client *ethclient.Client
}

// InitEthClient do the config  validation for make initial call to eth backend.
// Error return if malformat config or rpc server unreachable.
func (api *ApiServer) InitAndStartEthClient(cfg *ethclient.Config) (err error) {
	api.client, err = ethclient.New(cfg)
	return err
}

// Check eth rpc server connectivity.
func (api *ApiServer) HealthCheck() (err error) {
	err = api.client.Ping()
	if err != nil {
		err = errors.Wrap(err, "erc20: ")
	}

	return err
}

// Hook all HTTP routes and start listen on `addr`
func NewApiServer(addr, env string) *ApiServer {
	apiServer := &ApiServer{
		httpListenAddr: addr,
	}

	if strings.ToLower(env) == "production" ||
		strings.ToLower(env) == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	g := r.Group("/v1")
	g.GET("/addresses", apiServer.Addresses)
	g.GET("/transfer", apiServer.Transfer)
	g.GET("/list_balances", apiServer.ListBalances)

	// misc API
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/health", func(c *gin.Context) {
		err := apiServer.HealthCheck()
		if err != nil {
			c.JSON(500, gin.H{
				"message": fmt.Sprint(err),
			})
		} else {
			c.JSON(200, gin.H{
				"message": "healthy",
			})
		}
	})

	r.GET("/help", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"methods": METHODS_SUPPORTED,
		})
	})

	apiServer.engine = r
	return apiServer
}

func (api *ApiServer) HttpListen() error {
	return api.engine.Run(api.httpListenAddr)
}
