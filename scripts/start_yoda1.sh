#!/bin/bash

rm -rf ~/.yoda1
# export EXECUTOR_URL=https://ue3puk0mlg.execute-api.ap-southeast-1.amazonaws.com/default/executor
export EXECUTOR1_URL=http://127.0.0.1:5000/execute

# config chain id
yoda config chain-id bandchain --home ~/.yoda1

# add validator to yoda config
yoda config validator $(bandd keys show validator1 -a --bech val --keyring-backend test --keyring-dir ~/.band1)  --home ~/.yoda1

# setup execution endpoint
yoda config executor "rest:$EXECUTOR1_URL?timeout=500s" --home ~/.yoda1

# setup broadcast-timeout to yoda config
yoda config broadcast-timeout "5m" --home ~/.yoda1

# setup rpc-poll-interval to yoda config
yoda config rpc-poll-interval "1s" --home ~/.yoda1

# setup max-try to yoda config
yoda config max-try 5 --home ~/.yoda1

echo "y" | bandd tx oracle activate --from validator1 --keyring-backend test --chain-id bandchain --home ~/.band1 --keyring-dir ~/.band1 --node http://127.0.0.1:26657

# wait for activation transaction success
sleep 2

for i in $(eval echo {1..1}); do
	# add reporter key
	yoda keys add reporter$i --home ~/.yoda1
done

# send band tokens to reporters
echo "y" | bandd tx bank send validator1 $(yoda keys list -a --home ~/.yoda1) 1000000uband --keyring-backend test --chain-id bandchain  --keyring-dir ~/.band1 --home ~/.band1 --node http://127.0.0.1:26657

# wait for sending band tokens transaction success
sleep 2

# add reporter to bandchain
echo "y" | bandd tx oracle add-reporters $(yoda keys list -a --home ~/.yoda1) --from validator1 --keyring-backend test --chain-id bandchain --keyring-dir ~/.band1 --home ~/.band1 --node http://127.0.0.1:26657

# wait for addding reporter transaction success
sleep 2

# run yoda
yoda run --home ~/.yoda1 --node http://127.0.0.1:26657
