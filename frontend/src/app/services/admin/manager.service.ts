import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { environment } from '../../../environments/environment';
import { Observable } from 'rxjs';
import {
    CreateManager,
    CreateManagerOutput,
    GetAccommodationsOfManagerByAdminOutput,
    GetManagerOutput,
    VerifyAccommodationInput,
    VerifyAccommodationOutput,
} from '../../models/admin/manager.model';

@Injectable({
    providedIn: 'root',
})
export class ManagerService {
    private apiUrl = `${environment.apiUrl}/manager`;
    private adminUrl = `${environment.apiUrl}/admin`;

    constructor(private http: HttpClient) {}

    createNewManager(manager: CreateManager): Observable<CreateManagerOutput> {
        return this.http.post<CreateManagerOutput>(
            this.apiUrl + '/register',
            manager
        );
    }

    getManagers(): Observable<GetManagerOutput> {
        return this.http.get<GetManagerOutput>(`${this.adminUrl}/managers`);
    }

    getAccommodationsOfManagerByAdmin(id: string): Observable<GetAccommodationsOfManagerByAdminOutput> {
        return this.http.get<GetAccommodationsOfManagerByAdminOutput>(`${this.adminUrl}/manager/${id}/accommodations`);
    }

    updateVerified(newVerify: VerifyAccommodationInput): Observable<VerifyAccommodationOutput> {
        return this.http.put<VerifyAccommodationOutput>(this.adminUrl + '/verify-accommodation', newVerify);
    }
}
