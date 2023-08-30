import pkg from '@bandprotocol/bandchain.js'
const { Client, Wallet, Message, Coin, Transaction, Fee, Obi } = pkg
import fs from 'fs'
import { fileURLToPath } from 'url';
import { dirname } from 'path';
import path from 'path';

const currentFilePath = fileURLToPath(import.meta.url);
const currentDirectory = dirname(currentFilePath);

// Setup the client
const grpcURL = 'http://127.0.0.1:9091'
const client = new Client(grpcURL)

async function makeRequest() {
  // Setup the wallet
  const { PrivateKey } = Wallet
  const privateKey = PrivateKey.fromMnemonic("smile stem oven genius cave resource better lunar nasty moon company ridge brass rather supply used horn three panic put venue analyst leader comic")
  const publicKey = privateKey.toPubkey()
  const sender = publicKey.toAddress().toAccBech32()

  // Step 2.1: Prepare oracle request's properties
  const obi = new Obi('{repeat:u64}/{response:string}')
  const calldata = obi.encodeInput({ repeat: 1 })

  const oracleScriptId = 1
  const askCount = 1
  const minCount = 1
  const clientId = 'from_bandchain.js'

  let feeLimit = new Coin()
  feeLimit.setDenom('uband')
  feeLimit.setAmount('100000')

  const prepareGas = 100000
  const executeGas = 200000

  // Step 2.2: Create an oracle request message
  const requestMessage = new Message.MsgRequestData(
    oracleScriptId,
    calldata,
    askCount,
    minCount,
    clientId,
    sender,
    [feeLimit],
    prepareGas,
    executeGas
  )

  let feeCoin = new Coin()
  feeCoin.setDenom('uband')
  feeCoin.setAmount('50000')

  // Step 3.1: Construct a transaction
  const fee = new Fee()
  fee.setAmountList([feeCoin])
  fee.setGasLimit(1000000)

  const chainId = await client.getChainId()
  const txn = new Transaction()
  txn.withMessages(requestMessage)
  await txn.withSender(client, sender)
  txn.withChainId(chainId)
  txn.withFee(fee)
  txn.withMemo('')

  // Step 3.2: Sign the transaction using the private key
  const signDoc = txn.getSignDoc(publicKey)
  const signature = privateKey.sign(signDoc)

  const txRawBytes = txn.getTxData(signature, publicKey)

  // Step 4: Broadcast the transaction
  const sendTx = await client.sendTxBlockMode(txRawBytes)
  console.log(sendTx)
}

; (async () => {
  await makeRequest()
})()
