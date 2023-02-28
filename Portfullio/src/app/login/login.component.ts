import { HttpClient } from '@angular/common/http';
import { Component } from '@angular/core';
import { Login } from '../Login';
@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css'],
})
export class LoginComponent {
  public getJsonValue: any;
  public postJsonValue: any;
  constructor (private http: HttpClient) {

  }
  login: Login = {
    username: '',
    password: '',
  };
}
// This entire component is to help the navigate the website to the login page
