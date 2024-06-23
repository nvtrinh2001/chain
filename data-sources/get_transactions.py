import sys
import requests
import json
from decimal import Decimal

# Function to get transactions for a given block range
def get_transactions(rpc, start_block, end_block):
    try:
        headers = {
            'Content-Type': 'application/json',
        }
        transactions = []

        for block_num in range(start_block, end_block):
            payload = {
                "jsonrpc": "2.0",
                "method": "eth_getBlockByNumber",
                "params": [hex(block_num), True],
                "id": 1,
            }
            response = requests.post(rpc, headers=headers, json=payload)
            
            # Check if response is valid JSON
            try:
                block = response.json()['result']
                if block is None:
                    print(f"Block {block_num} is empty or does not exist.")
                    continue
            except json.JSONDecodeError:
                print(f"Error decoding JSON for block {block_num}. Raw response: {response.text}")
                continue

            for transaction in block['transactions']:
                tx_data = {
                    "hash": transaction['hash'],
                    "from": transaction['from'],
                    "to": transaction.get('to', None),
                    "value": float(Decimal(int(transaction['value'], 16)) / 10**18),  # Convert Wei to Ether
                    "gas": int(transaction['gas'], 16),
                }
                transactions.append(tx_data)

        return transactions

    except requests.RequestException as e:
        print(f"Network error: {e}")
        return None
    except Exception as e:
        print(f"Error retrieving transaction history: {e}")
        return None

# Function to perform analysis on transactions and return results as JSON
def analyze_transactions(transactions):
    if transactions is None:
        return None

    # Calculate total Ether transferred
    total_ether_transferred = sum(tx['value'] for tx in transactions)

    # Identify top 5 senders by transaction volume
    sender_totals = {}
    for tx in transactions:
        if tx['from'] in sender_totals:
            sender_totals[tx['from']] += tx['value']
        else:
            sender_totals[tx['from']] = tx['value']

    # top_senders = sorted(sender_totals.items(), key=lambda item: item[1], reverse=True)[:5]

    # Prepare results as JSON
    results = {
        "total_ether_transferred": total_ether_transferred,
        # "top_senders": [{"sender": sender, "total_amount": amount} for sender, amount in top_senders],
        # "transactions": transactions
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

