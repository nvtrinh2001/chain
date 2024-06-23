#!/bin/bash

rm -rf ~/.yoda2
# export EXECUTOR_URL=https://ue3puk0mlg.execute-api.ap-southeast-1.amazonaws.com/default/executor
export EXECUTOR_URL=http://127.0.0.1:7000/execute

# config chain id
yoda config chain-id bandchain --home ~/.yoda2

# add validator to yoda config
yoda config validator $(bandd keys show validator2 -a --bech val --keyring-backend test --keyring-dir ~/.band1)  --home ~/.yoda2

# setup execution endpoint
yoda config executor "rest:$EXECUTOR_URL?timeout=5000s" --home ~/.yoda2

# setup broadcast-timeout to yoda config
yoda config broadcast-timeout "100m" --home ~/.yoda2

# setup rpc-poll-interval to yoda config
yoda config rpc-poll-interval "1s" --home ~/.yoda2

# setup max-try to yoda config
yoda config max-try 5 --home ~/.yoda2

echo "y" | bandd tx oracle activate --from validator2 --keyring-backend test --chain-id bandchain --home ~/.band2 --keyring-dir ~/.band1 --node http://127.0.0.1:26657

# wait for activation transaction success
sleep 2

for i in $(eval echo {1..1}); do
	# add reporter key
	yoda keys add reporter$i --home ~/.yoda2
done

# send band tokens to reporters
echo "y" | bandd tx bank send validator2 $(yoda keys list -a --home ~/.yoda2) 1000000uband --keyring-backend test --chain-id bandchain  --keyring-dir ~/.band1 --home ~/.band2 --node http://127.0.0.1:26657

# wait for sending band tokens transaction success
sleep 2

# add reporter to bandchain
echo "y" | bandd tx oracle add-reporters $(yoda keys list -a --home ~/.yoda2) --from validator2 --keyring-backend test --chain-id bandchain --keyring-dir ~/.band1 --home ~/.band2 --node http://127.0.0.1:26657

# wait for addding reporter transaction success
sleep 2

# run yoda
yoda run --home ~/.yoda2 --node http://127.0.0.1:26657
