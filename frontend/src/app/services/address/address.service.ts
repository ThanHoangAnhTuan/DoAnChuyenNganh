import { Injectable } from '@angular/core';
import { environment } from '../../../environments/environment';
import { HttpClient } from '@angular/common/http';
import { map, Observable } from 'rxjs';
import { City, CityResponse } from '../../models/address/address.model';

@Injectable({
    providedIn: 'root',
})
export class AddressService {
    // private readonly localUrl = `tinh_tp.json`;
    private readonly localUrl = `data/tinh_tp.json`;

    constructor(private http: HttpClient) {}

    getCities(): Observable<CityResponse> {
        return this.http.get<CityResponse>(this.localUrl);
    }

    getCityByLevel1id(level1_id: string): Observable<City[]> {
        return this.getCities().pipe(
            map((cities) =>
                cities.data.filter((city) => city.level1_id === level1_id)
            )
        );
    }

    getCityBySlug(slug: string): Observable<City[]> {
        return this.getCities().pipe(
            map((cities) =>
                cities.data.filter((city) => city.slug === slug)
            )
        );
    }
}
