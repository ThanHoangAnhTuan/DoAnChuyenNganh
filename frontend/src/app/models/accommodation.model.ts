// get accommodations
export interface Facilities {
    wifi: boolean;
    airCondition: boolean;
    tv: boolean;
}

export interface PropertySurroundings {
    restaurant: boolean;
    bar: boolean;
}

export interface Accommodation {
    id: string;
    managerId: string;
    name: string;
    city: string;
    country: string;
    district: string;
    image: string;
    description: string;
    rating: string;
    facilities: Facilities;
    googleMap: string;
    propertySurrounds: PropertySurroundings;
    rules: string;
}

export interface GetAccommodationResponse {
    data: Accommodation[];
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
    googleMap: string;
    propertySurrounds: PropertySurroundings;
    rules: string;
    image: File[];
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
    googleMap: string;
    propertySurrounds: PropertySurroundings;
    rules: string;
    image: File[];
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