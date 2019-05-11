package ethclient

import (
	"context"
	"math/big"
	"time"

	"github.com/858chain/erc20-transfer/utils"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"golang.org/x/crypto/sha3"
)

type TransferRequest struct {
	ContractAddress string
	FromAddress     string
	ToAddress       string
	Amount          float64
	Decimals        int
}

func (c *Client) TokenTranser(tr *TransferRequest) (string, error) {
	hexContractAddress := common.HexToAddress(tr.ContractAddress)
	hexFromAddress := common.HexToAddress(tr.FromAddress)
	hexToAddress := common.HexToAddress(tr.ToAddress)

	nonce, err := c.PendingNonceAt(context.Background(), hexFromAddress)
	if err != nil {
		utils.L.Error(err)
		return "", err
	}

	zeroAmount := big.NewInt(0) // in wei (0 eth)
	gasPrice, err := c.SuggestGasPrice(context.Background())
	if err != nil {
		utils.L.Error(err)
		return "", err
	}

	transferFnSignature := []byte("transfer(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]

	paddedAddress := common.LeftPadBytes(hexToAddress.Bytes(), 32)

	utils.L.Info(tr.Amount)
	amountBig := floatToBigInt(tr.Amount, tr.Decimals)
	paddedAmount := common.LeftPadBytes(amountBig.Bytes(), 32)

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	gasLimit, err := c.EstimateGas(context.Background(), ethereum.CallMsg{
		To:   &hexToAddress,
		Data: data,
	})
	if err != nil {
		utils.L.Error(err)
		return "", err
	}

	chainID, err := c.NetworkID(context.Background())
	if err != nil {
		utils.L.Error(err)
		return "", err
	}

	unloadedAccount := accounts.Account{Address: hexFromAddress}
	err = c.store.TimedUnlock(unloadedAccount, c.config.EthPassword, time.Duration(time.Second*10))
	if err != nil {
		utils.L.Error(err)
		return "", err
	}

	tx := types.NewTransaction(nonce, hexContractAddress, zeroAmount, gasLimit, gasPrice, data)
	signedTx, err := c.store.SignTx(unloadedAccount, tx, chainID)
	if err != nil {
		utils.L.Error(err)
		return "", err
	}

	err = c.SendTransaction(context.Background(), signedTx)
	if err != nil {
		utils.L.Error(err)
		return "", err
	}

	utils.L.Infof("ERC20TokenTranser contractAddress: %s, fromAddress: %s, toAddress: %s with amount %f",
		tr.ContractAddress, tr.FromAddress, tr.ToAddress, tr.Amount)
	utils.L.Infof("txid: %s", signedTx.Hash().Hex())

	return signedTx.Hash().Hex(), nil
}
