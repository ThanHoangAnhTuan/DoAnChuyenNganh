import { Component, OnInit } from '@angular/core';
import { NavbarComponent } from '../../../components/navbar/navbar.component';
import SearchBoxComponent from '../../../components/search-box/search-box.component';
import { FormControl, FormsModule, ReactiveFormsModule } from '@angular/forms';
import { TuiIcon, tuiNumberFormatProvider, TuiSizeS } from '@taiga-ui/core';
import { TuiCheckbox } from '@taiga-ui/kit';
import { NgIf } from '@angular/common';
import {
    TuiInputRangeModule,
    TuiTextfieldControllerModule,
} from '@taiga-ui/legacy';
import { HotelService } from '../../../services/user/hotel.service';
import { ActivatedRoute, RouterLink } from '@angular/router';
import { AddressService } from '../../../services/address/address.service';
import { forkJoin, map, switchMap } from 'rxjs';

@Component({
    selector: 'app-search-page',
    imports: [
        NavbarComponent,
        SearchBoxComponent,
        FormsModule,
        TuiCheckbox,
        ReactiveFormsModule,
        TuiInputRangeModule,
        TuiTextfieldControllerModule,
        TuiIcon,
        NgIf,
        RouterLink,
    ],
    templateUrl: './search-page.component.html',
    standalone: true,
    styleUrl: './search-page.component.scss',
    providers: [
        tuiNumberFormatProvider({
            decimalSeparator: '.',
            thousandSeparator: ',',
            decimalMode: 'always',
        }),
    ],
})
export class SearchPageComponent implements OnInit {
    // Các FormControl dùng để kiểm tra validation
    protected readonly invalidTrue = new FormControl(true, () => ({
        invalid: true,
    }));
    protected readonly invalidFalse = new FormControl(false, () => ({
        invalid: true,
    }));

    // Giá trị min/max cho slider khoảng giá
    protected readonly max = 3_000_000;
    protected readonly min = 100_000;
    protected readonly control = new FormControl([this.min, this.max]);

    /**
     * Lấy kích thước cho các phần tử UI
     * @param first - Có phải phần tử đầu tiên không
     * @returns Kích thước 'm'
     */
    protected getSize(first: boolean): TuiSizeS {
        return first ? 'm' : 'm';
    }

    city: string = ''; // Thành phố tìm kiếm
    citySlug: string = '';
    checkIn: string = '';
    checkOut: string = '';
    hotels: any[] = []; // Danh sách khách sạn
    // level1AddressNames: string = '';
    // level2AddressNames: string = '';
    error = false; // Có lỗi khi tải dữ liệu không
    filteredHotels: any[] = [];

    constructor(
        private hotelService: HotelService,
        private addressService: AddressService,
        private route: ActivatedRoute
    ) {
        // Lấy tham số city từ URL
        this.route.params.subscribe((params) => {
            this.city = params['city'];
            if (this.hotels.length > 0) {
                this.applyFilters(); // Cập nhật khi params thay đổi
            }
        });

        // Lấy tham số city từ QueryParams
        this.route.queryParams.subscribe((params) => {
            this.citySlug = params['slug'];
            this.checkIn = params['checkIn'];
            this.checkOut = params['checkOut'];

            if (this.hotels.length > 0) {
                this.applyFilters(); // Cập nhật khi params thay đổi
            }
        });
    }

    // Khởi tạo component
    public ngOnInit(): void {
        this.invalidTrue.markAsTouched();
        this.invalidFalse.markAsTouched();
        this.loadHotels();
    }

    customCheckboxes = [
        {
            id: '4',
            label: '5 stars',
            checked: false,
        },
        {
            id: '5',
            label: '4 stars',
            checked: false,
        },
        {
            id: '6',
            label: '3 stars',
            checked: false,
        },
    ];

    page = 1;

    loadMoreHotels(): void {
        this.page++;
        this.loadHotels()
    }

    loadHotels(): void {
        this.hotelService
            .getAccommodationsByCity(this.citySlug, this.page)
            .pipe(
                switchMap((hotels) => {
                    const hotelList = hotels.data;
                    console.log('hotelList', hotels);
                    
                    const hotelWithCityName$ = hotelList.map((hotel) =>
                        this.addressService.getCityBySlug(hotel.city).pipe(
                            map((cityData) => {
                                const cityName = cityData[0]?.name || 'Unknow';

                                const district = cityData[0]?.level2s.find(
                                    (d) => d.slug === hotel.district
                                );
                                const districName = district?.name ?? 'Unknow';

                                return {
                                    ...hotel,
                                    city: cityName,
                                    district: districName,
                                };
                            })
                        )
                    );
                    return forkJoin(hotelWithCityName$);
                })
            )
            .subscribe({
                next: (hotelsWithCity) => {
                    this.hotels = hotelsWithCity;
                    // Lọc khách sạn theo thành phố từ URL
                    if (this.city && this.city.trim() !== '') {
                        this.filteredHotels = this.hotels.filter((hotel) =>
                            hotel.city
                                .toLowerCase()
                                .includes(this.city.toLowerCase())
                        );
                        this.filteredHotels = [...this.hotels];
                        this.applyFilters(); // Áp dụng bộ lọc ngay khi tải xong
                    }
                },
                error: (err) => {
                    this.error = true;
                },
            });
    }

    onCheckboxChange(item: any, checked: boolean): void {
        item.checked = checked;
        this.applyFilters();
    }

    applyFilters(): void {
        // Bước 1: Lọc theo city từ thanh search (URL params)
        let result = [...this.hotels];

        // Nếu có thành phố từ URL, lọc danh sách khách sạn
        if (this.city && this.city.trim() !== '') {
            result = result.filter((hotel) =>
                hotel.city.toLowerCase().includes(this.city.toLowerCase())
            );
        }

        // Bước 2: Tiếp tục lọc theo các filter checkbox
        const activeFilters = this.customCheckboxes.filter((cb) => cb.checked);

        // Nếu không có filter nào được chọn, trả về kết quả lọc theo URL
        if (activeFilters.length === 0) {
            this.filteredHotels = result;
            return;
        }

        const starFilters = activeFilters.filter((f) =>
            f.label.includes('stars')
        );
        const typeFilters = activeFilters.filter(
            (f) =>
                f.label === 'Guest houses' || f.label === 'Bed and breakfasts'
        );
        const ratingFilters = activeFilters.filter((f) =>
            f.label.includes('Very good:')
        );

        // Áp dụng filter lên kết quả đã lọc theo URL
        this.filteredHotels = result.filter((hotel) => {
            const passStarFilter =
                starFilters.length === 0 ||
                starFilters.some((filter) => {
                    const stars = parseInt(filter.label.split(' ')[0]);
                    return hotel.rating === stars;
                });

            const passTypeFilter =
                typeFilters.length === 0 ||
                typeFilters.some((filter) => hotel.type === filter.label);

            const passRatingFilter =
                ratingFilters.length === 0 ||
                ratingFilters.some((filter) => {
                    if (filter.label.includes('Very good: 8+')) {
                        return hotel.reviewScore >= 8;
                    }
                    return true;
                });

            return passStarFilter && passTypeFilter && passRatingFilter;
        });
    }

    getStars(rating: number): string {
        return '★'.repeat(rating) + ''.repeat(5 - rating);
    }
}
