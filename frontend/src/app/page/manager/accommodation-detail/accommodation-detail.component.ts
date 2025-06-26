import { FacilityDetailService } from './../../../services/facility-detail/facility-detail.service';
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
import { FacilityDetail } from '../../../models/facility/facility.model';
import { NavbarComponent } from '../../../components/navbar/navbar.component';
import { MessageService } from 'primeng/api';
import { Toast } from 'primeng/toast';
import { ButtonModule } from 'primeng/button';
import { Ripple } from 'primeng/ripple';

@Component({
    selector: 'app-accommodation-detail',
    imports: [
        TuiTable,
        FormsModule,
        ReactiveFormsModule,
        // TuiIcon,
        TuiButton,
        TuiInputModule,
        FormsModule,
        ReactiveFormsModule,
        TuiTextfield,
        TuiAppearance,
        TuiCardLarge,
        // TuiGroup,
        TuiCheckbox,
        RouterLink,
        TuiInputNumber,
        TuiDataList,
        TuiSelect,
        TuiChevron,
        NavbarComponent,
        Toast,
        ButtonModule,
        Ripple,
    ],
    templateUrl: './accommodation-detail.component.html',
    styleUrl: './accommodation-detail.component.scss',
    providers: [MessageService],
})
export class AccommodationDetailComponent implements OnInit {
    protected accommodationDetails!: AccommodationDetails[];
    protected facilities!: FacilityDetail[];
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
        'Action',
        // 'Discount',
        // 'FacilityDetails',
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
        accommodationId: new FormControl<string | ''>(''),
        discountId: new FormControl<string | ''>(''),
        facilityDetails: new FormControl<string | ''>(''),
    });
    protected formFacilityDetail = new FormGroup({});
    protected accommodationId: string = '';

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
        facilityDetails: '',
    };

    protected accommodations!: Accommodation[];

    protected accommodationItems: readonly AccommodationSelect[] = [];
    constructor(
        private route: ActivatedRoute,
        private accommodationDetailService: AccommodationDetailService,
        private accommodationService: AccommodationService,
        private facilityDetailService: FacilityDetailService,
        private messageService: MessageService
    ) {}
    showToast(
        severity: 'success' | 'info' | 'warn' | 'error',
        summary: string,
        detail: string
    ): void {
        this.messageService.add({ severity, summary, detail });
    }

    ngOnInit() {
        this.route.params.subscribe((params) => {
            this.accommodationId = params['id'];
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
        this.facilityDetailService.getFacilityDetail().subscribe((response) => {
            this.facilities = response.data;
            this.createFacilityControls();
        });
    }
    private createFacilityControls() {
        const facilityControls: { [key: string]: FormControl<boolean> } = {};

        if (!this.facilities || this.facilities.length === 0) {
            this.formFacilityDetail = new FormGroup(facilityControls);
            return;
        }

        this.facilities.forEach((facility) => {
            facilityControls[facility.id] = new FormControl<boolean>(false, {
                nonNullable: true,
            });
        });

        // Tạo FormGroup mới với các controls
        this.formFacilityDetail = new FormGroup(facilityControls);
    }
    getSelectedFacilityIds(): string[] {
        if (!this.facilities || this.facilities.length === 0) {
            return [];
        }
        return this.facilities
            .filter(
                (facility) =>
                    this.formFacilityDetail.get(facility.id)?.value === true
            )
            .map((facility) => facility.id);
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

        this.setFacilityDetailValues(accommodationDetail.facilities);
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
    private setFacilityDetailValues(
        accommodationFacilityDetail: FacilityDetail[]
    ) {
        const facilityValues: { [key: string]: boolean } = {};
        Object.keys(this.formFacilityDetail.controls).forEach((facilityId) => {
            facilityValues[facilityId] = false;
        });

        accommodationFacilityDetail.forEach((facilityId) => {
            if (facilityValues.hasOwnProperty(facilityId.id)) {
                facilityValues[facilityId.id] = true;
            }
        });

        this.formFacilityDetail.patchValue(facilityValues);
    }

    protected createAccommodationDetail() {
        // this.formAccommodationDetail.patchValue({
        //     accommodationId: this.accommodationId,
        // });

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
            price: `${this.formAccommodationDetail.get('price')?.value || 0}`,
            accommodation_id: this.accommodationId,
            discount_id:
                this.formAccommodationDetail.get('discountId')?.value || '',
            facilities: this.getSelectedFacilityIds(),
        };

        console.log(accommodationDetail);

        if (this.formAccommodationDetail.invalid) {
            this.showToast(
                'error',
                'Tạo chi tiết chỗ ở thất bại',
                'Vui lòng điền đầy đủ thông tin'
            );
            this.formAccommodationDetail.markAllAsTouched();
            console.log('here');
            return;
        }

        this.accommodationDetailService
            .createAccommodationDetail(accommodationDetail)
            .subscribe((response) => {
                this.accommodationDetails.push(response.data);
                this.formAccommodationDetail.reset();
                this.formFacilityDetail.reset();
                this.showToast(
                    'success',
                    'Tạo chi tiết chỗ ở thành công',
                    'Chi tiết chỗ ở đã được tạo thành công'
                );
            });
    }

    protected updateAccommodationDetail() {
        const accommodationDetail: UpdateAccommodationDetails = {
            id: this.idAccommodationDetailUpdating,
            accommodation_id: this.accommodationId,
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
            price: `${this.formAccommodationDetail.get('price')?.value || 0}`,
            facilities: this.getSelectedFacilityIds(),
        };

        this.accommodationDetailService
            .updateAccommodationDetail(accommodationDetail)
            .subscribe((response) => {
                // console.log(response);
                this.accommodationDetails = this.accommodationDetails.map(
                    (accommodationDetail) => {
                        if (accommodationDetail.id === response.data.id) {
                            return response.data;
                        } else {
                            return accommodationDetail;
                        }
                    }
                );
                this.showToast(
                    'success',
                    'Cập nhật chi tiết chỗ ở thành công',
                    'Chi tiết chỗ ở đã được cập nhật thành công'
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
                this.showToast(
                    'success',
                    'Xóa chi tiết chỗ ở thành công',
                    'Chi tiết chỗ ở đã được xóa thành công'
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
