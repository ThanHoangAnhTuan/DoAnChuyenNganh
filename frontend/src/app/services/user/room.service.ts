import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { map, Observable } from 'rxjs';
import { GetAccommodationDetailsResponse } from '../../models/manager/accommodation-detail.model';

@Injectable({
    providedIn: 'root',
})
export class RoomService {
    private baseUrl = 'http://localhost:8080/api/v1/accommodation-detail';

    constructor(private http: HttpClient) {}

    getRoomDetailByAccommodationId(
        id: string,
        checkIn: string,
        checkOut: string
    ): Observable<GetAccommodationDetailsResponse> {
        return this.http.get<GetAccommodationDetailsResponse>(
            `${this.baseUrl}/get-accommodation-details/${id}?check_in=${checkIn}&check_out=${checkOut}`
        );
    }
}
