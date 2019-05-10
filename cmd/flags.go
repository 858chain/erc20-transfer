package main

import (
	"github.com/urfave/cli"
)

var httpAddrFlag = cli.StringFlag{
	Name:   "http-listen-addr",
	Value:  "0.0.0.0:8001",
	Usage:  "http address of web application",
	EnvVar: "HTTP_LISTEN_ADDR",
}

var logLevelFlag = cli.StringFlag{
	Name:   "log-level",
	Value:  "info",
	Usage:  "default log level",
	EnvVar: "LOG_LEVEL",
}

var logDirFlag = cli.StringFlag{
	Name:   "log-dir",
	EnvVar: "LOG_DIR",
	Value:  "/var/log/",
}

var envFlag = cli.StringFlag{
	Name:   "env",
	Value:  "production",
	EnvVar: "ENV",
}

var rpcAddrFlag = cli.StringFlag{
	Name:   "rpc-addr",
	Value:  "http://192.168.0.101:8545",
	EnvVar: "RPCADDR",
}

var ethWalletDirFlag = cli.StringFlag{
	Name:   "eth-wallet-dir",
	EnvVar: "ETH_WALLET_DIR",
}

var ERC20ContractsDirFlag = cli.StringFlag{
	Name:   "erc20-contracts-dir",
	Value:  "",
	EnvVar: "ERC20_CONTRACTS_DIR",
}

var ethWalletPasswordFlag = cli.StringFlag{
	Name:   "eth-wallet-password",
	EnvVar: "ETH_WALLET_PASSWORD",
}
