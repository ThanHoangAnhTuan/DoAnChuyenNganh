import { Injectable } from '@angular/core';
import { environment } from '../../../environments/environment';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import {
    GetFacilitiesDetailOutput,
    GetFacilitiesOutput,
} from '../../models/facility/facility.model';

@Injectable({
    providedIn: 'root',
})
export class FacilityService {
    private apiUrl = `${environment.apiUrl}/facility`;
    constructor(private http: HttpClient) {}

    getFacilities(): Observable<GetFacilitiesOutput> {
        return this.http.get<GetFacilitiesOutput>(
            `${this.apiUrl}/get-facilities`
        );
    }
    // getFacilityDetail(): Observable<GetFacilitiesDetailOutput> {
    //     return this.http.get<GetFacilitiesDetailOutput>(
    //         `${this.apiUrl}/get-facility-detail`
    //     );
    // }
}
