import { Injectable } from '@angular/core';
import { environment } from '../../../environments/environment';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import {
    CreateRoom,
    CreateRoomResponse,
    DeleteAccommodationRoomResponse,
    GetAccommodationRoomResponse,
    Room,
    UpdateRoomResponse,
} from '../../models/manager/room.model';

@Injectable({
    providedIn: 'root',
})
export class RoomService {
    private readonly apiUrl = `${environment.apiUrl}/accommodation-room`;

    constructor(private http: HttpClient) {}

    getAccommodationRooms(
        accommodationDetailID: string
    ): Observable<GetAccommodationRoomResponse> {
        return this.http.get<GetAccommodationRoomResponse>(
            `${this.apiUrl}/${accommodationDetailID}`
        );
    }

    createAccommodationRoom(room: CreateRoom): Observable<CreateRoomResponse> {
        const newRoom: CreateRoom = {
            prefix: room.prefix,
            quantity: room.quantity,
            accommodation_type_id: room.accommodation_type_id,
        };
        return this.http.post<CreateRoomResponse>(this.apiUrl, newRoom);
    }

    updateAccommodationRoom(room: Room): Observable<UpdateRoomResponse> {
        const newRoom: Room = {
            id: room.id,
            name: room.name,
            status: room.status,
        };
        return this.http.put<UpdateRoomResponse>(this.apiUrl, newRoom);
    }

    deleteAccommodationRoom(
        id: string
    ): Observable<DeleteAccommodationRoomResponse> {
        return this.http.delete<DeleteAccommodationRoomResponse>(
            `${this.apiUrl}/${id}`
        );
    }
}
