import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { map, Observable } from 'rxjs';
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

    getHotelDetailByCity(city: string): Observable<any[]> {
        return this.http.get<any[]>(this.apiUrl).pipe(
            map((hotels: any) => hotels.filter((hotel: any) => hotel.location?.city === city))
        );
    }

    getHotelDetailByName(name: string): Observable<any[]> {
        return this.http.get<any[]>(this.apiUrl).pipe(
            map((hotels: any) => hotels.filter((hotel: any) => hotel?.name === name))
        );
    }

    getHotelImagesByName(name: string): Observable<any> {
        return this.getHotelDetailByName(name).pipe(
            map(data => data[0]?.images || [])
        );
    }

    getHotelsFilteredClient(filters: { city?: string; name?: string }): Observable<any[]> {
        return this.http.get<any[]>(this.apiUrl).pipe(
            map((hotels: any[]) => {
                return hotels.filter(hotel => {
                    const matchesCity = !filters.city || hotel.location?.city === filters.city;
                    const matchesName = !filters.name || hotel.name === filters.name;
                    return matchesCity && matchesName;
                });
            })
        );
    }
}
