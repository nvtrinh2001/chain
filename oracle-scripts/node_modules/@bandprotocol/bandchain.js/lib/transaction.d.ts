import { PublicKey } from './wallet';
import Client from './client';
import { BaseMsg } from './message';
import { SignModeMap } from '../proto/cosmos/tx/signing/v1beta1/signing_pb';
import { Fee } from '../proto/cosmos/tx/v1beta1/tx_pb';
export default class Transaction {
    msgs: Array<BaseMsg>;
    accountNum?: number;
    sequence?: number;
    chainId?: string;
    fee: Fee;
    memo: string;
    withMessages(...msg: Array<BaseMsg>): Transaction;
    withSender(client: Client, sender: string): Promise<Transaction>;
    withAccountNum(accountNum: number): Transaction;
    withSequence(sequence: number): Transaction;
    withChainId(chainId: string): Transaction;
    withFee(fee: Fee): Transaction;
    withMemo(memo: string): Transaction;
    private getInfo;
    getSignDoc(publicKey: PublicKey): Uint8Array;
    getTxData(signature: Uint8Array, publicKey: PublicKey, signMode?: SignModeMap[keyof SignModeMap]): Uint8Array;
    getSignMessage(): Uint8Array;
}
