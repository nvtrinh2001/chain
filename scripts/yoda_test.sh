#!/bin/bash

rm -rf ~/.yoda
# export EXECUTOR_URL=https://ue3puk0mlg.execute-api.ap-southeast-1.amazonaws.com/default/executor
export EXECUTOR1_URL=http://127.0.0.1:5000/execute
# export EXECUTOR2_URL=http://127.0.0.1:6000/execute

# config chain id
yoda config chain-id bandchain --home ~/.yoda

# add validator to yoda config
yoda config validator $(bandd keys show validator1 -a --bech val --keyring-backend test --keyring-dir ~/.band)  --home ~/.yoda

# setup execution endpoint
# yoda config executor "rest:$EXECUTOR1_URL?timeout=5000s","rest:$EXECUTOR2_URL?timeout=5000s" --home ~/.yoda
yoda config executor "rest:$EXECUTOR1_URL?timeout=5000s&lang=python" --home ~/.yoda

# setup broadcast-timeout to yoda config
yoda config broadcast-timeout "100m" --home ~/.yoda

# setup rpc-poll-interval to yoda config
yoda config rpc-poll-interval "1s" --home ~/.yoda

# setup max-try to yoda config
yoda config max-try 5 --home ~/.yoda

echo "y" | bandd tx oracle activate --from validator1 --keyring-backend test --chain-id bandchain --home ~/.band --keyring-dir ~/.band

# wait for activation transaction success
sleep 2

for i in $(eval echo {1..1}); do
	# add reporter key
	yoda keys add reporter$i --home ~/.yoda
done

# send band tokens to reporters
echo "y" | bandd tx bank send validator1 $(yoda keys list -a --home ~/.yoda) 1000000uband --keyring-backend test --chain-id bandchain  --keyring-dir ~/.band --home ~/.band

# wait for sending band tokens transaction success
sleep 2

# add reporter to bandchain
echo "y" | bandd tx oracle add-reporters $(yoda keys list -a --home ~/.yoda) --from validator1 --keyring-backend test --chain-id bandchain --keyring-dir ~/.band --home ~/.band

# wait for addding reporter transaction success
sleep 2

# run yoda
# yoda run --home ~/.yoda
