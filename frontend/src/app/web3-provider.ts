export const chainLinkAbi = [{
  constant: false,
  inputs: [{name: 'newCurrent', type: 'bytes32'}],
  name: 'update',
  outputs: [],
  type: 'function'
}, {
  constant: true,
  inputs: [],
  name: 'current',
  outputs: [{name: 'current', type: 'bytes32'}],
  type: 'function'
}, {inputs: [], type: 'constructor'}];

export const chainLinkAddress = '0x2a032eb0af76e6a0315f14b470f1fbe309393416';

export class Web3Provider {

  readonly contractInstance: any;
  readonly web3: any;

  constructor(contractAbi: any, contractAddress: any) {
    const eth = (window as any).web3.eth;
    const prepare = eth.contract(contractAbi);
    this.contractInstance = prepare.at(contractAddress);
    this.web3 = (window as any).web3;
  }

  public tbn(n) {
    return new this.web3.BigNumber(n);
  }

  public tw(n) {
    return this.tbn(n).mul(1e18).toString();
  }

  public fw(n) {
    return this.tbn(n).div(1e18).toString();
  }

  private isUnlocked = () => typeof this.web3.currentProvider !== 'undefined';

  public getAddress() {
    if (!this.isUnlocked()) {
      alert('Please, unlock your MetaMask account');
      throw new Error('Please, unlock your MetaMask account');
    }
    return this.web3.currentProvider.selectedAddress;
  }

  public getBalance(address: string) {
    if (!this.isUnlocked()) {
      alert('Please, unlock your MetaMask account');
      throw new Error('Please, unlock your MetaMask account');
    }
    return new Promise((resolve, reject) => {
      this.web3.eth.getBalance(address, (err, balance) => {
        if (err) {
          reject(err);
        }
        resolve(balance.toString());
      });
    });
  }

  public sendTransaction(toAddress: any, amount: any, data: any = '') {
    return new Promise((resolve, reject) => {
      this.web3.eth.sendTransaction({
        to: toAddress,
        value: amount,
        data
      }, (err, txHash) => {
        if (err !== null) {
          reject(err);
        }
        resolve(txHash);
      });
    });
  }

  public callSmartContract(methodName: string) {
    return new Promise((resolve, reject) => {
      this.contractInstance[methodName].call((err, res) => {
        if (err) {
          reject(err);
        }
        resolve(this.fromHexToString(res));
      });
    });
  }

  private fromHexToString(hexString) {
    const hex = hexString.toString();
    let str = '';
    for (let n = 0; n < hex.length; n += 2) {
      str += String.fromCharCode(parseInt(hex.substr(n, 2), 16));
    }
    return str;
  }

  public sendSmartContract(methodName: string, parameters: any[] = [], value: string = '') {
    return new Promise((resolve, reject) => {
      this.contractInstance.methods[methodName](...parameters).send({value}, (err, res) => {
        if (err) {
          reject(err);
        }
        resolve(res);
      });
    });
  }

  public getCallData(methodName: string, parameters: any[] = []) {
    if (!this.contractInstance.methods[methodName]) {
      throw new Error(`Method ${methodName} does not exist`);
    }
    return this.contractInstance.methods[methodName](...parameters).encodeABI();
  }

}
