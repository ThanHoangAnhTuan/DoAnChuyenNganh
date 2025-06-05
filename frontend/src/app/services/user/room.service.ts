import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { map, Observable } from 'rxjs';
import { GetAccommodationDetailsResponse } from '../../models/manager/accommodation-detail.model';

@Injectable({
  providedIn: 'root'
})
export class RoomService {
  private baseUrl = 'http://localhost:8080/api/v1/accommodation-detail';

  constructor(private http: HttpClient) { }

  getRoomDetailByAccommodationId(id: string): Observable<GetAccommodationDetailsResponse> {
    return this.http.get<GetAccommodationDetailsResponse>(this.baseUrl + '/get-accommodation-details/' + id);
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
