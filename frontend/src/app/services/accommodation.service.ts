import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import {
    Accommodation,
    GetAccommodationResponse,
    CreateAccommodation,
    CreateAccommodationResponse,
    UpdateAccommodationResponse,
    UpdateAccommodation,
    DeleteAccommodationResponse,
} from '../models/accommodation.model';
import { environment } from '../../environments/environment';

@Injectable({
    providedIn: 'root',
})
export class AccommodationService {
    private apiUrl = `${environment.apiUrl}/accommodation/`;

    constructor(private http: HttpClient) {}

    getAccommodations(): Observable<GetAccommodationResponse> {
        return this.http.get<GetAccommodationResponse>(
            this.apiUrl + 'get-accommodations'
        );
    }

    createAccommodation(
        accommodation: CreateAccommodation
    ): Observable<CreateAccommodationResponse> {
        const formData = new FormData();
        console.log(accommodation.facilities);
        formData.append('name', accommodation.name);
        formData.append('country', accommodation.country);
        formData.append('city', accommodation.city);
        formData.append('district', accommodation.district);
        formData.append('description', accommodation.description);
        formData.append('facilities', JSON.stringify(accommodation.facilities));
        formData.append('google_map', accommodation.googleMap);
        formData.append(
            'propertySurrounds',
            JSON.stringify(accommodation.propertySurrounds)
        );
        formData.append('rules', accommodation.rules);
        accommodation.image.forEach((file) => {
            formData.append('image', file, file.name);
        });
        return this.http.post<CreateAccommodationResponse>(
            this.apiUrl + 'create-accommodation',
            formData
        );
    }

    updateAccommodation(
        accommodation: UpdateAccommodation
    ): Observable<UpdateAccommodationResponse> {
        const formData = new FormData();
        formData.append('id', accommodation.id.toString());
        formData.append('name', accommodation.name);
        formData.append('country', accommodation.country);
        formData.append('city', accommodation.city);
        formData.append('district', accommodation.district);
        formData.append('description', accommodation.description);
        formData.append('facilities', JSON.stringify(accommodation.facilities));
        formData.append('google_map', accommodation.googleMap);
        formData.append(
            'propertySurrounds',
            JSON.stringify(accommodation.propertySurrounds)
        );
        formData.append('rules', accommodation.rules);
        accommodation.image.forEach((file) => {
            formData.append('image', file, file.name);
        });
        return this.http.put<UpdateAccommodationResponse>(
            this.apiUrl + 'update-accommodation',
            formData
        );
    }

    deleteAccommodation(id: string): Observable<DeleteAccommodationResponse> {
        return this.http.delete<DeleteAccommodationResponse>(this.apiUrl + 'delete-accommodation/' + id);
    }
}
