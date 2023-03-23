import { Component, OnInit } from '@angular/core';


@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css'],
})

export class AppComponent implements OnInit {
  title = "Portfullio";
  userNamee: string;

  constructor() {
    this.userNamee = 'OliverP';
  }

  ngOnInit(): void {
    fetch('localhost:3000/users/')
      .then((response) => response.json())
      .then((quotesData) => (this.userNamee = quotesData));
  }
}
