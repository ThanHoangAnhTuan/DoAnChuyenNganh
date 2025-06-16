import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { GetAccommodationByIdResponse, GetAccommodationResponse } from '../../models/manager/accommodation.model';

@Injectable({
  providedIn: 'root'
})
export class AccommodationDetailService {
  private baseUrl = 'http://localhost:8080/api/v1/accommodations';

  constructor(private http: HttpClient) { }

  getAllAccommodationDetail(): Observable<GetAccommodationResponse> {
    return this.http.get<GetAccommodationResponse>(this.baseUrl);
  }

  getAccommodationDetailByCity(city: string): Observable<GetAccommodationResponse> {
    return this.http.get<GetAccommodationResponse>(this.baseUrl + '?city=' + city);
  }

  getAccommodationDetailById(id: string): Observable<GetAccommodationByIdResponse> {
    return this.http.get<GetAccommodationByIdResponse>(this.baseUrl + '/' + id);
  }
}
