import sys
from web3 import Web3
import pandas as pd
import json
from decimal import Decimal

# Function to get transactions for a given block range
def get_transactions(rpc, start_block, end_block):
    try:
        w3 = Web3(Web3.HTTPProvider(rpc))
        transactions = []

        for block_num in range(start_block, end_block):
            block = w3.eth.get_block(block_num, full_transactions=True)
            for transaction in block.transactions:
                tx_data = {
                    "hash": transaction.hash.hex(),
                    "from": transaction['from'],
                    "to": transaction['to'],
                    "value": float(w3.from_wei(transaction['value'], 'ether')),  # Convert Wei to Ether and then to float
                    "gas": transaction['gas'],
                }
                transactions.append(tx_data)

        return transactions

    except Exception as e:
        print('Error retrieving transaction history:', e)
        return None

# Function to perform analysis on transactions and return results as JSON
def analyze_transactions(transactions):
    if transactions is None:
        return None
    
    # Convert transactions to DataFrame for easier manipulation
    df = pd.DataFrame(transactions)

    # Calculate total Ether transferred
    total_ether_transferred = float(df['value'].sum())  # Convert to float for JSON serialization

    # Identify top 5 senders by transaction volume
    top_senders = df.groupby('from')['value'].sum().sort_values(ascending=False).head(5)

    # Prepare results as JSON
    results = {
        "total_ether_transferred": total_ether_transferred,
        "top_senders": top_senders.reset_index().rename(columns={'from': 'sender', 'value': 'total_amount'}).to_dict(orient='records'),
        "transactions": df.to_dict(orient='records')
    }

    return json.dumps(results, indent=4)  # Convert results dict to JSON string with indentation

if __name__ == "__main__":
    if len(sys.argv) != 4:
        print("Usage: python script.py <rpc_url> <start_block> <end_block>")
        sys.exit(1)

    rpc_url = sys.argv[1]
    start_block = int(sys.argv[2])
    end_block = int(sys.argv[3])

    transactions = get_transactions(rpc_url, start_block, end_block)
    if transactions:
        results_json = analyze_transactions(transactions)
        if results_json:
            print(results_json)

