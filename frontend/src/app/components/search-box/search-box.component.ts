import {TuiComboBoxModule, TuiInputDateRangeModule} from '@taiga-ui/legacy';
import {TuiButton, TuiTextfieldOptionsDirective} from '@taiga-ui/core';
import {FormControl, FormsModule, ReactiveFormsModule, Validators} from '@angular/forms';
import {TuiDataListWrapper, TuiDataListWrapperComponent, TuiFilterByInputPipe} from '@taiga-ui/kit';
import {ActivatedRoute, Router} from '@angular/router';
import {ChangeDetectionStrategy, Component} from '@angular/core';
import {TuiDay} from '@taiga-ui/cdk';


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
    ],
    standalone: true,
    templateUrl: './search-box.component.html',
    styleUrl: './search-box.component.scss',
    changeDetection: ChangeDetectionStrategy.OnPush,
})

export default class SearchBoxComponent {
    protected city: string = '';
    //danh sách các thành phố có sẵn để người dùng chọn
    protected readonly cities = [
        'Ho Chi Minh',
        'Ha Noi',
        'Da Nang',
        'Nha Trang',
        'Hue',
        'Dong Nai',
        'Vung Tau',
        'Da Lat',
        'Can Tho',
        'Ninh Binh',
        'Bac Linh',
        'Binh Dinh',
        'Binh Thuan',
        'Cao Bang',
    ]
    protected readonly DayControl = new FormControl();
    protected searchCityControl = new FormControl('', Validators.required);

    protected readonly today = TuiDay.currentLocal(); // Lấy ngày hiện tại

    constructor(private activatedRoute: ActivatedRoute, private router: Router) {
        //Lấy tham số từ URL khi component khởi tạo
        this.activatedRoute.params.subscribe(params => {
            // this.searchCity = params['city'];
            this.city = params['city'];
        });
        //thiết lập validator cho date-range
        this.DayControl.setValidators([
            control => {
                const value = control.value;
                if (!value) return null;

                const today = TuiDay.currentLocal();
                const fromDay = TuiDay.fromLocalNativeDate(new Date(value.from));
                const toDay = TuiDay.fromLocalNativeDate(new Date(value.to));
                // Kiểm tra ngày không được trước ngày hiện tại
                if (fromDay.dayBefore(today) || toDay.dayBefore(today)) {
                    return {minDate: true};
                }
                return null;
            }
        ]);
    }

    /**
     * Kiểm tra xem input thành phố có hợp lệ không
     * @returns true nếu input không hợp lệ và đã được touch
     */
    get isSearchCityInvalid() {
        return this.searchCityControl.invalid && this.searchCityControl.touched;
    }

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
        if (this.DayControl.value) {
            // Định dạng ngày check-in và check-out
            const checkIn = `${this.DayControl.value?.from.formattedDayPart}-${this.DayControl.value?.from.formattedMonthPart}-${this.DayControl.value?.from.formattedYear}`;
            const checkOut = `${this.DayControl.value?.to.formattedDayPart}-${this.DayControl.value?.to.formattedMonthPart}-${this.DayControl.value?.to.formattedYear}`;
            //Chuyển hướng với thành phố và ngày người dùng đã nhập
            this.router.navigate(['/search', this.searchCityControl.value], {
                queryParams: {
                    checkIn,
                    checkOut
                }
            });
            return;
        }
        //chuyển hướng chỉ với thành phố
        this.router.navigate(['/search', this.searchCityControl.value]);
        return;
    }
}
