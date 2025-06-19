import { HttpClient } from '@angular/common/http';
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
}
