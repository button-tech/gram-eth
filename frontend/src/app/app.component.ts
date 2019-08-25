import { Component, OnDestroy, OnInit } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Web3Provider } from './web3-provider';
import { chainLinkAbi, chainLinkAddress } from './chainlink';
import { swapAddress, swapContractAbi } from './swap-conract';
import { combineLatest, from, interval, merge, Observable, of, Subscription, timer } from 'rxjs';
import { first, map, switchMap, take, tap, timeout } from 'rxjs/operators';

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
export class AppComponent implements OnInit, OnDestroy {
  isLinear = true;
  selectedCurrency: 'Gram' | 'ETH' = 'Gram';
  amount: number;
  address: string;
  email: string;

  balance: number;
  usdBalance: string;
  ton2eth = 0.01;
  eth2ton = 100;
  resolvingAddressByEmail = false;
  enterManually = true;

  private readonly chainLink;
  private readonly swap;
  private subscription: Subscription;

  constructor(private http: HttpClient) {
    this.chainLink = new Web3Provider(chainLinkAbi, chainLinkAddress);
    this.swap = new Web3Provider(swapContractAbi, swapAddress);
  }

  // TODO: move to helper
  private polling(doRequest: () => Promise<any>, interval: number): Observable<any> {
    return timer(0, interval).pipe(
      switchMap(() => {
        return from(doRequest());
      })
    );
  }

  async ngOnInit() {

    const ethBalance$ = this.polling(
      () => this.getEthBalance(),
      3000
    );

    const eth2usd$ = this.polling(
      () => this.requestEth2usd(),
      10000
    );

    // Subscribe to balances and exchange rates
    this.subscription = combineLatest([ethBalance$, eth2usd$]).pipe(
      tap((x) => {
        const [ethBalance, eth2usd] = x;
        this.balance = Number(ethBalance);
        const rate = Number(eth2usd);
        this.usdBalance = (this.balance * rate).toFixed(2);
        // this.eth2ton = rate * 100;
        // this.ton2eth = 1 / this.eth2ton;
      })
    ).subscribe();
  }

  ngOnDestroy(): void {
    this.subscription.unsubscribe();
  }

  // get eth2usd exchange rate from chainlink
  async requestEth2usd(): Promise<string> {
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
      amount: this.amount,
      address: 'd16c2312004621ff65ba4425d86aee437c8fb2ec7bef96824fe09099158c17ee',
      ethAddressToSend: this.address,
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

  enterAddressManually() {
    this.enterManually = true;
    this.address = '';
    this.email = '';
    this.resolvingAddressByEmail = false;
  }

  fetchAddressByEmail() {
    this.enterManually = false;
    this.address = '';

    if (!this.email || this.email.indexOf('@') === -1) {
      return;
    }

    this.resolvingAddressByEmail = true;

    // WARNING: web3 0.20.3
    const sha3 = (window as any).web3.sha3;

    let promise: Promise<string>;
    if ((window as any).torus) {
      // TODO: implement with debounce on keydown with observable
      // const promise = (window as any).torus.getPublicAddress(this.email);
      // Observable.from(promise).pipe();
      promise = (window as any).torus.getPublicAddress(this.email);
    } else {
      // Mock on the eth berlin to avoid conflicts with metamask
      promise = new Promise((resolve, reject) => {
        setTimeout(() => {
          if (sha3(this.email) === '0x342d98173a593ce5bd91af38752537722db74cc8d880fc81b0015c642d8c3f02') {
            // Mock address 1
            resolve('0x8907B733F664903512ce3F40f16fd67Ac5E7225C');
          } else if (sha3(this.email) === '0xc85207082f2dac41fe03915afe2d6431924df99d528d32c37709a57e0c2efaa9') {
            // Mock address 2
            resolve('0x3EFa2B2A67268C8B3F0bF219180021CC62A57a0c');
          }
        }, 1300);
      });
    }

    const p$ = from(promise);
    const t$ = interval(2500).pipe(map(() => ''));

    merge(p$, t$).pipe(
      take(1)
    ).subscribe((addr: any) => {
      this.address = addr;
      this.enterManually = true;
    }, () => {
      this.address = '';
    }, () => {
      this.resolvingAddressByEmail = false;
    });
  }

  resetAddress() {
    this.address = '';
  }
}
