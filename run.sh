#!/usr/bin/env bash

./bin/erc20-transfer --log-level debug \
  --log-dir /tmp \
  --env production \
  start \
  --http-listen-addr 0.0.0.0:8081 \
  --rpc-addr ws://107.150.126.20:9546 \
  --eth-wallet-dir /tmp/wallets \
  --eth-wallet-password password \
  --erc20-contracts-dir /tmp/contracts

