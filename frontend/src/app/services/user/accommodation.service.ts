import { Injectable } from '@angular/core';
import { environment } from '../../../environments/environment';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { GetAccommodationByIdResponse } from '../../models/manager/accommodation.model';

@Injectable({
    providedIn: 'root',
})
export class AccommodationService {
    private apiUrl = `${environment.apiUrl}/accommodations`;
    
    constructor(private http: HttpClient) { }

    getAccommodationDetailById(id: string): Observable<GetAccommodationByIdResponse> {
        return this.http.get<GetAccommodationByIdResponse>(this.apiUrl + '/' + id);
    }
}
