// This .ts file helps us navigate between different pages

import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LoginComponent } from './login/login.component';
import { MainpageComponent } from './mainpage/mainpage.component';
import { StocksComponent } from './stocks/stocks.component';

const routes: Routes = [
  { path: '', redirectTo: '/mainpage', pathMatch: 'full' }, // Here, the default route is set to the mainpage
  { path: 'mainpage', component: MainpageComponent },
  { path: 'login', component: LoginComponent },
  { path: 'stocks', component: StocksComponent },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule],
})
export class AppRoutingModule {}
