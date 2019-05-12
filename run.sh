#!/usr/bin/env bash

./bin/erc20-transfer --log-level debug \
  --log-dir /tmp \
  --env development \
  start \
  --http-listen-addr 0.0.0.0:8081 \
  --rpc-addr http://154.8.201.160:8545 \
  --eth-wallet-dir /tmp/wallets \
  --eth-unlock-password password \
  --erc20-contracts-dir /tmp/contracts

