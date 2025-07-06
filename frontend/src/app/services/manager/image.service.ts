import { HttpClient, HttpParams } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { environment } from '../../../environments/environment';
import {
    GetImagesResponse,
    UploadImagesResponse,
} from '../../models/manager/image.model';

@Injectable({
    providedIn: 'root',
})
export class ImageService {
    private apiUrl = `${environment.apiUrl}/image`;

    constructor(private http: HttpClient) {}

    getImages(id: string, isDetail: boolean): Observable<GetImagesResponse> {
        const params = new HttpParams().set('is_detail', isDetail.toString());

        return this.http.get<GetImagesResponse>(
            `${this.apiUrl}/get-images/${id}`,
            {
                params,
            }
        );
    }

    uploadImages(
        deleteImages: string[],
        formImages: File[],
        id: string,
        isDetail: boolean
    ): Observable<UploadImagesResponse> {
        const formData = new FormData();

        formImages.forEach((file) => {
            formData.append('images', file);
        });

        deleteImages.forEach((image) => {
            formData.append('delete_images', image);
        });
        formData.append('id', id);
        formData.append('is_detail', isDetail.toString());

        return this.http.post<UploadImagesResponse>(
            `${this.apiUrl}/upload-images`,
            formData
        );
    }
}
