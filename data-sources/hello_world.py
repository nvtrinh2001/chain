#!/usr/bin/env python3

import sys
import json
from web3 import Web3

abi = [
    {
        "constant": True,
        "inputs": [],
        "name": "name",
        "outputs": [{"name": "", "type": "string"}],
        "payable": False,
        "stateMutability": "view",
        "type": "function"
    },
    {
        "constant": False,
        "inputs": [{"name": "guy", "type": "address"}, {"name": "wad", "type": "uint256"}],
        "name": "approve",
        "outputs": [{"name": "", "type": "bool"}],
        "payable": False,
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "constant": True,
        "inputs": [],
        "name": "totalSupply",
        "outputs": [{"name": "", "type": "uint256"}],
        "payable": False,
        "stateMutability": "view",
        "type": "function"
    },
    {
        "constant": False,
        "inputs": [{"name": "src", "type": "address"}, {"name": "dst", "type": "address"},
                   {"name": "wad", "type": "uint256"}],
        "name": "transferFrom",
        "outputs": [{"name": "", "type": "bool"}],
        "payable": False,
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "constant": False,
        "inputs": [{"name": "wad", "type": "uint256"}],
        "name": "withdraw",
        "outputs": [],
        "payable": False,
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "constant": True,
        "inputs": [],
        "name": "decimals",
        "outputs": [{"name": "", "type": "uint8"}],
        "payable": False,
        "stateMutability": "view",
        "type": "function"
    },
    {
        "constant": True,
        "inputs": [{"name": "", "type": "address"}],
        "name": "balanceOf",
        "outputs": [{"name": "", "type": "uint256"}],
        "payable": False,
        "stateMutability": "view",
        "type": "function"
    },
    {
        "constant": True,
        "inputs": [],
        "name": "symbol",
        "outputs": [{"name": "", "type": "string"}],
        "payable": False,
        "stateMutability": "view",
        "type": "function"
    },
    {
        "constant": False,
        "inputs": [{"name": "dst", "type": "address"}, {"name": "wad", "type": "uint256"}],
        "name": "transfer",
        "outputs": [{"name": "", "type": "bool"}],
        "payable": False,
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "constant": False,
        "inputs": [],
        "name": "deposit",
        "outputs": [],
        "payable": True,
        "stateMutability": "payable",
        "type": "function"
    },
    {
        "constant": True,
        "inputs": [{"name": "", "type": "address"}, {"name": "", "type": "address"}],
        "name": "allowance",
        "outputs": [{"name": "", "type": "uint256"}],
        "payable": False,
        "stateMutability": "view",
        "type": "function"
    },
    {"payable": True, "stateMutability": "payable", "type": "fallback"},
    {
        "anonymous": False,
        "inputs": [
            {"indexed": True, "name": "src", "type": "address"},
            {"indexed": True, "name": "guy", "type": "address"},
            {"indexed": False, "name": "wad", "type": "uint256"}
        ],
        "name": "Approval",
        "type": "event"
    },
    {
        "anonymous": False,
        "inputs": [
            {"indexed": True, "name": "src", "type": "address"},
            {"indexed": True, "name": "dst", "type": "address"},
            {"indexed": False, "name": "wad", "type": "uint256"}
        ],
        "name": "Transfer",
        "type": "event"
    },
    {
        "anonymous": False,
        "inputs": [
            {"indexed": True, "name": "dst", "type": "address"},
            {"indexed": False, "name": "wad", "type": "uint256"}
        ],
        "name": "Deposit",
        "type": "event"
    },
    {
        "anonymous": False,
        "inputs": [
            {"indexed": True, "name": "src", "type": "address"},
            {"indexed": False, "name": "wad", "type": "uint256"}
        ],
        "name": "Withdrawal",
        "type": "event"
    }
]


def main():
    address = "0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c"
    provider_url = "https://bsc-dataseed1.ninicoin.io"
    w3 = Web3(Web3.HTTPProvider(provider_url))
    token_contract = w3.eth.contract(address=address, abi=abi)

    name = token_contract.functions.name().call()
    symbol = token_contract.functions.symbol().call()
    total_supply = token_contract.functions.totalSupply().call()

    result = name + "," + symbol + "," + str(total_supply)
    # result = "Hello"
    return result


# def main():
#     return "Hello World!"


if __name__ == "__main__":
    try:
        print(main())
    except Exception as e:
        print(str(e), file=sys.stderr)
        sys.exit(1)
