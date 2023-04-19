import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { MaterialModule } from './material/material.module';
import { LoginComponent } from './login/login.component';
import { MainpageComponent } from './mainpage/mainpage.component';
import { UserhomeComponent } from './userhome/userhome.component';
import { StocksComponent } from './stocks/stocks.component';
import { ChartModule } from 'angular-highcharts';

@NgModule({
  declarations: [
    AppComponent,
    LoginComponent,
    MainpageComponent,
    UserhomeComponent,
    StocksComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    MaterialModule,
    FormsModule,
    HttpClientModule,
    ReactiveFormsModule,
    ChartModule,
  ],
  providers: [],
  bootstrap: [AppComponent],
})
export class AppModule {}
