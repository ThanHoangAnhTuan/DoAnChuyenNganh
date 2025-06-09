export interface RoomSelected {
    id: string;
    quantity: number;
}

export interface Payment {
    check_in: string;
    check_out: string;
    accommodation_id: string;
    room_selected: RoomSelected[];
}

export interface CreatePayment {
    check_in: string;
    check_out: string;
    accommodation_id: string;
    room_selected: RoomSelected[];
}