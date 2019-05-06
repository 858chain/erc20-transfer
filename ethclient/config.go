package ethclient

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
)

type ContractConfig struct {
	Address   string `json:"address"`
	AbiBase64 string `json:"abi"`
	Abi       []byte
}

type Config struct {
	// rpc addr, should be one of http://, ws://, ipc
	RpcAddr           string
	LogDir            string
	EthWalletDir      string
	ERC20ContractsDir string
	ContractConfigs   map[string]ContractConfig
}

// Check config is valid.
func (c *Config) SanityAndValidCheck() error {
	if len(c.RpcAddr) == 0 {
		return errors.New("RpcAddr should not empty")
	}

	// rpcaddr format check
	if !(strings.HasPrefix(c.RpcAddr, "http://") ||
		strings.HasPrefix(c.RpcAddr, "ws://") ||
		strings.HasSuffix(c.RpcAddr, ".ipc")) {
		return errors.New("rpcaddr should like http://, ws:// or /xxx/xx/foo.ipc")
	}

	if len(c.EthWalletDir) == 0 {
		return errors.New("WalletDir should not empty")
	}

	err := isValidDir(c.EthWalletDir)
	if err != nil {
		return err
	}

	err = isValidDir(c.LogDir)
	if err != nil {
		return err
	}

	err = isValidDir(c.ERC20ContractsDir)
	if err != nil {
		return err
	}

	c.ContractConfigs = make(map[string]ContractConfig)

	erc20Configs, err := ioutil.ReadDir(c.ERC20ContractsDir)
	if err != nil {
		return err
	}

	for _, cfgFile := range erc20Configs {
		name := filepath.Base(cfgFile.Name())
		cc, err := loadContractConfig(name, cfgFile.Name())
		if err != nil {
			return err
		}

		if !common.IsHexAddress(cc.Address) {
			return errors.New(fmt.Sprintf("%s address not valid", name))
		}

		if len(cc.Abi) == 0 {
			return errors.New(fmt.Sprintf("%s abi not valid", name))
		}
		c.ContractConfigs[name] = cc
	}

	return nil
}

func isValidDir(dir string) error {
	stat, err := os.Stat(dir)
	if err != nil {
		return errors.Wrap(err, dir)
	}

	if !stat.IsDir() {
		return errors.New(fmt.Sprintf("%s is not a directory", dir))
	}

	return nil
}

func loadContractConfig(name, path string) (ContractConfig, error) {
	_, err := os.Stat(path)
	if err != nil {
		return ContractConfig{}, errors.Wrap(err, name)
	}

	contractFile, err := os.Open(path)
	if err != nil {
		return ContractConfig{}, errors.Wrap(err, path)
	}
	defer contractFile.Close()

	var contractConfig ContractConfig
	err = json.NewDecoder(contractFile).Decode(&contractConfig)
	if err != nil {
		return ContractConfig{}, err
	}

	contractConfig.Abi, err = base64.StdEncoding.DecodeString(contractConfig.AbiBase64)
	if err != nil {
		return ContractConfig{}, err
	}

	return contractConfig, nil
}
