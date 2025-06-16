export interface Beds {
    single_bed: number;
    double_bed: number;
    large_double_bed: number;
    extra_large_double_bed: number;
}

// export interface Facilities {
//     wifi: boolean;
//     air_condition: boolean;
//     tv: boolean;
// }

export interface AccommodationDetails {
    id: string;
    accommodation_id: string;
    name: string;
    guests: number;
    beds: Beds;
    // facilities: Facilities;
    available_rooms: number;
    price: number;
    discount_id: string;
    images: string[];
}

export interface GetAccommodationDetailsResponse {
    data: AccommodationDetails[];
    code: number;
    message: string;
}

export interface AccommodationSelect {
    id: string;
    name: string;
}

export interface DiscountSelect {
    id: string;
    name: string;
}

export interface CreateAccommodationDetails {
    accommodation_id: string;
    name: string;
    guests: number;
    beds: Beds;
    // facilities: Facilities;
    available_rooms: number;
    price: string;
    discount_id: string;
}

export interface CreateAccommodationDetailResponse {
    data: AccommodationDetails;
    code: number;
    message: string;
}

export interface UpdateAccommodationDetails {
    id: string;
    accommodation_id: string;
    name: string;
    guests: number;
    beds: Beds;
    // facilities: Facilities;
    available_rooms: number;
    price: string;
    discount_id: string;
}

export interface UpdateAccommodationDetailResponse {
    data: AccommodationDetails;
    code: number;
    message: string;
}

export interface DeleteAccommodationDetailResponse {
    data: null;
    code: number;
    message: string;
}
