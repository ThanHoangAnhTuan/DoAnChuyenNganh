import {ChangeDetectionStrategy, Component, OnInit} from '@angular/core';
import {NavbarComponent} from '../../components/navbar/navbar.component';
import SearchBoxComponent from '../../components/search-box/search-box.component';
import {FormControl, FormsModule, ReactiveFormsModule} from '@angular/forms';
import {TuiIcon, tuiNumberFormatProvider, TuiSizeS} from '@taiga-ui/core';
import {TuiCheckbox} from '@taiga-ui/kit';
import {NgForOf, NgIf} from '@angular/common';
import {TuiInputRangeModule, TuiTextfieldControllerModule} from '@taiga-ui/legacy';
import {HotelService} from '../../services/hotel/hotel.service';
import {ActivatedRoute} from '@angular/router';

@Component({
    selector: 'app-search-page',
    imports: [
        NavbarComponent,
        SearchBoxComponent,
        FormsModule,
        TuiCheckbox,
        NgForOf,
        ReactiveFormsModule,
        TuiInputRangeModule,
        TuiTextfieldControllerModule,
        TuiIcon,
        NgIf,
    ],
    templateUrl: './search-page.component.html',
    styleUrl: './search-page.component.scss',
    changeDetection: ChangeDetectionStrategy.OnPush,
    providers: [
        tuiNumberFormatProvider({
            decimalSeparator: '.',
            thousandSeparator: ',',
            decimalMode: 'always',
        }),
    ],
})
export class SearchPageComponent implements OnInit {
    protected readonly invalidTrue = new FormControl(true, () => ({invalid: true}));
    protected readonly invalidFalse = new FormControl(false, () => ({invalid: true}));
    protected readonly max = 3_000_000;
    protected readonly min = 100_000;
    protected readonly control = new FormControl([this.min, this.max]);

    protected getSize(first: boolean): TuiSizeS {
        return first ? 'm' : 'm';
    }

    city: string = '';
    hotels: any[] = [];
    error = false;

    constructor(private hotelService: HotelService, private route: ActivatedRoute) {
        this.route.params.subscribe(params => {
            this.city = params['city'];
            console.log(this.city);
        });
    }

    public ngOnInit(): void {
        this.invalidTrue.markAsTouched();
        this.invalidFalse.markAsTouched();
        this.loadHotels();
    }

    customCheckboxes = [
        {id: '1', label: 'Guest houses', checked: false},
        {id: '2', label: 'Very good: 8+', checked: false},
        {id: '3', label: '5 stars', checked: false},
        {id: '4', label: 'Bed and breakfasts', checked: false},
        {id: '5', label: '4 stars', checked: false},
        {id: '6', label: '3 stars', checked: false},
    ];


    loadHotels(): void {
        this.hotelService.getHotels().subscribe({
            next: (hotels) => {
                this.hotels = hotels;
            },
            error: (err) => {
                console.error('Error loading hotels:', err);
            }
        });
    }

    formatPrice(price: number): string {
        return new Intl.NumberFormat('vi-VN', {
            style: 'currency',
            currency: 'VND',
            minimumFractionDigits: 0
        }).format(price).replace('₫', 'VND');
    }

    getStars(rating: number): string {
        return '★'.repeat(rating) + '☆'.repeat(5 - rating);
    }
}
