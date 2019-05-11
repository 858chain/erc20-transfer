Get any erc20 token balance of arbitray address and transfer token

## Examples

### return all addresses in wallet
```bash
$ curl -sSL localhost:8081/v1/addresses  | jq  .
{
  "message": [
    "0x3b8830D7EaA1D98FDa4E67A8607877537241d71c",
    "0x6a151f72Ab86afe61232d2368f41115E8c5a5B7B"
  ]
}
```

### Get ERC20 Balance of Address
```bash
$ curl -sSL 'localhost:8081/v1/getbalance?contract=0x722dd3F80BAC40c951b51BdD28Dd19d435762180&address=0x3b8830D7EaA1D98FDa4E67A8607877537241d71c'  | jq  .
{
  "message": {
    "balance": 0,
    "decimal": 18,
    "name": "Test Standard Token",
    "symbol": "TST"
  }
}
```

### transfer erc20 token to target address
```bash
curl -sSL 'localhost:8081/v1/transfer?contract=0x722dd3F80BAC40c951b51BdD28Dd19d435762180&from=0x3b8830D7EaA1D98FDa4E67A8607877537241d71c&to=0x6a151f72ab86afe61232d2368f41115e8c5a5b7b&amount=0.01'  | jq  .
```
