import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class AccommodationDetailService {
  private baseUrl = 'http://localhost:3000/accommodation-detail';

  constructor(private http: HttpClient) { }

  getAccommodationDetail(): Observable<any> {
    return this.http.get<any[]>(this.baseUrl);
  }

  getAccommodationDetailByCity(city: string): Observable<any> {
    return this.http.get<any>(`${this.baseUrl}/${city}`);
  }
}
