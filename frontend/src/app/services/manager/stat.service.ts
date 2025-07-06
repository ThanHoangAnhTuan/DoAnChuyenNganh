import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { environment } from '../../../environments/environment';
import {
    GetDailyEarningsResponse,
    GetMonthlyEarningsResponse,
} from '../../models/manager/stats.model';

@Injectable({
    providedIn: 'root',
})
export class StatsService {
    private readonly statsUrl = `${environment.apiUrl}/stats`;
    constructor(private http: HttpClient) {}

    getMonthlyEarnings(): Observable<GetMonthlyEarningsResponse> {
        return this.http.get<GetMonthlyEarningsResponse>(`${this.statsUrl}`);
    }

    getDailyEarnings(): Observable<GetDailyEarningsResponse> {
        return this.http.get<GetDailyEarningsResponse>(
            `${this.statsUrl}/daily`
        );
    }

    getDailyEarningsByMonth(
        month: number,
        year: number
    ): Observable<GetDailyEarningsResponse> {
        return this.http.get<GetDailyEarningsResponse>(
            `${this.statsUrl}/daily/${year}/${month}`
        );
    }

    getMonthlyEarningsByYear(
        year: number
    ): Observable<GetMonthlyEarningsResponse> {
        return this.http.get<GetMonthlyEarningsResponse>(
            `${this.statsUrl}/monthly/${year}`
        );
    }

    exportDailyEarningsCSV(month: number, year: number): Observable<Blob> {
        return this.http.get(
            `${this.statsUrl}/export/daily-earnings/csv/${year}/${month}`,
            {
                responseType: 'blob',
                headers: new HttpHeaders({
                    Accept: 'text/csv',
                }),
            }
        );
    }

    exportMonthlyEarningsCSV(year: number): Observable<Blob> {
        return this.http.get(
            `${this.statsUrl}/export/monthly-earnings/csv/${year}`,
            {
                responseType: 'blob',
                headers: new HttpHeaders({
                    Accept: 'text/csv',
                }),
            }
        );
    }

    downloadFile(blob: Blob, filename: string): void {
        const url = window.URL.createObjectURL(blob);
        const link = document.createElement('a');
        link.href = url;
        link.download = filename;
        link.click();
        window.URL.revokeObjectURL(url);
    }
}
