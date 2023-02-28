import { Component } from '@angular/core';
import { Login } from '../Login';
@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css'],
})
export class LoginComponent {
  login: Login = {
    username: '',
    password: '',
  };
}
// This entire component is to help the navigate the website to the login page
