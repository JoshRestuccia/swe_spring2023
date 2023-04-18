import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { firstValueFrom } from 'rxjs';
import { NgModule } from '@angular/core';

interface IUserItem {
  username: String;
  password: String;
  email: String;
}

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css'],
})
export class LoginComponent implements OnInit {
  //create variable to temporarily hold values for new objects
  public username = '';
  public password = '';
  public email = '';
  //create an array to hold users to display on the page for debug purposes
  public userItems: IUserItem[] = [];

  constructor(private httpClient: HttpClient) {}

  async ngOnInit() {
    //runs loadUsers on initialization of the component
    await this.loadUsers();
  }

  async loadUsers() {
    //sends a get request that returns an array of IuserItem. the userItems array is then
    //set to the returned value
    this.userItems = await firstValueFrom(
      this.httpClient.get<IUserItem[]>('/api/users')
    );
  }

  async addUser() {
    //sends post request with the body of user and it's attributes
    this.httpClient
      .post<IUserItem>('/api/users', {
        username: this.username,
        password: this.password,
        email: this.email,
      })
      .subscribe((response) => {
        console.log(response);
        this.loadUsers();
      });
    //clears temp variables since they are connected to the html inputs
    this.username = '';
    this.password = '';
    this.email = '';
  }
}
