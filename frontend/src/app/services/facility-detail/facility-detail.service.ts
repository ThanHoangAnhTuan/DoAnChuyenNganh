import { Injectable } from '@angular/core';
import { environment } from '../../../environments/environment';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import {
    CreateFacilityDetailOutput,
    DeleteFacilityDetailOutput,
    GetFacilitiesDetailOutput,
    UpdateFacilityDetailOutput,
} from '../../models/facility/facility.model';

@Injectable({
    providedIn: 'root',
})
export class FacilityDetailService {
    private readonly facilityUrl = `${environment.apiUrl}/facility-detail`;
    constructor(private http: HttpClient) {}

    getFacilityDetail(): Observable<GetFacilitiesDetailOutput> {
        return this.http.get<GetFacilitiesDetailOutput>(
            `${this.facilityUrl}/get-facility-detail`
        );
    }

    getFacilities(): Observable<GetFacilitiesDetailOutput> {
        return this.http.get<GetFacilitiesDetailOutput>(
            `${this.facilityUrl}/get-facility-detail`
        );
    }

    createFacility(data: any): Observable<CreateFacilityDetailOutput> {
        return this.http.post<CreateFacilityDetailOutput>(
            `${this.facilityUrl}/create-facility-detail`,
            data
        );
    }
    updateFacility(data: any): Observable<UpdateFacilityDetailOutput> {
        return this.http.put<UpdateFacilityDetailOutput>(
            `${this.facilityUrl}/update-facility-detail`,
            data
        );
    }

    deleteFacility(id: string): Observable<DeleteFacilityDetailOutput> {
        return this.http.delete<DeleteFacilityDetailOutput>(
            `${this.facilityUrl}/delete-facility-detail/${id}`
        );
    }
}
