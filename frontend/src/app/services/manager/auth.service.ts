import { Injectable } from '@angular/core';
import {
    ManagerLoginInput,
    ManagerLoginOutput,
} from '../../models/manager/accommodation.model';
import { Observable } from 'rxjs';
import { environment } from '../../../environments/environment';
import { HttpClient } from '@angular/common/http';
import { jwtDecode } from 'jwt-decode';

@Injectable({
    providedIn: 'root',
})
export class AuthService {
    private apiUrl = `${environment.apiUrl}/manager`;

    constructor(private http: HttpClient) {}

    login(managerLogin: ManagerLoginInput): Observable<ManagerLoginOutput> {
        return this.http.post<ManagerLoginOutput>(`${this.apiUrl}/login`, {
            account: managerLogin.account,
            password: managerLogin.password,
        });
    }

    private getToken(): string | null {
        const cookieString = document.cookie;
        const cookies = cookieString.split(';');

        for (let i = 0; i < cookies.length; i++) {
            const cookie = cookies[i].trim();
            if (cookie.startsWith('auth_token=')) {
                return cookie.substring('auth_token='.length, cookie.length);
            }
        }

        return null;
    }

    getUserRole(): string | null {
        const token = this.getToken();
        if (!token) return null;
        try {
            const decoded: any = jwtDecode(token);
            return decoded.role || null;
        } catch (e) {
            return null;
        }
    }

    isLoggedIn(): boolean {
        return !!this.getToken();
    }
}
