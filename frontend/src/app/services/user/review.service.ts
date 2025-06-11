import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { CreateAccommodationResponse, CreateNewReview, GetReviewsByAccommodationIdResponse, Review } from '../../models/user/review.model';
import { environment } from '../../../environments/environment';

@Injectable({
  providedIn: 'root'
})
export class ReviewService {
  private baseUrl = `${environment.apiUrl}/review`;

  constructor(private http: HttpClient) { }

  getReviewsByAccommodationId(accommodationId: string): Observable<GetReviewsByAccommodationIdResponse> {
    return this.http.get<GetReviewsByAccommodationIdResponse>(this.baseUrl + '?accommodation_id=' + accommodationId);
  }

  addReview(review: CreateNewReview): Observable<CreateAccommodationResponse> {
    return this.http.post<CreateAccommodationResponse>(this.baseUrl + '/', review);
  }
}
