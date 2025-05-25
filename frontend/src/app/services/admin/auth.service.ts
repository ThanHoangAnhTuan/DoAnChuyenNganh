import { Injectable } from '@angular/core';
import { environment } from '../../../environments/environment';
import { HttpClient } from '@angular/common/http';
import {
    AdminLoginInput,
    AdminLoginOutput,
} from '../../models/admin/admin.model';
import { Observable } from 'rxjs';

@Injectable({
    providedIn: 'root',
})
export class AuthService {
    private apiUrl = `${environment.apiUrl}/admin`;
    constructor(private http: HttpClient) {}

    login(adminLogin: AdminLoginInput): Observable<AdminLoginOutput> {
        return this.http.post<AdminLoginOutput>(`${this.apiUrl}/login`, {
            account: adminLogin.account,
            password: adminLogin.password,
        });
    }
}
