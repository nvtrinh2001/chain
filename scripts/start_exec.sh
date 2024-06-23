bandd tx oracle create-requirement-file --name python-requirements --description "dependencies for python script" --script ~/ddap/requirement-files/requirements.txt --owner band1m5lq9u533qaya4q3nfyl6ulzqkpkhge9q8tpzs --treasury band1m5lq9u533qaya4q3nfyl6ulzqkpkhge9q8tpzs --fee 10uband --from requester1 --keyring-backend test --chain-id bandchain --node http://127.0.0.1:26657 --keyring-dir ~/.band1

sleep 5 

bandd tx oracle create-data-source --name get-transactions --description "get transactions" --script ~/ddap/data-sources/get_transactions.py --owner band1m5lq9u533qaya4q3nfyl6ulzqkpkhge9q8tpzs --treasury band1m5lq9u533qaya4q3nfyl6ulzqkpkhge9q8tpzs --fee 10uband --from requester1 --keyring-backend test --chain-id bandchain --node http://127.0.0.1:26657 --requirement-file-id 1 --keyring-dir ~/.band1

sleep 5

bandd tx oracle create-oracle-script --name gettransactions --description hello --script ~/ddap/oracle-scripts/get_transactions/target/wasm32-unknown-unknown/release/get_transactions.wasm --owner band1m5lq9u533qaya4q3nfyl6ulzqkpkhge9q8tpzs --schema '{rpc:string,start_block:u64,end_block:u64}/{response:string}' --url test.org --node http://127.0.0.1:26657 --chain-id bandchain --from band1m5lq9u533qaya4q3nfyl6ulzqkpkhge9q8tpzs --keyring-backend test --gas 300000 --keyring-dir ~/.band1

sleep 5

# 10 blocks
bandd tx oracle request 1 2 2 --calldata "0000001368747470733a2f2f317270632e696f2f65746800000000000f423600000000000f4240" --node http://127.0.0.1:26657 --client-id "my-client" --fee-limit 100000000000uband --offchain-fee-limit 400000000000uband --from requester2 --chain-id bandchain --keyring-backend test --gas 3000000 --keyring-dir ~/.band1

