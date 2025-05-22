import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class AccommodationDetailService {
  private baseUrl = 'http://localhost:8081/api/v1/accommodation/get-accommodations';

  constructor(private http: HttpClient) { }

  getAccommodationDetail(): Observable<any[]> {
    return this.http.get<any>(this.baseUrl).pipe(
      map(response => response.data)
    );
  }

  getAccommodationDetailByName(name: string): Observable<any[]> {
    return this.http.get<any>(`${this.baseUrl}?name=${encodeURIComponent(name)}`).pipe(
      map(response => response.data)
    );
  }

  getAccommodationDetailByCity(city: string): Observable<any[]> {
    return this.http.get<any[]>(`${this.baseUrl}?city=${encodeURIComponent(city)}`);
  }

  getAccommodationImagesByName(name: string): Observable<any> {
    return this.getAccommodationDetailByName(name).pipe(
      map(data => data[0]?.images || [])
    );
  }
}
