import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { CreatePayment, Payment } from '../../models/user/payment.model';
import { HttpClient, HttpHeaders } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class PaymentService {
  private baseUrl = 'http://localhost:8080/api/v1/payment/';

  constructor(private http: HttpClient) { }

  createPayment(payment: Payment): Observable<any> {
    console.log("payment: ", payment);

    return this.http.post<any>(
      'http://localhost:8080/api/v1/payment/create-payment-url',
      payment,
      { observe: 'response' }
    );
  }
}
