import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

import { NgxLoadersCssModule } from 'ngx-loaders-css';
import {
  MatButtonModule,
  MatDividerModule,
  MatFormFieldModule,
  MatInputModule, MatListModule,
  MatSelectModule,
  MatStepperModule,
  MatProgressBarModule
} from '@angular/material';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';

@NgModule({
  declarations: [
    AppComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    HttpClientModule,
    FormsModule,
    ReactiveFormsModule,

    MatButtonModule,
    MatInputModule,
    MatStepperModule,
    MatFormFieldModule,
    MatSelectModule,
    MatDividerModule,
    MatListModule,
    MatProgressBarModule,
    BrowserModule, NgxLoadersCssModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule {
}
