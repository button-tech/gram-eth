import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';

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

  constructor(private http: HttpClient) {

  }

  ngOnInit() {
    // web3.currentProvider.enable()
    // window.torus.getPublicAddress('artall64@gmail.com').then( (addr) => {debugger})
    //
  }

  sendEth() {

  }

  sendGrams() {
    // d16c2312004621ff65ba4425d86aee437c8fb2ec7bef96824fe09099158c17ee
    const payload = {
      currency: 'TonTestNet',
      amount: 3,
      address: this.address,
      tokenAddress: null,
      amountInUsd: false,
      description: '',
      webHookUrl: 'http://34.90.64.237:30923/api/ton/eth'
    };

    // console.log(payload);
    this.http.post('https://client.buttonwallet.com/api/TonFastLink/create', payload)
      .subscribe((resp) => {
        const {uuid, botLink} = (resp as any);
        window.open('https://t.me/wallet_test_bot?start=69907838898');
      }, (error) => {
        console.log(error);
      });
  }
}
