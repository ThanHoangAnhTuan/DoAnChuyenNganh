import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { GetAccommodationResponse } from '../../models/manager/accommodation.model';

@Injectable({
    providedIn: 'root',
})
export class HotelService {
    private apiUrl = 'http://localhost:8080/api/v1/accommodations';
    constructor(private http: HttpClient) {}

    getAccommodationsByCity(
        city: string,
        page: number = 1,
    ): Observable<GetAccommodationResponse> {
        return this.http.get<GetAccommodationResponse>(
            this.apiUrl + '?city=' + city + '&page=' + page,
        );
    }
}
