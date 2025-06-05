import {
    OTPResponse,
    UpdateResponse,
} from './../../models/user/register.model';
import { HttpClient } from '@angular/common/http';
import {
    OTP,
    RegisterModel,
    RegisterResponse,
    UpdatePassword,
} from '../../models/user/register.model';
import { Observable } from 'rxjs';
import { LoginModel, LoginResponse } from '../../models/user/login.model';
import { Injectable } from '@angular/core';
@Injectable({
    providedIn: 'root',
})
export class UserService {
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
        return (
            !!localStorage.getItem('token') || !!sessionStorage.getItem('token')
        );
    }
}
