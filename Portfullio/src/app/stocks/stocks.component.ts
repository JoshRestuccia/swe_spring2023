import { Component } from '@angular/core';
import { Chart } from 'angular-highcharts';
import { MatCardModule } from '@angular/material/card';
import { firstValueFrom } from 'rxjs';

import { HttpClient } from '@angular/common/http';

interface IStockItem {
  symbol: String;
  name: String;
  price: number;
  quantity: number;
  userRefer: number;
}
@Component({
  selector: 'app-stocks',
  templateUrl: './stocks.component.html',
  styleUrls: ['./stocks.component.css'],
})
export class StocksComponent {
  public symbol = '';
  public name = '';
  public price = 0;
  public quantity = 0;

  public stockItems: IStockItem[] = [];

  constructor(private httpClient: HttpClient) {}

  async ngOnInit() {
    //runs loadUsers on initialization of the component
    await this.loadStocks();
  }

  async loadStocks() {
    //sends a get request that returns an array of IuserItem. the userItems array is then
    //set to the returned value
    this.stockItems = await firstValueFrom(
      this.httpClient.get<IStockItem[]>('/api/stocks/1')
    );
  }

  async addStock() {
    //sends post request with the body of user and it's attributes
    this.httpClient
      .post<IStockItem>('/api/stocks/1', {
        symbol: this.symbol,
        name: this.name,
        price: Number(this.price),
        quantity: Number(this.quantity),
        userRefer: 1,
      })
      .subscribe((response) => {
        console.log(response);
        this.loadStocks();
      });
    //clears temp variables since they are connected to the html inputs
    this.symbol = '';
    this.name = '';
    this.price = 0;
    this.quantity = 0;
  }

  title = 'Portfullio';
  lineChart = new Chart({
    chart: {
      type: 'line',
    },
    title: {
      text: 'Your Portfolio the Past Months',
    },
    yAxis: {
      title: {
        text: 'Value in USD',
      },
    },
    credits: {
      enabled: false,
    },
    series: [
      {
        name: 'USD',
        data: [
          14000, 12000, 18000, 19000, 18500, 20000, 23000, 19000, 20040, 23005,
          25500,
        ],
      } as any,
    ],
  });
  pieChart = new Chart({
    chart: {
      type: 'pie',
    },
    tooltip: {
      pointFormat: '{series.name}: <b>{point.percentage:.1f}%</b>',
    },
    credits: {
      enabled: false,
    },
    accessibility: {
      point: {
        valueSuffix: '%',
      },
    },
    plotOptions: {
      pie: {
        allowPointSelect: true,
        cursor: 'pointer',
        dataLabels: {
          enabled: true,
          format: '<b>{point.name}</b>: {point.percentage:.1f} %',
        },
      },
    },
    series: [
      {
        colorByPoint: true,
        type: 'pie',
        data: [
          { name: 'Stocks', y: 40.32, color: 'blue' },
          { name: 'Real Estate', y: 10.2, color: 'red' },
          { name: 'Cryptocurrency', y: 14.2, color: 'orange' },
          { name: 'Bonds', y: 5.8, color: 'green' },
          { name: '401k', y: 29.48, color: 'yellow' },
        ],
      },
    ],
  });
}
