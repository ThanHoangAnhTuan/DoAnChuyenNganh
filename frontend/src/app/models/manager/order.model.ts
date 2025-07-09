export interface GetOrdersByManagerResponse {
    code: number;
    message: string;
    data: Order[] | [];
}

export interface Order {
    id: string;
    accommodationId: string;
    accommodationName: string;
    checkIn: string;
    checkOut: string;
    finalTotal: string;
    orderStatus: string;
    orderDetail: OrderDetail[] | [];
    email: string;
    username: string;
    phone: string;
    createdAt: string;
    updatedAt: string;
}

export interface OrderDetail {
    accommodationDetailId: string;
    accommodationDetailName: string;
    accommodationDetailPrice: string;
    accommodationDetailGuests: number;
}
