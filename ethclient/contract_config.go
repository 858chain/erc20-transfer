package ethclient

// ContractConfig describe config info for any contract
type ContractConfig struct {
	Address   string `json:"address"`
	AbiBase64 string `json:"abi"`
	Abi       []byte
}
