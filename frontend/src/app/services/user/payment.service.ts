import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Payment } from '../../models/user/payment.model';
import { HttpClient } from '@angular/common/http';

@Injectable({
    providedIn: 'root',
})
export class PaymentService {
    private baseUrl = 'http://localhost:8080/api/v1/payment/';

    constructor(private http: HttpClient) {}

    createPayment(payment: Payment): Observable<any> {
        return this.http.post<any>(
            'http://localhost:8080/api/v1/payment/create-payment-url',
            payment,
            { observe: 'response' }
        );
    }
}
