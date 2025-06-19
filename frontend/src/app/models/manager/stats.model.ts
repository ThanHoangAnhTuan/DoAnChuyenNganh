export interface GetMonthlyEarnings {
    month: number;
    total_orders: number;
    total_revenue: string;
}

export interface GetMonthlyEarningsResponse {
    code: number;
    message: string;
    data: GetMonthlyEarnings[];
}

export interface GetDailyEarnings {
    day: string;
    total_orders: number;
    total_revenue: string;
}

export interface GetDailyEarningsResponse {
    code: number;
    message: string;
    data: GetDailyEarnings[];
}
