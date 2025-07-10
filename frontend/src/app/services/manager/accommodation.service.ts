import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import {
    GetAccommodationResponse,
    CreateAccommodation,
    CreateAccommodationResponse,
    UpdateAccommodationResponse,
    UpdateAccommodation,
    DeleteAccommodationResponse,
} from '../../models/manager/accommodation.model';
import { environment } from '../../../environments/environment';

@Injectable({
    providedIn: 'root',
})
export class AccommodationService {
    private readonly accommodationUrl = `${environment.apiUrl}/accommodations`;
    private readonly managerUrl = `${environment.apiUrl}/manager/accommodations`;

    constructor(private http: HttpClient) {}

    getAccommodations(): Observable<GetAccommodationResponse> {
        return this.http.get<GetAccommodationResponse>(this.managerUrl);
    }

    getAccommodationWithPage(page: number): Observable<GetAccommodationResponse> {
        return this.http.get<GetAccommodationResponse>(this.managerUrl + '?page=' + page)
    };

    createAccommodation(
        accommodation: CreateAccommodation
    ): Observable<CreateAccommodationResponse> {
        const newAccommodation: CreateAccommodation = {
            name: accommodation.name,
            city: accommodation.city,
            country: accommodation.country,
            description: accommodation.description,
            district: accommodation.district,
            google_map: accommodation.google_map,
            address: accommodation.address,
            rating: accommodation.rating,
            facilities: accommodation.facilities,
        };
        return this.http.post<CreateAccommodationResponse>(
            this.accommodationUrl,
            newAccommodation
        );
    }

    updateAccommodation(
        accommodation: UpdateAccommodation
    ): Observable<UpdateAccommodationResponse> {
        const newAccommodation: UpdateAccommodation = {
            id: accommodation.id,
            name: accommodation.name,
            city: accommodation.city,
            country: accommodation.country,
            description: accommodation.description,
            district: accommodation.district,
            address: accommodation.address,
            google_map: accommodation.google_map,
            rating: accommodation.rating,
            facilities: accommodation.facilities,
        };
        return this.http.put<UpdateAccommodationResponse>(
            this.accommodationUrl,
            newAccommodation
        );
    }

    deleteAccommodation(id: string): Observable<DeleteAccommodationResponse> {
        return this.http.delete<DeleteAccommodationResponse>(
            this.accommodationUrl,
            {
                body: { id: id },
            }
        );
    }
}
