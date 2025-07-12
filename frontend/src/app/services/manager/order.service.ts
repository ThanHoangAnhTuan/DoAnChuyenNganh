import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { environment } from '../../../environments/environment';
import { Observable } from 'rxjs';
import { GetOrderChageStatusResponse, GetOrdersByManagerResponse } from '../../models/manager/order.model';

@Injectable({
    providedIn: 'root',
})
export class OrderService {
    private apiUrl = `${environment.apiUrl}/order`;

    constructor(private http: HttpClient) {}

    getOrdersByManager(): Observable<GetOrdersByManagerResponse> {
        return this.http.get<GetOrdersByManagerResponse>(`${this.apiUrl}/get-orders-by-manager`);
    }

    changeOrderStatus(orderId: string, status: string): Observable<GetOrderChageStatusResponse> {
        const url = this.apiUrl + '/' + status;
        const body = { order_id: orderId };
        
        return this.http.post<GetOrderChageStatusResponse>(url, body);
    }
}
