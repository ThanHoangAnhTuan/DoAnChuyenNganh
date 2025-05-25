import { Injectable } from '@angular/core';
import { environment } from '../../../environments/environment';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import {
    CreateAccommodationDetailResponse,
    CreateAccommodationDetails,
    DeleteAccommodationDetailResponse,
    GetAccommodationDetailsResponse,
    UpdateAccommodationDetailResponse,
    UpdateAccommodationDetails,
} from '../../models/manager/accommodation-detail.model';

@Injectable({
    providedIn: 'root',
})
export class AccommodationDetailService {
    private readonly apiUrl = `${environment.apiUrl}/accommodation-detail`;

    constructor(private http: HttpClient) {}

    getAccommodationDetails(
        accommodationId: string
    ): Observable<GetAccommodationDetailsResponse> {
        return this.http.get<GetAccommodationDetailsResponse>(
            `${this.apiUrl}/get-accommodation-details/${accommodationId}`
        );
    }

    createAccommodationDetail(
        accommodationDetail: CreateAccommodationDetails
    ): Observable<CreateAccommodationDetailResponse> {
        const newAccommodationDetail: CreateAccommodationDetails = {
            name: accommodationDetail.name,
            accommodation_id: accommodationDetail.accommodation_id,
            available_rooms: accommodationDetail.available_rooms,
            beds: {
                single_bed: accommodationDetail.beds.single_bed,
                double_bed: accommodationDetail.beds.double_bed,
                large_double_bed: accommodationDetail.beds.large_double_bed,
                extra_large_double_bed:
                    accommodationDetail.beds.extra_large_double_bed,
            },
            discount_id: accommodationDetail.discount_id,
            guests: accommodationDetail.guests,
            price: accommodationDetail.price,
        };

        console.log(newAccommodationDetail);

        return this.http.post<CreateAccommodationDetailResponse>(
            `${this.apiUrl}/create-accommodation-detail`,
            newAccommodationDetail
        );
    }

    updateAccommodationDetail(
        accommodationDetail: UpdateAccommodationDetails
    ): Observable<UpdateAccommodationDetailResponse> {
        const newAccommodationDetail: UpdateAccommodationDetails = {
            id: accommodationDetail.id,
            name: accommodationDetail.name,
            accommodation_id: accommodationDetail.accommodation_id,
            available_rooms: accommodationDetail.available_rooms,
            beds: {
                single_bed: accommodationDetail.beds.single_bed,
                double_bed: accommodationDetail.beds.double_bed,
                large_double_bed: accommodationDetail.beds.large_double_bed,
                extra_large_double_bed:
                    accommodationDetail.beds.extra_large_double_bed,
            },
            discount_id: accommodationDetail.discount_id,
            guests: accommodationDetail.guests,
            price: accommodationDetail.price,
        };

        return this.http.put<UpdateAccommodationDetailResponse>(
            `${this.apiUrl}/update-accommodation-detail`,
            newAccommodationDetail
        );
    }

    deleteAccommodationDetail(
        id: string
    ): Observable<DeleteAccommodationDetailResponse> {
        return this.http.delete<DeleteAccommodationDetailResponse>(
            `${this.apiUrl}/delete-accommodation-detail`,
            {
                body: {
                    id: id,
                },
            }
        );
    }
}
