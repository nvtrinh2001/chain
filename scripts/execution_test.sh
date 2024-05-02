bandd tx oracle create-requirement-file --name python-requirements --description "dependencies for python script" --script ~/bandchain.js/example/mock/requirements.txt --owner band1m5lq9u533qaya4q3nfyl6ulzqkpkhge9q8tpzs --treasury band1m5lq9u533qaya4q3nfyl6ulzqkpkhge9q8tpzs --fee 10uband --from requester1 --keyring-backend test --chain-id bandchain --node http://127.0.0.1:26657 

sleep 5 

bandd tx oracle create-data-source --name get-transactions --description "get transactions" --script ~/bandchain.js/example/mock/get_transactions.py --owner band1m5lq9u533qaya4q3nfyl6ulzqkpkhge9q8tpzs --treasury band1m5lq9u533qaya4q3nfyl6ulzqkpkhge9q8tpzs --fee 10uband --from requester1 --keyring-backend test --chain-id bandchain --node http://127.0.0.1:26657 --requirement-file-id 1

sleep 5

bandd tx oracle create-oracle-script --name gettransactions --description hello --script /home/nvt/bandchain.js/example/mock/get_transactions.wasm --owner band1m5lq9u533qaya4q3nfyl6ulzqkpkhge9q8tpzs --schema '{repeat:u64}/{response:string}' --url test.org --node http://127.0.0.1:26657 --chain-id bandchain --from band1m5lq9u533qaya4q3nfyl6ulzqkpkhge9q8tpzs --keyring-backend test --gas 300000

sleep 5

bandd tx oracle request 1 1 1 --calldata "0000000000000001" --node http://127.0.0.1:26657 --client-id "my-client" --fee-limit 100000000000uband --offchain-fee-limit 400000000000uband --from requester3 --chain-id bandchain --keyring-backend test --gas 3000000
