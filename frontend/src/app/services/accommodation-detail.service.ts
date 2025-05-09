import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class AccommodationDetailService {
  private baseUrl = 'http://localhost:3000/accommodation-detail';

  constructor(private http: HttpClient) {  }

  getAccommodationDetail(): Observable<any[]> {
    return this.http.get<any[]>(this.baseUrl);
  }

  getAccommodationDetailByName(name: string): Observable<any[]> {
    return this.http.get<any[]>(`${this.baseUrl}?name=${encodeURIComponent(name)}`);
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
