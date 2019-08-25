import { Pipe, PipeTransform } from '@angular/core';
import { supportedCurrency } from './currencies';

export const mapping: { [key in supportedCurrency]: string } = {
  ETH: 'ethereum',
  Gram: 'TON',
  BNB: 'binance',
  Waves: 'waves',
  '': ''
};

@Pipe({
  name: 'symbol2name'
})
export class Symbol2namePipe implements PipeTransform {

  transform(value: any, ...args: any[]): any {
    return mapping[value];
  }

}
