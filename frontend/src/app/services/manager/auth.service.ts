import { Injectable } from '@angular/core';
import { jwtDecode } from "jwt-decode";

@Injectable({
    providedIn: 'root',
})
export class AuthService {
    constructor() {}

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
