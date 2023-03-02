import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { Login} from '../Login';
@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css'],
})
export class LoginComponent implements OnInit {
  public getJsonValue: any;
  public postJsonValue: any;

  ngOnInit(): void {
    fetch('localhost:3000/users/')
    .then((response) => response.json()) 
    .then(console.log);
  }
  constructor (private http: HttpClient) {

  }
  login: Login = {
    username: '',
    password: '',
  };
}
// This entire component is to help the navigate the website to the login page
