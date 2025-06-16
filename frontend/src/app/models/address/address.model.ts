export interface Ward {
    "level3_id": string,
    "name": string,
    "type": string,
}

export interface District {
    "level2_id": string,
    "name": string,
    "type": string,
    "level3s": Ward[],
}

export interface City {
    "id": string
    "level1_id": string,
    "name": string,
    "type": string,
    "level2s": District[],
}

export interface CityResponse {
    "city": City[]
}