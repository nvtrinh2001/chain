#!/home/nvt/chain/data-sources/venv python3

from web3 import Web3
import ipfshttpclient
import hashlib
import json

START_BLOCK = 9_999_990
END_BLOCK = 10_000_000


# Share TCP connections until the client session is closed
class IPFSClient:
    def __init__(self):
        self._hash = None
        self._client = ipfshttpclient.connect(session=True)

    def upload(self, content):
        self._hash = self._client.add_json(content)

    def get_hash(self):
        return self._hash

    def get_value(self):
        return self._client.get_json(self._hash)

    def close(self):  # Call this when your done
        self._client.close()


def get_transactions(ipfs_cli):
    try:
        w3 = Web3(Web3.HTTPProvider('https://cloudflare-eth.com'))
        transactions = []

        for block in range(START_BLOCK, END_BLOCK):
            block = w3.eth.get_block(block, full_transactions=True)
            for transaction in block.transactions:
                tx_data = {
                    "hash": transaction.hash.hex(),
                    "from": transaction['from'],
                    "to": transaction['to'],
                    "value": transaction['value'],
                    "gas": transaction['gas'],
                }
                transactions.append(tx_data)

        ipfs_cli.upload(transactions)
        # return hashlib.sha256(json.dumps(transactions, sort_keys=True).encode('utf-8')).hexdigest()

    except Exception as e:
        print('Error retrieving transaction history:', e)


if __name__ == "__main__":
    ipfs_cli = IPFSClient()

    # start_time = time.time()
    get_transactions(ipfs_cli)
    # end_time = time.time()

    hash = ipfs_cli.get_hash()
    ipfs_cli.close()

    # print(hash1, hash2)
    print(hash)
    #verified_content = ipfs_cli.get_value()
    #print(verified_content[0])

    #print("Hash of saved content: ", hash)
    #print("Number of transactions: ", len(verified_content))
    # print("Execution Time: ", end_time - start_time)
