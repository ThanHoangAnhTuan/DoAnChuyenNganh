import { TuiComboBoxModule, TuiInputDateRangeModule } from '@taiga-ui/legacy';
import {
    TuiButton,
    tuiItemsHandlersProvider,
    TuiTextfield,
    TuiTextfieldOptionsDirective,
} from '@taiga-ui/core';
import {
    FormControl,
    FormsModule,
    ReactiveFormsModule,
    Validators,
} from '@angular/forms';
import {
    TuiChevron,
    TuiDataListWrapper,
    TuiDataListWrapperComponent,
    TuiFilterByInputPipe,
} from '@taiga-ui/kit';
import { ActivatedRoute, Router } from '@angular/router';
import { Component, EventEmitter, OnInit, Output, signal } from '@angular/core';
import { TuiDay, TuiDayRange } from '@taiga-ui/cdk';
import { AddressService } from '../../services/address/address.service';
import { City } from '../../models/address/address.model';

@Component({
    selector: 'app-search-box',
    imports: [
        TuiComboBoxModule,
        TuiTextfieldOptionsDirective,
        FormsModule,
        TuiDataListWrapperComponent,
        TuiDataListWrapper,
        TuiFilterByInputPipe,
        TuiInputDateRangeModule,
        ReactiveFormsModule,
        TuiButton,
        TuiTextfield,
    ],
    standalone: true,
    templateUrl: './search-box.component.html',
    styleUrl: './search-box.component.scss',
})
export default class SearchBoxComponent implements OnInit {
    protected city: string = '';
    //danh sách các thành phố có sẵn để người dùng chọn
    protected cities: City[] = [];
    protected cityNames: string[] = [];
    protected selectedCitySlug: string = '';
    protected selectedCityId: string = '';
    protected level2Adress: any;
    protected readonly DayControl = new FormControl();
    protected searchCityControl = new FormControl('', Validators.required);
    protected readonly today = TuiDay.currentLocal(); // Lấy ngày hiện tại
    protected selectedCityId: string | null = null;
    constructor(
        private activatedRoute: ActivatedRoute,
        private router: Router,
        private addressService: AddressService
    ) {
        //Lấy tham số từ URL khi component khởi tạo
        this.activatedRoute.params.subscribe((params) => {
            this.city = params['city'];
        });
        //thiết lập validator cho date-range
        this.DayControl.setValidators([
            (control) => {
                const value = control.value;
                if (!value) return null;
                const today = TuiDay.currentLocal();

                const newFromDate = new Date(
                    Number(value.from.formattedYear),
                    Number(value.from.formattedMonthPart) - 1,
                    Number(value.from.formattedDayPart)
                );
                console.log(newFromDate);
                const newToDate = new Date(
                    Number(value.to.formattedYear),
                    Number(value.to.formattedMonthPart) - 1,
                    Number(value.to.formattedDayPart)
                );

                const fromDay = TuiDay.fromLocalNativeDate(newFromDate);
                const toDay = TuiDay.fromLocalNativeDate(newToDate);
                // Kiểm tra ngày không được trước ngày hiện tại
                if (fromDay.dayBefore(today) || toDay.dayBefore(today)) {
                    return { minDate: true };
                }
                return null;
            },
        ]);
    }
    ngOnInit(): void {
        this.addressService.getCities().subscribe((data) => {
            console.log(data);
            this.cities = data.data;
            console.log(this.cities);

            this.cityNames = this.cities.map((city) => city.name);
            console.log("data:", this.cities);
        });

        // Lấy thành phố từ URL parameter
        this.activatedRoute.params.subscribe((params) => {
            const cityParam = params['city'];
            if (cityParam) {
                this.city = cityParam;
                this.searchCityControl.setValue(cityParam);
            }
        });

        this.searchCityControl.valueChanges.subscribe(
            (selectedCityName: string | null) => {
                const selectedCity = this.cities.find(
                    (city) => city.name === selectedCityName
                );
                if (selectedCity) {
                    this.selectedCityId = selectedCity.level1_id;
                    console.log('City id đã chọn:', this.selectedCityId);
                }
            }
        );

        // Lấy ngày từ query parameters
        this.activatedRoute.queryParams.subscribe((queryParams) => {
            if (queryParams['checkIn'] && queryParams['checkOut']) {
                try {
                    // Parse ngày từ định dạng "dd-MM-yyyy" (giống với cách bạn đang gửi lên URL)
                    const [checkInDay, checkInMonth, checkInYear] = queryParams[
                        'checkIn'
                    ]
                        .split('-')
                        .map(Number);
                    const [checkOutDay, checkOutMonth, checkOutYear] =
                        queryParams['checkOut'].split('-').map(Number);

                    const fromDay = new TuiDay(
                        checkInYear,
                        checkInMonth - 1, // Trừ 1 vì TuiDay dùng tháng 0-11
                        checkInDay
                    );

                    const toDay = new TuiDay(
                        checkOutYear,
                        checkOutMonth - 1, // Trừ 1 vì TuiDay dùng tháng 0-11
                        checkOutDay
                    );

                    // Tạo TuiDayRange và gán vào form control
                    this.DayControl.setValue(new TuiDayRange(fromDay, toDay));
                } catch (e) {
                    console.error('Lỗi khi parse ngày từ URL:', e);
                }
            }
        });
    }

