import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { map, Observable } from 'rxjs';
import { Hotel } from '../../models/user/hotel.model';
import { AccommodationByCityResponse } from '../../models/manager/accommodation.model';

@Injectable({
    providedIn: 'root',
})
export class HotelService {
    private apiUrl =
        'http://localhost:8080/api/v1/accommodation';
    constructor(private http: HttpClient) {}

    getHotels(): Observable<any> {
        return this.http.get<any>(this.apiUrl);
    }

    getHotelDetailByCity(city: string): Observable<AccommodationByCityResponse> {
        return this.http.get<AccommodationByCityResponse>(this.apiUrl + '/get-accommodation-by-city/' + city);
    }
}
