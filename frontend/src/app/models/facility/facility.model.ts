// facility
export interface Facility {
    id: string;
    name: string;
    image: string;
}

export interface CreateFacilityOutput {
    code: number;
    message: string;
    data: Facility[];
}

export interface DeleteFacilityResponse {
    data: null;
    code: number;
    message: string;
}

export interface GetFacilitiesOutput {
    code: number;
    message: string;
    data: Facility[];
}

export interface UpdateFacilityResponse {
    code: number;
    message: string;
    data: Facility;
}

// facility detail

export interface FacilityDetail {
    id: string;
    name: string;
}

export interface CreateFacilityDetailOutput {
    code: number;
    message: string;
    data: FacilityDetail[];
}

export interface DeleteFacilityDetailOutput {
    data: null;
    code: number;
    message: string;
}

export interface GetFacilitiesDetailOutput {
    code: number;
    message: string;
    data: FacilityDetail[];
}

export interface UpdateFacilityDetailOutput {
    code: number;
    message: string;
    data: FacilityDetail;
}

// export interface FacilityDetail {
//     id: string;
//     name: string;
// }

// export interface UpdateFacility {
//     id: string;
//     name: string;
//     image: string;
// }

// export interface CreateFacilityDetailInput {
//     name: string;
//     image: string;
// }

// export interface CreateFacilityDetailOutput {
//     code: number;
//     message: string;
//     data: FacilityDetail[];
// }

// export interface GetFacilitiesDetailOutput {
//     code: number;
//     message: string;
//     data: FacilityDetail[];
// }
