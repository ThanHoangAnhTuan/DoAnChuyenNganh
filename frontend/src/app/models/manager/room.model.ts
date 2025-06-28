export interface Room {
    id: string;
    name: string;
    status: string;
}

export interface GetAccommodationRoomResponse {
    data: Room[];
    code: number;
    message: string;
}

export interface CreateRoom {
    prefix: string;
    quantity: number;
    accommodation_type_id: string;
}

export interface CreateRoomResponse {
    data: Room[];
    code: number;
    message: string;
}

export interface UpdateRoomResponse {
    data: Room;
    code: number;
    message: string;
}

export interface DeleteAccommodationRoomResponse {
    data: null;
    code: number;
    message: string;
}
