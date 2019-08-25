import { Component, OnDestroy, OnInit } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Web3Provider } from './web3-provider';
import { chainLinkAbi, chainLinkAddress } from './chainlink';
import { swapAddress, swapContractAbi } from './swap-conract';
import { combineLatest, from, interval, merge, Observable, Subscription, timer } from 'rxjs';
import { map, switchMap, take, tap } from 'rxjs/operators';
import { LoadersCSS } from 'ngx-loaders-css';
import { supportedCurrency, supportedCurrencyList } from './currencies';
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

  srcList = [...supportedCurrencyList];
  dstList = [...supportedCurrencyList];

  bwLink: string;

  // tslint:disable-next-line:variable-name
  private _dstCurrency: supportedCurrency;
  get dstCurrency(): supportedCurrency {
    return this._dstCurrency;
  }

  set dstCurrency(value: supportedCurrency) {
    this._dstCurrency = value;
    this.updateCryptoExchangeRate();
  }

  // tslint:disable-next-line:variable-name
  private _srcCurrency: supportedCurrency;
  get srcCurrency(): supportedCurrency {
    return this._srcCurrency;
  }

  set srcCurrency(value: supportedCurrency) {
    this._srcCurrency = value;
    this.dstList = supportedCurrencyList.filter(x => x !== value);
    this.dstCurrency = this.dstList[0];
    this.updateCryptoExchangeRate();
  }

  loader: LoadersCSS = 'line-scale';
  bgColor = 'white';
  color = 'rgb(63, 81, 181) ';
  isSent = false;

  isLinear = true;

  amount: number;
  address: string;
  email: string;

  balance: number;
  usdBalance: string;

  eth2ton = 1000;
  ton2eth = 1 / this.eth2ton;

  resolvingAddressByEmail = false;
  enterManually = true;

  private readonly chainLink;
  private readonly swap;
  private subscription: Subscription;
  exchangeRate = 0;

  // (srcCurrency == 'ETH' ? eth2ton : ton2eth)

  constructor(private http: HttpClient) {
    this.chainLink = new Web3Provider(chainLinkAbi, chainLinkAddress);
    this.swap = new Web3Provider(swapContractAbi, swapAddress);
    this.srcCurrency = 'Gram';
    this.dstCurrency = 'ETH';
  }

  // TODO: move to helper
  private polling(doRequest: () => Promise<any>, pollingPeriod: number): Observable<any> {
    return timer(0, pollingPeriod).pipe(
      switchMap(() => {
        return from(doRequest());
      })
    );
  }

  updateCryptoExchangeRate() {
    if (!this.srcCurrency || !this.dstCurrency) {
      this.exchangeRate = 0;
      return;
    }

    if (this.srcCurrency === 'ETH' && this.dstCurrency === 'Gram') {
      this.exchangeRate = this.eth2ton;
      return;
    }

    if (this.srcCurrency === 'Gram' && this.dstCurrency === 'ETH') {
      this.exchangeRate = this.ton2eth;
      return;
    }

    //
    // this.dstCurrency && fetchExchangeRate(this.srcCurrency && this.dstCurrency)
    //
  }

  ngOnInit() {

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
        this.bwLink = botLink;
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

  refresh(): void {
    this.isSent = true;
    setTimeout(() => {
      window.location.reload();
    }, 10000);

  }

  send() {
    if (this.srcCurrency === 'ETH' && this.dstCurrency === 'Gram') {
      this.sendEth();
      // TODO: other ETH pairs
    } else if (this.srcCurrency === 'Gram' && this.dstCurrency === 'ETH') {
      this.sendGrams();
      // TODO: other Gram pairs
    } else if (this.srcCurrency === 'BNB') {
      // TODO: ....
    } else if (this.srcCurrency === 'Waves' && this.dstCurrency === 'Gram') {
      this.sendWaves();
    }
    this.refresh();
  }

  private sendBnb() {

  }

  private sendWaves() {
    (window as any).Waves.signAndPublishTransaction({
      type: 4, // 4 - transfer transaction
      data: {
        amount: {
          assetId: 'WAVES',
          tokens: this.amount
        },
        fee: {
          assetId: 'WAVES',
          tokens: '0.01'
        },
        recipient: '3N2TA9QQ11dmq1eM3khsWF1bTWNG3s4xoHo'
      }
    });

    this.http.post('/api/waves/ton', null).subscribe();
  }
}
