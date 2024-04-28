import pkg from '@bandprotocol/bandchain.js'
const { Client, Wallet, Message, Coin, Transaction, Fee } = pkg
import fs from 'fs'
import { fileURLToPath } from 'url';
import { dirname } from 'path';
import path from 'path';

const currentFilePath = fileURLToPath(import.meta.url);
const currentDirectory = dirname(currentFilePath);

// Setup the client
const grpcURL = 'http://127.0.0.1:9091'
const client = new Client(grpcURL)

async function createDataSource() {
  // Setup the wallet
  const { PrivateKey } = Wallet
  const mnemonic = process.env.MNEMONIC
  const privateKey = PrivateKey.fromMnemonic("smile stem oven genius cave resource better lunar nasty moon company ridge brass rather supply used horn three panic put venue analyst leader comic")
  const publicKey = privateKey.toPubkey()
  const sender = publicKey.toAddress().toAccBech32()

  // Setup the transaction's properties
  const chainId = await client.getChainId()
  const execPath = path.resolve(currentDirectory, 'get_transactions.py')
  const file = fs.readFileSync(execPath, 'utf8')
  const executable = Buffer.from(file).toString('base64')

  let feeCoin = new Coin()
  feeCoin.setDenom('uband')
  feeCoin.setAmount('50000')

  const requestMessage = new Message.MsgCreateDataSource(
    'get transactions', // Data source name
    executable, // Data source executable
    sender, // Treasury address
    sender, // Owner address
    sender, // Sender address
    [feeCoin], // Fee
    '' // Data source description
  )

  // Construct the transaction
  const fee = new Fee()
  fee.setAmountList([feeCoin])
  fee.setGasLimit(100000)

  const txn = new Transaction()
  txn.withMessages(requestMessage)
  await txn.withSender(client, sender)
  txn.withChainId(chainId)
  txn.withFee(fee)
  txn.withMemo('')

  // Sign the transaction
  const signDoc = txn.getSignDoc(publicKey)
  const signature = privateKey.sign(signDoc)
  const txRawBytes = txn.getTxData(signature, publicKey)

  // Broadcast the transaction
  const sendTx = await client.sendTxBlockMode(txRawBytes)

  return sendTx
}

; (async () => {
  console.log(await createDataSource())
})()
