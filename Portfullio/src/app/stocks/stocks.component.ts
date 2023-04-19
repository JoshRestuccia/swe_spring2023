import { Component } from '@angular/core';
import { Chart } from 'angular-highcharts';
import { MatCardModule } from '@angular/material/card';
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
