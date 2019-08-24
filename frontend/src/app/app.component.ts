import { Component, OnInit } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Web3Provider } from './web3-provider';
import { chainLinkAbi, chainLinkAddress } from "./chainlink";
import { swapAddress, swapContractAbi } from "./swap-conract";

//
// import Web3 from 'web3';
// import Torus from '@toruslabs/torus-embed';
// const web3Obj = {
//   web3: new Web3(),
//   setweb3: (provider) => {
//     web3Obj.web3 = new Web3(provider);
//     sessionStorage.setItem('pageUsingTorus', 'true');
//   },
//   initialize: async () => {
//     const torus = new Torus();
//     await torus.init();
//     await torus.login();
//     web3Obj.setweb3(torus.provider);
//   }
// };
// // export default web3Obj
//
// web3Obj.enable();
// web3Obj.web3.enable();

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent implements OnInit {
  isLinear = true;
  selectedCurrency = 'TON';
  amount: number;
  address: string;

  balance: number;
  usdBalance: string;

  private readonly chainLink;
  private readonly swap;

  constructor(private http: HttpClient) {
    this.chainLink = new Web3Provider(chainLinkAbi, chainLinkAddress);
    this.swap = new Web3Provider(swapContractAbi, swapAddress);
  }

  async ngOnInit() {
    // web3.currentProvider.enable()
    // window.torus.getPublicAddress('artall64@gmail.com').then( (addr) => {debugger})
    //

    //
    // http://35.228.96.248/getBalance/{address}/?network=basechain
    // http://35.228.96.248/getBalance/2dc356e6c07379ae86c09fadd6ba1f858ec65bab0252f4e36c05c5ff73b9806c/?network=basechain
    //

    // this.getExchangeRate();
    // this.getBalance();
    this.balance = await this.getEthBalance();
    const rate = await this.eth2usd();
    this.usdBalance = (Number(this.balance) * Number(rate)).toFixed(2);
  }

  // get eth2usd exchange rate from chainlink
  async eth2usd(): Promise<string> {
    return await this.chainLink.callSmartContract('current');
  }

  async getEthBalance(): Promise<number> {
    const addr = this.chainLink.getAddress();
    if (!addr) {
      console.log(`Probably you didn't share address with host, check your metamask`);
      return 0;
    }
    // console.log(addr);
    const wei = await this.chainLink.getBalance(addr);
    return +this.chainLink.fw(wei);
  }

  async sendEth() {
    const tonAddress = this.address;
    const amount = this.swap.tw(this.amount);
    const txHash = await this.swap.sendSmartContract('sendTon', [tonAddress], amount);
    console.log(txHash); // show tx hash in UI
  }

  sendGrams() {
    //
    // Sample tx address
    // d16c2312004621ff65ba4425d86aee437c8fb2ec7bef96824fe09099158c17ee
    //
    const payload = {
      currency: 'TonTestNet',
      amount: 3,
      address: this.address,
      tokenAddress: null,
      amountInUsd: false,
      description: '',
      webHookUrl: 'http://34.90.64.237:30923/api/ton/eth'
    };

    let headers = new HttpHeaders();
    headers = headers.set('Authorization', '2abfac55-18a0-4d5d-bb61-47175743b606');

    this.http.post('https://client.buttonwallet.com/api/TonFastLink/create', payload, {headers})
      .subscribe((resp) => {
        const {uuid, botLink} = (resp as any);
        window.open(botLink);
      }, (error) => {
        console.log(error);
      });
  }

  fetchEmail() {

  }
}
