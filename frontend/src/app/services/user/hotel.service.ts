import { GetAccommodationByIdResponse } from './../../models/manager/accommodation.model';
import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { map, Observable } from 'rxjs';
import { Hotel } from '../../models/user/hotel.model';
import { GetAccommodationResponse } from '../../models/manager/accommodation.model';

@Injectable({
    providedIn: 'root',
})
export class HotelService {
    private apiUrl = 'http://localhost:8080/api/v1/accommodations';
    constructor(private http: HttpClient) {}

    // getHotels(): Observable<any> {
    //     return this.http.get<any>(this.apiUrl + '/get-accommodations');
    // }

    getAccommodationsByCity(
        city: string
    ): Observable<GetAccommodationResponse> {
        return this.http.get<GetAccommodationResponse>(
            this.apiUrl + '?city=' + city
        );
    }

    // getHotelDetailByCity(city: string): Observable<any[]> {
    //     return this.http
    //         .get<any[]>(this.apiUrl + '/get-accommodation-by-city/' + city)
    // }

    // getHotelDetailByName(name: string): Observable<any[]> {
    //     return this.http
    //         .get<any[]>(this.apiUrl)
    //         .pipe(
    //             map((hotels: any) =>
    //                 hotels.filter((hotel: any) => hotel?.name === name)
    //             )
    //         );
    // }

    // getHotelImagesByName(name: string): Observable<any> {
    //     return this.getHotelDetailByName(name).pipe(
    //         map((data) => data[0]?.images || [])
    //     );
    // }

    // getHotelsFilteredClient(filters: {
    //     city?: string;
    //     name?: string;
    // }): Observable<any[]> {
    //     return this.http.get<any[]>(this.apiUrl).pipe(
    //         map((hotels: any[]) => {
    //             return hotels.filter((hotel) => {
    //                 const matchesCity =
    //                     !filters.city || hotel.location?.city === filters.city;
    //                 const matchesName =
    //                     !filters.name || hotel.name === filters.name;
    //                 return matchesCity && matchesName;
    //             });
    //         })
    //     );
    // }
}
