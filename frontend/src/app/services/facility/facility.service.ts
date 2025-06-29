import { Injectable } from '@angular/core';
import { environment } from '../../../environments/environment';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import {
    CreateFacilityOutput,
    DeleteFacilityResponse,
    GetFacilitiesOutput,
    UpdateFacilityResponse,
} from '../../models/facility/facility.model';

@Injectable({
    providedIn: 'root',
})
export class FacilityService {
    private readonly facilityUrl = `${environment.apiUrl}/facility`;
    constructor(private http: HttpClient) {}

    getFacilities(): Observable<GetFacilitiesOutput> {
        return this.http.get<GetFacilitiesOutput>(
            `${this.facilityUrl}/get-facilities`
        );
    }

    createFacility(formData: FormData): Observable<CreateFacilityOutput> {
        return this.http.post<CreateFacilityOutput>(
            `${this.facilityUrl}/create-facility`,
            formData
        );
    }

    updateFacility(formData: FormData): Observable<UpdateFacilityResponse> {
        return this.http.put<UpdateFacilityResponse>(
            `${this.facilityUrl}/update-facility`,
            formData
        );
    }

    deleteFacility(id: string): Observable<DeleteFacilityResponse> {
        return this.http.delete<DeleteFacilityResponse>(
            `${this.facilityUrl}/delete-facility/${id}`
        );
    }
}
