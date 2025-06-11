import { Facility } from '../facility/facility.model';
import { Pagination } from '../pagination/pagination.model';

export interface Accommodation {
    id: string;
    manager_id: string;
    name: string;
    city: string;
    country: string;
    district: string;
    address: string;
    description: string;
    rating: number;
    facilities: Facility[];
    images: string[];
    google_map: string;
}

export interface GetAccommodationResponse {
    data: Accommodation[];
    code: number;
    message: string;
    pagination: Pagination
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
    address: string;
    rating: number;
    description: string;
    google_map: string;
    facilities: string[];
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
    address: string;
    rating: number;
    description: string;
    google_map: string;
    facilities: string[];
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
    address: string;
    rating: string;
    google_map: string;
}
