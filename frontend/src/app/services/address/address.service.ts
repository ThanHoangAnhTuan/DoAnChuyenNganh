import { Injectable } from '@angular/core';
import { environment } from '../../../environments/environment';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { City } from '../../models/address/address.model';

@Injectable({
  providedIn: 'root'
})
export class AddressService {
  // private readonly localUrl = `tinh_tp.json`;
  private readonly localUrl = `${environment.localUrl}/data`;

  constructor(private http: HttpClient) { }

  getCities(): Observable<City[]> {
    return this.http.get<City[]>(this.localUrl);
  }

  // http://localhost:3000/data?level1_id=79
  getCityByLevel1id(id: string): Observable<City[]> {
    return this.http.get<City[]>(this.localUrl + '?level1_id=' + id);
  }

  // http://localhost:3000/data?slug=
  getCityBySlug(slug: string): Observable<City[]> {
    return this.http.get<City[]>(this.localUrl + '?slug=' + slug);
  }
}
