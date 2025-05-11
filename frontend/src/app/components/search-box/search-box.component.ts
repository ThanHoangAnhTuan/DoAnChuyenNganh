import {TuiComboBoxModule, TuiInputDateRangeModule} from '@taiga-ui/legacy';
import {TuiButton, TuiTextfieldOptionsDirective} from '@taiga-ui/core';
import {FormControl, FormsModule, ReactiveFormsModule, Validators} from '@angular/forms';
import {TuiDataListWrapper, TuiDataListWrapperComponent, TuiFilterByInputPipe} from '@taiga-ui/kit';
import {ActivatedRoute, Router, RouterLink} from '@angular/router';
import {ChangeDetectionStrategy, Component} from '@angular/core';


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
        RouterLink,
    ],
    standalone: true,
    templateUrl: './search-box.component.html',
    styleUrl: './search-box.component.scss',
    changeDetection: ChangeDetectionStrategy.OnPush,
})

export default class SearchBoxComponent {
    protected city = null;
    protected readonly cities = [
        'Ho Chi Minh',
        'Ha Noi',
        'Da Nang',
        'Nha Trang',
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

    protected readonly control = new FormControl();
    protected readonly searchCity = new FormControl('', Validators.required);

    // protected readonly today = TuiDay.currentLocal();

    constructor(private activatedRoute: ActivatedRoute, private router: Router) {
        this.activatedRoute.params.subscribe(params => {
            this.city = params['city'];
            console.log(this.city);
        });

        // this.control.setValidators([
        //     control => {
        //         const value = control.value;
        //         if (!value) return null;
        //
        //         const today = TuiDay.currentLocal();
        //         const fromDay = TuiDay.fromLocalNativeDate(new Date(value.from));
        //         const toDay = TuiDay.fromLocalNativeDate(new Date(value.to));
        //
        //         if (fromDay.dayBefore(today) || toDay.dayBefore(today)) {
        //             return {minDate: true};
        //         }
        //         return null;
        //     }
        // ]);
    }

    get isSearchCityInvalid() {
        return this.searchCity.invalid && this.searchCity.touched;
    }

    search() {
        if (this.searchCity.invalid) {
            this.searchCity.markAllAsTouched();
            return;
        }

        const checkIn = `${this.control.value.from.formattedDayPart}-${this.control.value.from.formattedMonthPart}-${this.control.value.from.formattedYear}`;
        const checkOut = `${this.control.value.to.formattedDayPart}-${this.control.value.to.formattedMonthPart}-${this.control.value.to.formattedYear}`;

        this.router.navigate(['/search', this.city], {
            queryParams: {
                checkIn,
                checkOut
            }
        });
    }
}
