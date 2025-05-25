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
    private apiUrl = `${environment.apiUrl}/accommodation/`;

    constructor(private http: HttpClient) {}

    getAccommodations(): Observable<GetAccommodationResponse> {
        return this.http.get<GetAccommodationResponse>(
            this.apiUrl + 'get-accommodations-by-manager'
        );
    }

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
        };
        return this.http.post<CreateAccommodationResponse>(
            this.apiUrl + 'create-accommodation',
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
        };
        return this.http.put<UpdateAccommodationResponse>(
            this.apiUrl + 'update-accommodation',
            newAccommodation
        );
    }

    deleteAccommodation(id: string): Observable<DeleteAccommodationResponse> {
        return this.http.delete<DeleteAccommodationResponse>(
            this.apiUrl + 'delete-accommodation/' + id
        );
    }
}
