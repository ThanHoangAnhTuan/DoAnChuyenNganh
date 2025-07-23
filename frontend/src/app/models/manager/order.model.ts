export interface GetOrdersByManagerResponse {
    code: number;
    message: string;
    data: Order[] | [];
}

export interface GetOrdersByUserResponse {
    code: number;
    message: string;
    data: Order[] | [];
}

export interface GetOrderChageStatusResponse {
    code: number;
    message: string;
}

export interface Order {
    id: string;
    accommodation_id: string;
    accommodation_name: string;
    check_in: string;
    check_out: string;
    final_total: string;
    order_status: string;
    order_detail: OrderDetail[] | [];
    email: string;
    username: string;
    phone: string;
    created_at: string;
    updated_at: string;
}

export interface OrderDetail {
    accommodation_detail_id: string;
    accommodation_detail_name: string;
    price: string;
    guests: number;
    room_bookings: RoomBooking[];
}

export interface RoomBooking {
    id: string;
    accommodation_room_id: string;
    room_name: string;
    booking_status: string;
}