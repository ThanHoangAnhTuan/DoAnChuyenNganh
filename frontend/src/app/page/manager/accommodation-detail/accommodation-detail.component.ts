import { FacilityDetailService } from './../../../services/facility-detail/facility-detail.service';
import { Component, inject, OnInit } from '@angular/core';
import { Accommodation } from '../../../models/manager/accommodation.model';
import {
    TuiAppearance,
    TuiButton,
    TuiDataList,
    TuiDialogService,
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
import { finalize, forkJoin } from 'rxjs';
import { LoaderComponent } from '../../../components/loader/loader.component';

@Component({
    selector: 'app-accommodation-detail',
    imports: [
        TuiTable,
        FormsModule,
        ReactiveFormsModule,
        TuiButton,
        TuiInputModule,
        FormsModule,
        ReactiveFormsModule,
        TuiTextfield,
        TuiAppearance,
        TuiCardLarge,
        TuiCheckbox,
        RouterLink,
        TuiInputNumber,
        TuiDataList,
        TuiSelect,
        TuiChevron,
        NavbarComponent,
        Toast,
        ButtonModule,
        LoaderComponent,
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
        'Price',
        'Available Rooms',
        'Image',
        'Room',
        'Action',
    ];
    protected readonly baseUrl: string = 'http://localhost:8080/uploads/';
    isLoading: boolean = false;
    protected idAccommodationDetailUpdating = '';
    private readonly dialogs = inject(TuiDialogService);
    protected accommodationId: string = '';
    protected formAccommodationDetail = new FormGroup({
        name: new FormControl<string | ''>('', Validators.required),
        guests: new FormControl<number | 0>(0, [Validators.required, Validators.min(1)]),
        singleBed: new FormControl<number | 0>(0),
        doubleBed: new FormControl<number | 0>(0),
        largeDoubleBed: new FormControl<number | 0>(0),
        extraLargeDoubleBed: new FormControl<number | 0>(0),
        price: new FormControl<number | 0>(0, Validators.min(1)),
        accommodationId: new FormControl<string | ''>(''),
        discountId: new FormControl<string | ''>(''),
        facilityDetails: new FormControl<string | ''>(''),
    });
    protected formFacilityDetail = new FormGroup({});
    protected readonly resetFormAccommodationDetail = {
        accommodationId: '',
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

    protected readonly contentAccommodation: PolymorpheusContent<
        TuiContext<string | null>
    > = ({ $implicit: id }) =>
        this.accommodationItems.find((item) => item.id === id)?.name ?? '';
    protected readonly discountItems: readonly DiscountSelect[] = [];
    protected readonly contentDiscount: PolymorpheusContent<
        TuiContext<string | null>
    > = ({ $implicit: id }) =>
        this.discountItems.find((item) => item.id === id)?.name ?? '';
    constructor(
        private route: ActivatedRoute,
        private accommodationDetailService: AccommodationDetailService,
        private accommodationService: AccommodationService,
        private facilityDetailService: FacilityDetailService,
        private messageService: MessageService
    ) {}

    ngOnInit() {
        this.route.params.subscribe((params) => {
            this.accommodationId = params['id'];

            this.isLoading = true;

            // Sử dụng forkJoin để chờ tất cả API calls hoàn thành
            forkJoin({
                accommodationDetails:
                    this.accommodationDetailService.getAccommodationDetailsByManager(
                        params['id']
                    ),
                accommodations: this.accommodationService.getAccommodations(),
                facilities: this.facilityDetailService.getFacilityDetail(),
            })
                .pipe(
                    finalize(() => {
                        this.isLoading = false; // Tắt loading sau khi tất cả hoàn thành
                    })
                )
                .subscribe({
                    next: (results) => {
                        // Xử lý accommodation details
                        this.accommodationDetails =
                            results.accommodationDetails.data;

                        // Xử lý accommodations
                        this.accommodationItems =
                            results.accommodations.data.map((item) => ({
                                id: item.id,
                                name: item.name,
                            }));

                        // Xử lý facilities
                        this.facilities = results.facilities.data;
                        this.createFacilityControls();
                    },
                    error: (error) => {
                        console.error('Error loading data:', error);
                        this.showToast(
                            'error',
                            'Thất bại',
                            error.error.message ||
                                'Đã xảy ra lỗi khi tải dữ liệu. Vui lòng thử lại sau.'
                        );
                    },
                });
        });
    }

    showToast(
        severity: 'success' | 'info' | 'warn' | 'error',
        summary: string,
        detail: string
    ): void {
        this.messageService.add({ severity, summary, detail });
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

        this.formFacilityDetail = new FormGroup(facilityControls);
    }

    getSelectedFacilityIds(): string[] {
        if (!this.facilities || this.facilities.length === 0) {
            return [];
        }
        return this.facilities
            .filter(
                (facilities) =>
                    this.formFacilityDetail.get(facilities.id)?.value === true
            )
            .map((facilities) => facilities.id);
    }

    protected openDialogCreate(content: PolymorpheusContent): void {
        this.formAccommodationDetail.reset(this.resetFormAccommodationDetail);
        this.formFacilityDetail.reset();
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
        if (this.formAccommodationDetail.invalid) {
            this.formAccommodationDetail.markAllAsTouched();
            return;
        }
        this.isLoading = true;
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
            price: `${this.formAccommodationDetail.get('price')?.value || 0}`,
            accommodation_id: this.accommodationId,
            discount_id:
                this.formAccommodationDetail.get('discountId')?.value || '',
            facilities: this.getSelectedFacilityIds(),
        };
        this.accommodationDetailService
            .createAccommodationDetail(accommodationDetail)
            .pipe(
                finalize(() => {
                    this.isLoading = false;
                })
            )
            .subscribe({
                next: (response) => {
                    this.accommodationDetails.push(response.data);
                    this.formAccommodationDetail.reset(
                        this.resetFormAccommodationDetail
                    );
                    this.formFacilityDetail.reset();
                    this.showToast(
                        'success',
                        'Thành công',
                        'Tạo loại phòng thành công'
                    );
                },
                error: (error) => {
                    this.showToast(
                        'error',
                        'Thất bại',
                        error.error.message ||
                            'Đã xảy ra lỗi khi tạo loại phòng. Vui lòng thử lại sau.'
                    );
                },
            });
    }

    protected updateAccommodationDetail() {
        if (this.formAccommodationDetail.invalid) {
            this.formAccommodationDetail.markAllAsTouched();
            return;
        }
        this.isLoading = true;
        const accommodationDetail: UpdateAccommodationDetails = {
            id: this.idAccommodationDetailUpdating,
            accommodation_id: this.accommodationId,
            name: this.formAccommodationDetail.get('name')?.value || '',
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
        console.log("accommodationDetail: ", accommodationDetail);
        this.accommodationDetailService
            .updateAccommodationDetail(accommodationDetail)
            .pipe(
                finalize(() => {
                    this.isLoading = false;
                })
            )
            .subscribe({
                next: (response) => {
                    this.accommodationDetails = this.accommodationDetails.map(
                        (detail) => {
                            if (detail.id === response.data.id) {
                                return response.data;
                            }
                            return detail;
                        }
                    );
                    this.showToast(
                        'success',
                        'Thành công',
                        'Cập nhật loại phòng thành công'
                    );
                },
                error: (error) => {
                    this.showToast(
                        'error',
                        'Thất bại',
                        error.error.message ||
                            'Đã xảy ra lỗi khi cập nhật loại phòng. Vui lòng thử lại sau.'
                    );
                },
            });
    }

    protected deleteAccommodationDetail(id: string) {
        this.isLoading = true;
        this.accommodationDetailService
            .deleteAccommodationDetail(id)
            .pipe(
                finalize(() => {
                    this.isLoading = false;
                })
            )
            .subscribe({
                next: (response) => {
                    this.accommodationDetails =
                        this.accommodationDetails.filter(
                            (detail) => detail.id !== id
                        );
                    this.showToast(
                        'success',
                        'Thành công',
                        'Xóa loại phòng thành công'
                    );
                },
                error: (error) => {
                    this.showToast(
                        'error',
                        'Thất bại',
                        error.error.message ||
                            'Đã xảy ra lỗi khi xoá loại phòng. Vui lòng thử lại sau.'
                    );
                },
            });
    }
}