    /**
     * Kiểm tra xem input thành phố có hợp lệ không
     * @returns true nếu input không hợp lệ và đã được touch
     */
    get isSearchCityInvalid() {
        return this.searchCityControl.invalid && this.searchCityControl.touched;
    }

    private getToday(): string {
        const today = new Date();
        const day = today.getDate().toString().padStart(2, '0');
        const month = (today.getMonth() + 1).toString().padStart(2, '0');
        const year = today.getFullYear();
        return `${day}-${month}-${year}`;
    }

    private getDateAfterDays(daysToAdd: number): string {
        const date = new Date();
        date.setDate(date.getDate() + daysToAdd);

        const day = date.getDate().toString().padStart(2, '0');
        const month = (date.getMonth() + 1).toString().padStart(2, '0');
        const year = date.getFullYear();

        return `${day}-${month}-${year}`;
    }

    getDateFromQueryParam() {}

    /**
     * Xử lý sự kiện tìm kiếm
     * - Kiểm tra validation
     * - Nếu có date range: chuyển hướng với query params checkIn và checkOut
     * - Nếu không có date range: chỉ chuyển hướng với thành phố
     */
    search() {
        if (this.searchCityControl.invalid) {
            this.searchCityControl.markAllAsTouched(); //Đánh dấu touched để hiển thị lỗi
            return;
        }
        this.searchChanged.emit(this.searchCityControl.value ?? undefined);
        const city_name = this.searchCityControl.value;
        const slug = this.selectedCitySlug;
        const checkIn = this.getToday();
        const checkOut = this.getDateAfterDays(7);

        if (this.DayControl.value) {
            // Định dạng ngày check-in và check-out
            const checkIn = `${this.DayControl.value?.from.formattedDayPart}-${this.DayControl.value?.from.formattedMonthPart}-${this.DayControl.value?.from.formattedYear}`;
            const checkOut = `${this.DayControl.value?.to.formattedDayPart}-${this.DayControl.value?.to.formattedMonthPart}-${this.DayControl.value?.to.formattedYear}`;
            //Chuyển hướng với thành phố và ngày người dùng đã nhập
            this.router.navigate(['/search', city_name], {
                queryParams: {
                    slug,
                    checkIn,
                    checkOut,
                },
            });
            return;
        }
        //chuyển hướng chỉ với thành phố
        this.router.navigate(['/search', city_name], {
            queryParams: {
                slug,
                checkIn,
                checkOut,
            },
        });
        return;
    }
    @Output() searchChanged = new EventEmitter<string>();
}