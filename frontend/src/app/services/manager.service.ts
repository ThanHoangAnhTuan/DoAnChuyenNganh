import { Injectable } from '@angular/core';
import {
    ManagerLoginInput,
    ManagerLoginOutput,
} from '../models/accommodation.model';
import { Observable } from 'rxjs';
import { environment } from '../../environments/environment';
import { HttpClient } from '@angular/common/http';

@Injectable({
    providedIn: 'root',
})
export class ManagerService {
    private apiUrl = `${environment.apiUrl}/manager/`;

    constructor(private http: HttpClient) {}

    login(managerLogin: ManagerLoginInput): Observable<ManagerLoginOutput> {
        return this.http.post<ManagerLoginOutput>(this.apiUrl + 'login', {
            account: managerLogin.account,
            password: managerLogin.password,
        });
    }
}
