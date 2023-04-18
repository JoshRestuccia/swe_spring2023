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
  public username = '';
  public password = '';
  public email = '';
  public userItems: IUserItem[] = [];

  constructor(private httpClient: HttpClient) {}

  async ngOnInit() {
    await this.loadUsers();
  }

  async loadUsers() {
    this.userItems = await firstValueFrom(
      this.httpClient.get<IUserItem[]>('/api/users')
    );
  }

  async addUser() {
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
    this.username = '';
    this.password = '';
    this.email = '';
  }
}
