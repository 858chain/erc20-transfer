#!/usr/bin/env bash

./bin/erc20-transfer --log-level debug \
  --log-dir /tmp \
  start \
  --http-listen-addr 0.0.0.0:8081 \
  --rpc-addr ws://107.150.126.20:9546 \
  --watch-list eth,usdt,dusd \
  --eth-wallet-dir /tmp/wallets \
  --erc20-contracts-dir /tmp/contracts

