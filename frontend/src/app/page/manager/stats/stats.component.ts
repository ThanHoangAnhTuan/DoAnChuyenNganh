import { NgFor, NgIf } from '@angular/common';
import { Component } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { ChartConfiguration, ChartOptions } from 'chart.js';
import { BaseChartDirective } from 'ng2-charts';
import { StatsService } from '../../../services/manager/stat.service';
import { Observable } from 'rxjs';

@Component({
    selector: 'app-stats',
    imports: [FormsModule, NgIf, NgFor, BaseChartDirective],
    templateUrl: './stats.component.html',
    styleUrl: './stats.component.scss',
})
export class StatsComponent {
    selectedMode: string = 'current-year';
    selectedYear: number = new Date().getFullYear();
    selectedMonth: number = new Date().getMonth() + 1;

    months = Array.from({ length: 12 }, (_, i) => ({
        value: i + 1,
        label: `Tháng ${i + 1}`,
    }));

    constructor(private statsService: StatsService) {}

    ngOnInit() {
        this.loadRevenueData();
    }

    updateUI() {
        const now = new Date();
        if (
            this.selectedMode === 'current-year' ||
            this.selectedMode === 'current-month'
        ) {
            this.selectedYear = now.getFullYear();
            this.selectedMonth = now.getMonth() + 1;
        }
        this.loadRevenueData();
    }

    loadRevenueData() {
        switch (this.selectedMode) {
            case 'current-year':
                this.loadMonthlyRevenue(this.statsService.getMonthlyEarnings());
                break;

            case 'custom-year':
                this.loadMonthlyRevenue(
                    this.statsService.getMonthlyEarningsByYear(
                        this.selectedYear
                    )
                );
                break;

            case 'current-month':
                this.loadDailyRevenue(
                    this.statsService.getDailyEarnings(),
                    this.selectedMonth,
                    this.selectedYear
                );
                break;

            case 'custom-month':
                this.loadDailyRevenue(
                    this.statsService.getDailyEarningsByMonth(
                        this.selectedMonth,
                        this.selectedYear
                    ),
                    this.selectedMonth,
                    this.selectedYear
                );
                break;
        }
    }

    private loadMonthlyRevenue(obs: Observable<any>) {
        const data = Array(12).fill(0);
        const labels = this.months.map((m) => m.label);

        obs.subscribe({
            next: (res) => {
                if (res.data) {
                    res.data.forEach(
                        (el: { month: number; total_revenue: string }) => {
                            const index = el.month - 1;
                            if (index >= 0 && index < 12) {
                                data[index] = Number(el.total_revenue);
                            }
                        }
                    );
                }

                this.updateChart(labels, data, 'Doanh thu theo tháng');
            },
            error: (err) =>
                console.error('Failed to load monthly earnings', err),
        });
    }

    private loadDailyRevenue(
        obs: Observable<any>,
        month: number,
        year: number
    ) {
        const daysInMonth = new Date(year, month, 0).getDate();
        const data = Array(daysInMonth).fill(0);
        const labels = Array.from(
            { length: daysInMonth },
            (_, i) => `Ngày ${i + 1}`
        );

        console.log(labels);

        obs.subscribe({
            next: (res) => {
                if (res.data) {
                    res.data.forEach(
                        (el: { day: string; total_revenue: string }) => {
                            const dayNum = parseInt(el.day.split('-')[0], 10);
                            if (dayNum >= 1 && dayNum <= daysInMonth) {
                                data[dayNum - 1] = Number(el.total_revenue);
                            }
                        }
                    );
                }

                this.updateChart(labels, data, 'Doanh thu theo tháng');
            },
            error: (err) => console.error('Failed to load daily earnings', err),
        });
    }

    private updateChart(labels: string[], data: number[], label: string) {
        this.lineChartData = {
            labels,
            datasets: [
                {
                    data,
                    label,
                    borderColor: 'blue',
                    backgroundColor: 'rgba(0,0,255,0.3)',
                    tension: 0.3,
                    fill: true,
                },
            ],
        };
    }

    public lineChartData: ChartConfiguration<'line'>['data'] = {
        labels: [],
        datasets: [],
    };
    public lineChartOptions: ChartOptions<'line'> = {
        responsive: false,
    };
    public lineChartLegend = true;
}
