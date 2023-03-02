import { Component, OnInit } from '@angular/core';


@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css'],
})
export class AppComponent implements OnInit {
  title = 'Portfullio';
  private url: string = 'https://type.fit/api/quotes';

  ngOnInit(): void {
    fetch(this.url)
    .then((response) => response.json()) 
    .then(console.log);
  }
}
