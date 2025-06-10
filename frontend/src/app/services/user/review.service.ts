import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { CreateNewReview, Review } from '../../models/user/review.model';

@Injectable({
  providedIn: 'root'
})
export class ReviewService {
  private baseUrl = 'http://localhost:8080/api/v1/review';

  constructor(private http: HttpClient) { }

  getReviewsByAccommodationId(accommodationId: string): Observable<Review[]> {
    return this.http.get<Review[]>(`${this.baseUrl}?accommodation_id=${accommodationId}`);
  }

  addReview(review: Review): Observable<CreateNewReview> {
    return this.http.post<CreateNewReview>(this.baseUrl + '/', review);
  }
}
