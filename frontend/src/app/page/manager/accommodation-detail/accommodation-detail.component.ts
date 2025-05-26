import { Component, inject, OnInit } from '@angular/core';
import { Accommodation } from '../../../models/manager/accommodation.model';
import {
    TuiAppearance,
    TuiButton,
    TuiDataList,
    TuiDialogService,
    TuiGroup,
    TuiIcon,
    TuiTextfield,
} from '@taiga-ui/core';
import {
    FormControl,
    FormGroup,
    FormsModule,
    ReactiveFormsModule,
    Validators,
} from '@angular/forms';
import { AccommodationService } from '../../../services/manager/accommodation.service';
import type { PolymorpheusContent } from '@taiga-ui/polymorpheus';
import { ActivatedRoute, RouterLink } from '@angular/router';
import { AccommodationDetailService } from '../../../services/manager/accommodation-detail.service';
import {
    AccommodationDetails,
    AccommodationSelect,
    CreateAccommodationDetails,
    DiscountSelect,
    UpdateAccommodationDetails,
} from '../../../models/manager/accommodation-detail.model';
import { TuiTable } from '@taiga-ui/addon-table';
import { TuiInputModule } from '@taiga-ui/legacy';
import { TuiCardLarge } from '@taiga-ui/layout';
import {
    TuiCheckbox,
    TuiChevron,
    TuiInputNumber,
    TuiSelect,
} from '@taiga-ui/kit';
import { TuiContext } from '@taiga-ui/cdk';

@Component({
    selector: 'app-accommodation-detail',
    imports: [
        TuiTable,
        TuiIcon,
        TuiButton,
        TuiInputModule,
        FormsModule,
        ReactiveFormsModule,
        TuiTextfield,
        TuiAppearance,
        TuiCardLarge,
        TuiGroup,
        TuiCheckbox,
        RouterLink,
        TuiInputNumber,
        TuiDataList,
        TuiSelect,
        TuiChevron,
    ],
    templateUrl: './accommodation-detail.component.html',
    styleUrl: './accommodation-detail.component.scss',
})
export class AccommodationDetailComponent implements OnInit {
    protected accommodationDetails!: AccommodationDetails[];
    protected readonly columns: string[] = [
        'ID',
        'Name',
        'Guests',
        'Single Bed',
        'Double Bed',
        'Large Double Bed',
        'Extra Large Double Bed',
        'Available Rooms',
        'Price',
        'Image',
        'Accommodation',
        'Discount',
        'Action',
    ];
    protected readonly baseUrl: string = 'http://localhost:8080/uploads/';
    protected idAccommodationDetailUpdating = '';

    private readonly dialogs = inject(TuiDialogService);

    protected formAccommodationDetail = new FormGroup({
        name: new FormControl<string | ''>('', Validators.required),
        guests: new FormControl<number | 0>(0, Validators.min(1)),
        singleBed: new FormControl<number | 0>(0),
        doubleBed: new FormControl<number | 0>(0),
        largeDoubleBed: new FormControl<number | 0>(0),
        extraLargeDoubleBed: new FormControl<number | 0>(0),
        price: new FormControl<number | 0>(0, Validators.min(1)),
        availableRooms: new FormControl<number | 0>(0),
        accommodationId: new FormControl<string | ''>('', Validators.required),
        discountId: new FormControl<string | ''>(''),
    });

    protected readonly resetFormAccommodationDetail = {
        accommodationId: '',
        availableRooms: 0,
        discountId: '',
        doubleBed: 0,
        extraLargeDoubleBed: 0,
        guests: 0,
        largeDoubleBed: 0,
        name: '',
        price: 0,
        singleBed: 0,
    };

    protected accommodations!: Accommodation[];

    protected accommodationItems: readonly AccommodationSelect[] = [];

    constructor(
        private route: ActivatedRoute,
        private accommodationDetailService: AccommodationDetailService,
        private accommodationService: AccommodationService
    ) {}

    ngOnInit() {
        this.route.params.subscribe((params) => {
            this.accommodationDetailService
                .getAccommodationDetails(params['id'])
                .subscribe((response) => {
                    this.accommodationDetails = response.data;
                });
        });

        this.accommodationService.getAccommodations().subscribe((response) => {
            this.accommodationItems = response.data.map((item) => ({
                id: item.id,
                name: item.name,
            }));
        });
    }

    protected openDialogCreate(content: PolymorpheusContent): void {
        this.formAccommodationDetail.reset(this.resetFormAccommodationDetail);

        this.dialogs
            .open(content, {
                label: 'Create Accommodation Detail',
            })
            .subscribe({
                complete: () => {
                    this.formAccommodationDetail.reset(
                        this.resetFormAccommodationDetail
                    );
                },
            });
    }

