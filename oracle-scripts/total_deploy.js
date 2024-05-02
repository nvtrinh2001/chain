import pkg from '@bandprotocol/bandchain.js'
const { Client, Wallet, Message, Coin, Transaction, Fee } = pkg
import fs from 'fs'
import path from 'path';

// Setup the client
const grpcURL = 'http://127.0.0.1:9091'
const client = new Client(grpcURL)

async function createOracleScript() {
  // Setup the wallet
  const { PrivateKey } = Wallet
  const privateKey = PrivateKey.fromMnemonic("smile stem oven genius cave resource better lunar nasty moon company ridge brass rather supply used horn three panic put venue analyst leader comic")
  const publicKey = privateKey.toPubkey()
  const sender = publicKey.toAddress().toAccBech32()

  // Setup the transaction's properties
  const chainId = await client.getChainId()
  const execPath = path.resolve(currentDirectory, 'total_supply/target/wasm32-unknown-unknown/release/total_supply.wasm')
  const code = fs.readFileSync(execPath)

  let feeCoin = new Coin()
  feeCoin.setDenom('uband')
  feeCoin.setAmount('0')

  const requestMessage = new Message.MsgCreateOracleScript(
    'Total Supply', // oracle script name
    code, // oracle script code
    sender, // owner
    sender, // sender
    '', // description
    '{repeat:u64}/{response:string}', // schema
    'https://ipfs.io/ipfs/QmU5B8ag1KeNARqcdkGVyGcVs6Lq8WB5SmFcKo787LYQM6' // source code url
  )

  // Construct the transaction
  const fee = new Fee()
  fee.setAmountList([feeCoin])
  fee.setGasLimit(350000)

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
  console.log(await createOracleScript())
})()