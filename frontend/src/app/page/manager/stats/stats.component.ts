import { NgFor, NgIf } from '@angular/common';
import { Component } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { ChartConfiguration, ChartOptions, ChartType } from 'chart.js';
import { BaseChartDirective } from 'ng2-charts';

@Component({
    selector: 'app-stats',
    imports: [FormsModule, NgIf, NgFor, BaseChartDirective],
    templateUrl: './stats.component.html',
    styleUrl: './stats.component.scss',
})
export class StatsComponent {
    selectedMode: string = 'current-year';
    selectedYear: number = new Date().getFullYear();
    selectedMonth: number = new Date().getMonth() + 1; // tháng 1 = 0

    months = [
        { value: 1, label: 'Tháng 1' },
        { value: 2, label: 'Tháng 2' },
        { value: 3, label: 'Tháng 3' },
        { value: 4, label: 'Tháng 4' },
        { value: 5, label: 'Tháng 5' },
        { value: 6, label: 'Tháng 6' },
        { value: 7, label: 'Tháng 7' },
        { value: 8, label: 'Tháng 8' },
        { value: 9, label: 'Tháng 9' },
        { value: 10, label: 'Tháng 10' },
        { value: 11, label: 'Tháng 11' },
        { value: 12, label: 'Tháng 12' },
    ];

    ngOnInit() {
        this.loadRevenueData();
    }

    updateUI() {
        if (this.selectedMode === 'current-year') {
            this.selectedYear = new Date().getFullYear();
        } else if (this.selectedMode === 'current-month') {
            this.selectedYear = new Date().getFullYear();
            this.selectedMonth = new Date().getMonth() + 1;
        }
        this.loadRevenueData();
    }

    loadRevenueData() {
        const year = this.selectedYear;
        const month = this.selectedMonth;

        if (
            this.selectedMode === 'current-year' ||
            this.selectedMode === 'custom-year'
        ) {
            const labels = [
                'Tháng 1',
                'Tháng 2',
                'Tháng 3',
                'Tháng 4',
                'Tháng 5',
                'Tháng 6',
                'Tháng 7',
                'Tháng 8',
                'Tháng 9',
                'Tháng 10',
                'Tháng 11',
                'Tháng 12',
            ];
            const data = [12, 15, 18, 16, 20].map((x) => x * 1_000_000);

            this.lineChartData = {
                labels: labels,
                datasets: [
                    {
                        data: data,
                        label: 'Doanh thu theo tháng',
                        borderColor: 'blue',
                        backgroundColor: 'rgba(0,0,255,0.3)',
                        tension: 0.3,
                        fill: true,
                    },
                ],
            };
        } else {
            const daysInMonth = 30;
            const labels = Array.from(
                { length: daysInMonth },
                (_, i) => `Ngày ${i + 1}`
            );
            const data = Array.from({ length: daysInMonth }, () =>
                Math.floor(Math.random() * 5_000_000)
            );

            this.lineChartData = {
                labels: labels,
                datasets: [
                    {
                        data: data,
                        label: 'Doanh thu theo ngày',
                        borderColor: 'green',
                        backgroundColor: 'rgba(0,255,0,0.3)',
                        tension: 0.3,
                        fill: true,
                    },
                ],
            };
        }
    }

    public lineChartData: ChartConfiguration<'line'>['data'] = {
        labels: [
            'January',
            'February',
            'March',
            'April',
            'May',
            'June',
            'July',
        ],
        datasets: [
            {
                data: [65, 59, 80, 81, 56, 55, 40],
                label: 'Series A',
                fill: true,
                tension: 0.5,
                borderColor: 'black',
                backgroundColor: 'rgba(255,0,0,0.3)',
            },
        ],
    };
    public lineChartOptions: ChartOptions<'line'> = {
        responsive: false,
    };
    public lineChartLegend = true;
}
