import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { Hotel } from '../../models/user/hotel.model';

@Injectable({
    providedIn: 'root',
})
export class HotelService {
    private apiUrl =
        'http://localhost:8080/api/v1/accommodation/get-accommodations';
    constructor(private http: HttpClient) {}

    getHotels(): Observable<any> {
        return this.http.get<any>(this.apiUrl);
    }
}
