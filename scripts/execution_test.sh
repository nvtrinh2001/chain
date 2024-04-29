# bandd tx oracle create-requirement-file --name python-requirements --description "dependencies for python script" --script ~/bandchain.js/example/mock/requirements.txt --owner band1p40yh3zkmhcv0ecqp3mcazy83sa57rgjp07dun --treasury band1p40yh3zkmhcv0ecqp3mcazy83sa57rgjp07dun --fee 10uband --from requester --keyring-backend test --chain-id bandchain --node http://127.0.0.1:26657

bandd tx oracle create-data-source --name get-transactions --description "get transactions" --script ~/bandchain.js/example/mock/get_transactions.py --owner band1p40yh3zkmhcv0ecqp3mcazy83sa57rgjp07dun --treasury band1p40yh3zkmhcv0ecqp3mcazy83sa57rgjp07dun --fee 10uband --from requester --keyring-backend test --chain-id bandchain --node http://127.0.0.1:26657 --requirement-file-id 1

