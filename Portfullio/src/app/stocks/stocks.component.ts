import { Component } from '@angular/core';
import { Chart } from 'angular-highcharts';
@Component({
  selector: 'app-stocks',
  templateUrl: './stocks.component.html',
  styleUrls: ['./stocks.component.css'],
})
export class StocksComponent {
  title = 'Portfullio';
  lineChart = new Chart({
    chart: {
      type: 'line',
    },
    title: {
      text: 'Sample Stock Chart',
    },
    credits: {
      enabled: false,
    },
    series: [
      {
        name: 'Price',
        data: [140, 120, 180, 190, 185, 200, 230, 190, 240, 235, 255],
      } as any,
    ],
  });
}
