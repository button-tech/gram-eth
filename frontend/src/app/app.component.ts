import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';


@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent implements OnInit {
  isLinear = true;
  firstFormGroup: FormGroup;
  secondFormGroup: FormGroup;

  selectedCurrency = 'TON';
  currencies = ['TON', 'ETH'];
  // currencies = [
  //   {viewValue: 'TON', value: 'ton' },
  //   {viewValue: 'ETH', value: 'eth' },
  // ];

  constructor(private formBuilder: FormBuilder) {
  }

  ngOnInit() {

    this.firstFormGroup = this.formBuilder.group({
      amountCtrl: ['', Validators.required],
      currencyCtrl: ['', Validators.required],
    });

    this.firstFormGroup.controls.currencyCtrl.setValue('TON');


    this.secondFormGroup = this.formBuilder.group({
      secondCtrl: ['', Validators.required]
    });
  }
}
