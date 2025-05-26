import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';
import { GetAccommodationByIdResponse, GetAccommodationResponse } from '../../models/manager/accommodation.model';

@Injectable({
  providedIn: 'root'
})
export class AccommodationDetailService {
  private baseUrl = 'http://localhost:8080/api/v1/accommodation';

  constructor(private http: HttpClient) { }

  getAllAccommodationDetail(): Observable<GetAccommodationResponse> {
    return this.http.get<GetAccommodationResponse>(this.baseUrl + '/get-accommodations');
  }

  getAccommodationDetailByCity(city: string): Observable<GetAccommodationResponse> {
    return this.http.get<GetAccommodationResponse>(this.baseUrl + '/get-accommodation-by-city/' + city);
  }

  getAccommodationDetailById(id: string): Observable<GetAccommodationByIdResponse> {
    return this.http.get<GetAccommodationByIdResponse>(this.baseUrl + '/get-accommodation-by-id/' + id);
  }
}