    protected openDialogUpdate(
        content: PolymorpheusContent,
        accommodationDetail: AccommodationDetails
    ) {
        this.formAccommodationDetail.reset(this.resetFormAccommodationDetail);

        this.formAccommodationDetail.patchValue({
            name: accommodationDetail.name,
            accommodationId: accommodationDetail.accommodation_id,
            availableRooms: accommodationDetail.available_rooms,
            discountId: accommodationDetail.discount_id,

            doubleBed: accommodationDetail.beds.double_bed,
            singleBed: accommodationDetail.beds.single_bed,
            largeDoubleBed: accommodationDetail.beds.large_double_bed,
            extraLargeDoubleBed:
                accommodationDetail.beds.extra_large_double_bed,
            guests: accommodationDetail.guests,
            price: accommodationDetail.price,
        });

        this.idAccommodationDetailUpdating = accommodationDetail.id;

        this.dialogs
            .open(content, {
                label: 'Update Accommodation',
            })
            .subscribe({
                complete: () => {
                    this.formAccommodationDetail.reset(
                        this.resetFormAccommodationDetail
                    );
                },
            });
    }

    protected createAccommodationDetail() {
        const accommodationDetail: CreateAccommodationDetails = {
            name: this.formAccommodationDetail.get('name')?.value || '',
            guests: this.formAccommodationDetail.get('guests')?.value || 0,
            beds: {
                single_bed:
                    this.formAccommodationDetail.get('singleBed')?.value || 0,
                double_bed:
                    this.formAccommodationDetail.get('doubleBed')?.value || 0,
                large_double_bed:
                    this.formAccommodationDetail.get('largeDoubleBed')?.value ||
                    0,
                extra_large_double_bed:
                    this.formAccommodationDetail.get('extraLargeDoubleBed')
                        ?.value || 0,
            },
            available_rooms:
                this.formAccommodationDetail.get('availableRooms')?.value || 0,
            price: this.formAccommodationDetail.get('price')?.value || 0,
            accommodation_id:
                this.formAccommodationDetail.get('accommodationId')?.value ||
                '',
            discount_id:
                this.formAccommodationDetail.get('discountId')?.value || '',
        };

        if (this.formAccommodationDetail.invalid) {
            this.formAccommodationDetail.markAllAsTouched();
            return;
        }

        this.accommodationDetailService
            .createAccommodationDetail(accommodationDetail)
            .subscribe((response) => {
                this.accommodationDetails.push(response.data);
            });
    }

    protected updateAccommodationDetail() {
        const accommodationDetail: UpdateAccommodationDetails = {
            id: this.idAccommodationDetailUpdating,
            accommodation_id:
                this.formAccommodationDetail.get('accommodationId')?.value ||
                '',
            name: this.formAccommodationDetail.get('name')?.value || '',
            available_rooms:
                this.formAccommodationDetail.get('availableRooms')?.value || 0,
            beds: {
                single_bed:
                    this.formAccommodationDetail.get('singleBed')?.value || 0,
                double_bed:
                    this.formAccommodationDetail.get('doubleBed')?.value || 0,
                large_double_bed:
                    this.formAccommodationDetail.get('largeDoubleBed')?.value ||
                    0,
                extra_large_double_bed:
                    this.formAccommodationDetail.get('extraLargeDoubleBed')
                        ?.value || 0,
            },
            discount_id:
                this.formAccommodationDetail.get('discountId')?.value || '',
            guests: this.formAccommodationDetail.get('guests')?.value || 0,
            price: this.formAccommodationDetail.get('price')?.value || 0,
        };

        this.accommodationDetailService
            .updateAccommodationDetail(accommodationDetail)
            .subscribe((response) => {
                this.accommodationDetails = this.accommodationDetails.map(
                    (accommodationDetail) => {
                        if (accommodationDetail.id === response.data.id) {
                            return response.data;
                        } else {
                            return accommodationDetail;
                        }
                    }
                );
            });
    }

    protected deleteAccommodationDetail(id: string) {
        this.accommodationDetailService
            .deleteAccommodationDetail(id)
            .subscribe((response) => {
                this.accommodationDetails = this.accommodationDetails.filter(
                    (accommodationDetail) => accommodationDetail.id !== id
                );
            });
    }

    protected readonly contentAccommodation: PolymorpheusContent<
        TuiContext<string | null>
    > = ({ $implicit: id }) =>
        this.accommodationItems.find((item) => item.id === id)?.name ?? '';

    protected readonly discountItems: readonly DiscountSelect[] = [];

    protected readonly contentDiscount: PolymorpheusContent<
        TuiContext<string | null>
    > = ({ $implicit: id }) =>
        this.discountItems.find((item) => item.id === id)?.name ?? '';
}
