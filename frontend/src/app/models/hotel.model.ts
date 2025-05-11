export interface Hotel {
    id: number;
    featured: boolean;
    deal?: boolean;
    image: string;
    name: string;
    rating: number;
    location: {
        district: string;
        city: string;
        distance: string;
    };
    roomType: string;
    roomDetails?: string;
    bedInfo: string;
    amenities: string[];
    price: {
        original: number;
        discounted: number;
        taxes: number;
        currency: string;
    };
    reviews: {
        score: number;
        rating: string;
        count: number;
        new?: boolean;
    };
    availability: {
        roomsLeft?: number;
        stayInfo: string;
    };
}
export interface HotelResponse {
    hotels: Hotel[];
}
