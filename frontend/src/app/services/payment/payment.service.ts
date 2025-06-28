import { Injectable } from '@angular/core';
import { environment } from '../../../environments/environment';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
    providedIn: 'root',
})
export class PaymentService {
    private apiUrl = `${environment.apiUrl}/payment`;

    constructor(private http: HttpClient) {}

    createPayment(): Observable<any> {
        return this.http.post<any>(`${this.apiUrl}/create_payment_url`, {
            amount: 1000000,
        });
    }
}
