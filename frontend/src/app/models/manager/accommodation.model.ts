// get accommodations
export interface Facilities {
    wifi: boolean;
    air_condition: boolean;
    tv: boolean;
}

export interface PropertySurroundings {
    restaurant: boolean;
    bar: boolean;
}

export interface Accommodation {
    id: string;
    manager_id: string;
    name: string;
    city: string;
    country: string;
    district: string;
    description: string;
    rating: string;
    facilities: Facilities;
    images: string[];
    google_map: string;
    property_surrounds: PropertySurroundings;
    rules: string;
}

export interface GetAccommodationResponse {
    data: Accommodation[];
    code: number;
    message: string;
}

export interface GetAccommodationByIdResponse {
    data: Accommodation;
    code: number;
    message: string;
}

// create accommodation
export interface CreateAccommodation {
    name: string;
    country: string;
    city: string;
    district: string;
    description: string;
    facilities: Facilities;
    google_map: string;
    property_surrounds: PropertySurroundings;
    rules: string;
}

export interface CreateAccommodationResponse {
    data: Accommodation;
    code: number;
    message: string;
}

// update accommodation
export interface UpdateAccommodation {
    id: string;
    name: string;
    country: string;
    city: string;
    district: string;
    description: string;
    facilities: Facilities;
    google_map: string;
    property_surrounds: PropertySurroundings;
    rules: string;
}
export interface UpdateAccommodationResponse {
    data: Accommodation;
    code: number;
    message: string;
}

export interface DeleteAccommodationResponse {
    data: null;
    code: number;
    message: string;
}

// manager
export interface ManagerLoginInput {
    account: string;
    password: string;
}

export interface ManagerLoginOutput {
    code: number;
    message: string;
    data: {
        token: string;
        account: string;
        userName: string;
    };
}

export interface AccommodationByCityResponse {
    code: number;
    message: string;
    data: AccommodationByCity[];
}

// search accommodation by city
export interface AccommodationByCity {
    id: string;
    name: string;
    city: string;
    country: string;
    district: string;
    rating: string;
    google_map: string;
}
