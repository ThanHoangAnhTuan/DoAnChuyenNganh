import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { map, Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class RoomService {
  private baseUrl = 'http://localhost:8080/api/v1/accommodation-detail/get-accommodation-details';

  constructor(private http: HttpClient) { }

  getRoomDetailByAccommodationId(id: string): Observable<any> {
    return this.http.get<any>(`${this.baseUrl}/${id}`);
  }

  // getAccommodationDetailByName(name: string): Observable<any[]> {
  //   return this.http.get<any>(`${this.baseUrl}?name=${encodeURIComponent(name)}`).pipe(
  //     map(response => response.data)
  //   );
  // }

  // getAccommodationDetailByCity(city: string): Observable<any[]> {
  //   return this.http.get<any[]>(`${this.baseUrl}?city=${encodeURIComponent(city)}`);
  // }

  // getAccommodationImagesByName(name: string): Observable<any> {
  //   return this.getAccommodationDetailByName(name).pipe(
  //     map(data => data[0]?.images || [])
  //   );
  // }
}
