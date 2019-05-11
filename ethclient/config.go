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

type Config struct {
	// rpc addr, should be one of http://, ws://, ipc
	RpcAddr      string
	LogDir       string
	EthWalletDir string

	// password to unlock account
	EthPassword       string
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
		contractName := determineContractNameFromPath(cfgFile.Name())
		cc, err := loadContractConfig(contractName,
			filepath.Join(c.ERC20ContractsDir, cfgFile.Name()))
		if err != nil {
			return err
		}

		if !common.IsHexAddress(cc.Address) {
			return errors.New(fmt.Sprintf("%s address not valid", contractName))
		}

		if len(cc.Abi) == 0 {
			return errors.New(fmt.Sprintf("%s abi not valid", contractName))
		}

		if cc.Decimals <= 0 {
			return errors.New("decimals not valid")
		}

		// cache valid contractConfig
		c.ContractConfigs[contractName] = cc
	}

	return nil
}

// return ContractConfig have the address as `address`
func (c *Config) ContractConfigForAddress(address string) (ContractConfig, bool) {
	for _, cc := range c.ContractConfigs {
		if cc.Address == address {
			return cc, true
		}
	}
	return ContractConfig{}, false
}

// check directory is valid
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

// load ContractConfig from fs
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

func determineContractNameFromPath(path string) string {
	return strings.TrimSuffix(filepath.Base(path), ".json")
}
