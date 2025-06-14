import { Injectable } from '@angular/core';
import { environment } from '../../../environments/environment';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import {
    CreateFacilityInput,
    CreateFacilityOutput,
    DeleteFacilityResponse,
    GetFacilitiesDetailOutput,
    GetFacilitiesOutput,
    UpdateFacility,
    UpdateFacilityResponse,
} from '../../models/facility/facility.model';

@Injectable({
    providedIn: 'root',
})
export class FacilityService {
    private readonly facilityUrl = `${environment.apiUrl}/facility`;
    // private adminUrl = `${environment.apiUrl}/admin/facility`;
    constructor(private http: HttpClient) {}

    getFacilities(): Observable<GetFacilitiesOutput> {
        return this.http.get<GetFacilitiesOutput>(
            `${this.facilityUrl}/get-facilities`
        );
    }
    // createFacility(
    //     facilityData: CreateFacilityInput
    // ): Observable<CreateFacilityOutput> {
    //     const newFacility: CreateFacilityInput = {
    //         name: facilityData.name,
    //         image: facilityData.image,
    //     };
    //     return this.http.post<CreateFacilityOutput>(
    //         this.facilityUrl,
    //         newFacility
    //     );
    // }
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
    // updateFacility(
    //     facility: UpdateFacility
    // ): Observable<UpdateFacilityResponse> {
    //     const newFacility: UpdateFacility = {
    //         id: facility.id,
    //         name: facility.name,
    //         image: facility.image,
    //     };
    //     return this.http.put<UpdateFacilityResponse>(
    //         this.facilityUrl,
    //         newFacility
    //     );
    // }

    deleteFacility(id: string): Observable<DeleteFacilityResponse> {
        return this.http.delete<DeleteFacilityResponse>(this.facilityUrl, {
            body: { id },
        });
    }
    // deleteAccommodation(id: string): Observable<DeleteAccommodationResponse> {
    //         return this.http.delete<DeleteAccommodationResponse>(
    //             this.accommodationUrl,
    //             {
    //                 body: { id: id },
    //             }
    //         );
    //     }

    // getFacilityDetail(): Observable<GetFacilitiesDetailOutput> {
    //     return this.http.get<GetFacilitiesDetailOutput>(
    //         `${this.apiUrl}/get-facility-detail`
    //     );
    // }
}
