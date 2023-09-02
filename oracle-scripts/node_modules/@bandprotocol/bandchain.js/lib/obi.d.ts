/// <reference types="node" />
declare abstract class ObiBase {
    abstract encode(value: any): Buffer;
    abstract decode(buff: Buffer): any[];
}
export declare class ObiInteger extends ObiBase {
    static REGEX: RegExp;
    isSigned: boolean;
    sizeInBytes: number;
    constructor(schema: string);
    encode(value: bigint): Buffer;
    decode(buff: Buffer): any;
}
export declare class ObiBool extends ObiBase {
    static REGEX: RegExp;
    encode(value: boolean): Buffer;
    decode(buff: Buffer): any;
}
export declare class ObiVector extends ObiBase {
    static REGEX: RegExp;
    internalObi: any;
    constructor(schema: string);
    encode(value: any): Buffer;
    decode(buff: Buffer): any[];
}
export declare class ObiStruct extends ObiBase {
    static REGEX: RegExp;
    internalObiKvs: any;
    constructor(schema: string);
    encode(value: any): Buffer;
    decode(buff: Buffer): any;
}
export declare class ObiString extends ObiBase {
    static REGEX: RegExp;
    encode(value: string): Buffer;
    decode(buff: Buffer): any[];
}
export declare class ObiBytes {
    static REGEX: RegExp;
    encode(value: any): Buffer;
    decode(buff: Buffer): any[];
}
export declare class Obi {
    inputObi: ObiBase;
    outputObi: ObiBase;
    constructor(schema: string);
    encodeInput(value: any): Buffer;
    decodeInput(buff: Buffer): any;
    encodeOutput(value: any): Buffer;
    decodeOutput(buff: Buffer): any;
}
export declare class ObiSpec {
    static impls: (typeof ObiInteger | typeof ObiBool | typeof ObiVector | typeof ObiStruct | typeof ObiString)[];
    static fromSpec(schema: string): ObiBase;
}
export {};
