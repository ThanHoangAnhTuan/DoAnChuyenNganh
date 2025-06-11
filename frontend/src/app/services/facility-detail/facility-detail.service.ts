import { Injectable } from '@angular/core';
import { environment } from '../../../environments/environment';
import { HttpClient } from '@angular/common/http';
import { GetFacilitiesDetailOutput } from '../../models/facility/facility.model';
import { Observable } from 'rxjs';

@Injectable({
    providedIn: 'root',
})
export class FacilityDetailService {
    private apiUrl = `${environment.apiUrl}/facility-detail`;
    constructor(private http: HttpClient) {}

    // getFacilities(): Observable<GetFacilitiesOutput> {
    //     return this.http.get<GetFacilitiesOutput>(
    //         `${this.apiUrl}/get-facilities`
    //     );
    // }
    getFacilityDetail(): Observable<GetFacilitiesDetailOutput> {
        return this.http.get<GetFacilitiesDetailOutput>(
            `${this.apiUrl}/get-facility-detail`
        );
    }
}
