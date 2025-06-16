import {
    LoginModel,
    LoginResponse,
    OTPResponse,
    UpdateResponse,
} from '../../models/user/auth.model';
import { HttpClient } from '@angular/common/http';
import {
    OTP,
    RegisterModel,
    RegisterResponse,
    UpdatePassword,
} from '../../models/user/auth.model';
import { Observable } from 'rxjs';
import { Injectable } from '@angular/core';
import { jwtDecode } from 'jwt-decode';
@Injectable({
    providedIn: 'root',
})
export class AuthService {
    private apiUrl = 'http://localhost:8080/api/v1/user';
    constructor(private http: HttpClient) {}

    registerUser(userData: RegisterModel): Observable<RegisterResponse> {
        return this.http.post<RegisterResponse>(
            this.apiUrl + '/register',
            userData
        );
    }

    verifyOTP(otpData: OTP): Observable<OTPResponse> {
        return this.http.post<OTPResponse>(
            this.apiUrl + '/verify-otp',
            otpData
        );
    }

    loginUser(userData: LoginModel): Observable<LoginResponse> {
        return this.http.post<LoginResponse>(this.apiUrl + '/login', userData);
    }

    updatePassword(update: UpdatePassword): Observable<UpdateResponse> {
        return this.http.post<UpdateResponse>(
            this.apiUrl + '/update-password-register',
            update
        );
    }
    // getUserProfile(): Observable<any> {
    //     return this.http.get<any>(`${this.apiUrl}/profile`);
    // }

    isLoggedIn(): boolean {
        return !!this.getToken();
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
}
