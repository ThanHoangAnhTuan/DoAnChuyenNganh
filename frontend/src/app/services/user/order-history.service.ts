import { Injectable } from '@angular/core';
import { environment } from '../../../environments/environment';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { GetOrdersByUserResponse } from '../../models/manager/order.model';

@Injectable({
  providedIn: 'root'
})
export class OrderHistoryService {
  private apiUrl = `${environment.apiUrl}/order`;

  constructor(private http: HttpClient) { }

  public getOrdersByUser(): Observable<GetOrdersByUserResponse> {
    return this.http.get<GetOrdersByUserResponse>(`${this.apiUrl}/get-orders-by-user`);
  }
}
