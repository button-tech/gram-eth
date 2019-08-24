import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';


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

  constructor() {
  }

  ngOnInit() {
    //
  }
}
