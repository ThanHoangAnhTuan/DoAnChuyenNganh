export interface Ward {
    level3_id: string;
    name: string;
    type: string;
    slug: string;
}

export interface District {
    level2_id: string;
    name: string;
    type: string;
    slug: string;
    level3s: Ward[];
}

export interface City {
    level1_id: string;
    name: string;
    type: string;
    slug: string;
    level2s: District[];
}

export interface CityResponse {
    data: City[];
}
