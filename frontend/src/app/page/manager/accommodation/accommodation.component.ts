import {
    AfterViewInit,
    Component,
    ElementRef,
    inject,
    Injector,
    OnInit,
    QueryList,
    ViewChildren,
} from '@angular/core';
import {
    FormControl,
    FormGroup,
    FormsModule,
    ReactiveFormsModule,
    Validators,
} from '@angular/forms';
import {
    DomSanitizer,
    SafeHtml,
    SafeResourceUrl,
} from '@angular/platform-browser';
import { AccommodationService } from '../../../services/manager/accommodation.service';
import {
    Accommodation,
    CreateAccommodation,
    UpdateAccommodation,
} from '../../../models/manager/accommodation.model';
import { TuiTable } from '@taiga-ui/addon-table';
import {
    TuiIcon,
    TuiButton,
    TuiDialogService,
    TuiTextfield,
    TuiAppearance,
} from '@taiga-ui/core';
import type { PolymorpheusContent } from '@taiga-ui/polymorpheus';
import { TuiInputModule, TuiSelectModule } from '@taiga-ui/legacy';
import {
    TuiConfirmService,
    TuiFiles,
    tuiCreateTimePeriods,
    TuiRating,
    TuiSelect,
    TuiTooltip,
    TuiDataListWrapperComponent,
    TuiDataListWrapper,
} from '@taiga-ui/kit';
import { TuiResponsiveDialogService } from '@taiga-ui/addon-mobile';
import { TuiCardLarge } from '@taiga-ui/layout';
import {
    TUI_EDITOR_DEFAULT_EXTENSIONS,
    TUI_EDITOR_DEFAULT_TOOLS,
    TUI_EDITOR_EXTENSIONS,
    TuiEditor,
} from '@taiga-ui/editor';
import { RouterLink } from '@angular/router';
import { TuiInputTimeModule } from '@taiga-ui/legacy';
import { Facility } from '../../../models/facility/facility.model';
import { FacilityService } from '../../../services/facility/facility.service';
import { AddressService } from '../../../services/address/address.service';
import { City, District } from '../../../models/address/address.model';
import { NavbarComponent } from '../../../components/navbar/navbar.component';
import { MessageService } from 'primeng/api';
import { Toast } from 'primeng/toast';
import { ButtonModule } from 'primeng/button';

@Component({
    standalone: true,
    selector: 'app-accommodation',
    imports: [
        TuiIcon,
        TuiTooltip,
        TuiTable,
        TuiButton,
        TuiInputModule,
        FormsModule,
        ReactiveFormsModule,
        TuiTextfield,
        TuiAppearance,
        TuiCardLarge,
        TuiFiles,
        TuiEditor,
        RouterLink,
        TuiInputTimeModule,
        TuiRating,
        TuiSelect,
        TuiSelectModule,
        TuiDataListWrapperComponent,
        TuiDataListWrapper,
        NavbarComponent,
        Toast,
        ButtonModule,
    ],
    templateUrl: './accommodation.component.html',
    styleUrl: './accommodation.component.scss',
    providers: [
        MessageService,
        TuiConfirmService,
        {
            provide: TuiDialogService,
            useExisting: TuiResponsiveDialogService,
        },
        {
            provide: TUI_EDITOR_EXTENSIONS,
            deps: [Injector],
            useFactory: (injector: Injector) => [
                ...TUI_EDITOR_DEFAULT_EXTENSIONS,
                import('@taiga-ui/editor').then(({ setup }) =>
                    setup({ injector })
                ),
            ],
        },
    ],
})
export class AccommodationComponent implements OnInit, AfterViewInit {
    @ViewChildren('descEl') descEls!: QueryList<ElementRef<HTMLDivElement>>;

    protected accommodations!: Accommodation[];
    protected facilities!: Facility[];
    protected columns: string[] = [
        'ID',
        'Name',
        'Country',
        'City',
        'District',
        'Address',
        'Description',
        'Rating',
        'Google Map',
        'Is Verified',
        'Is Deleted',
        'Image',
        'Action',
        'Show Accommodation Detail',
    ];
    protected readonly tools = TUI_EDITOR_DEFAULT_TOOLS;
    protected idAccommodationUpdating = '';
    protected cities: City[] = [];
    protected districts: District[] = [];
    protected cityNames: string[] = [];
    protected districtNames: string[] = [];
    protected cityName: string = '';
    protected citySlug: string = '';
    protected districtName: string = '';
    protected districtSlug: string = '';
    protected showFullMap: { [id: string]: boolean } = {};
    protected elList: { [id: string]: any } = {};
    protected showButtonStates: { [id: string]: boolean } = {};

    private readonly dialogs = inject(TuiDialogService);

    protected formAccommodation = new FormGroup({
        name: new FormControl('', Validators.required),
        country: new FormControl('Việt Nam'),
        city: new FormControl('', Validators.required),
        district: new FormControl(
            { value: '', disabled: true },
            Validators.required
        ),
        address: new FormControl('', Validators.required),
        rating: new FormControl(0, [
            Validators.required,
            Validators.min(1),
            Validators.max(5),
        ]),
        description: new FormControl('', Validators.required),
        googleMap: new FormControl('', Validators.required),
    });
    protected formFacilities = new FormGroup({});

    protected timePeriods = tuiCreateTimePeriods();

    constructor(
        private accommodationService: AccommodationService,
        private facilityService: FacilityService,
        private addressService: AddressService,
        private messageService: MessageService,
        private sanitizer: DomSanitizer
    ) {}

    // Method để sanitize URL
    getSafeUrl(url: string): SafeResourceUrl {
        return this.sanitizer.bypassSecurityTrustResourceUrl(url);
    }

    showToast(
        severity: 'success' | 'info' | 'warn' | 'error',
        summary: string,
        detail: string
    ): void {
        this.messageService.add({ severity, summary, detail });
    }

    ngOnInit() {
        this.formAccommodation
            .get('city')
            ?.valueChanges.subscribe((selectedCity: string | null) => {
                this.onCitySelected(selectedCity);
                this.formAccommodation.get('district')?.reset();
            });

        this.formAccommodation
            .get('district')
            ?.valueChanges.subscribe((selectedDistrict: string | null) => {
                this.onDistrictSelected(selectedDistrict);
            });

        this.addressService.getCities().subscribe((res) => {
            this.cities = res.data;
            this.cityNames = res.data.map((city) => city.name);
        });

        this.accommodationService.getAccommodations().subscribe((response) => {
            console.log(response);
            this.accommodations = response.data;
        });

        this.facilityService.getFacilities().subscribe((response) => {
            this.facilities = response.data;
            this.createFacilityControls();
        });
    }

    private onCitySelected(selectedCityName: string | null): void {
        const selectedCity = this.cities.find(
            (city) => city.name === selectedCityName
        );

        if (selectedCity) {
            this.citySlug = selectedCity.slug;
            this.districts = selectedCity.level2s;
            this.districtNames = this.districts.map((d) => d.name);
            this.formAccommodation.get('district')?.enable();
        } else {
            this.citySlug = '';
            this.districts = [];
            this.districtNames = [];
            this.formAccommodation.get('district')?.disable();
        }
    }

    private onDistrictSelected(selectedDistrictName: string | null): void {
        const selectedDistrict = this.districts.find(
            (d) => d.name === selectedDistrictName
        );
        this.districtSlug = selectedDistrict?.slug ?? '';
    }

    changeCitySlugToName(slug: string): string {
        const city = this.cities.find((city) => city.slug === slug);

        return city?.name ?? '';
    }

    changeDistrictSlugToName(citySlug: string, districtSlug: string): string {
        const city = this.cities.find((city) => city.slug === citySlug);
        let districts = city?.level2s ?? [];
        let district = districts.find(
            (district) => district.slug === districtSlug
        );
        return district?.name ?? '';
    }

    private createFacilityControls() {
        const facilityControls: { [key: string]: FormControl<boolean> } = {};

        if (!this.facilities || this.facilities.length === 0) {
            this.formFacilities = new FormGroup(facilityControls);
            return;
        }

        this.facilities.forEach((facility) => {
            facilityControls[facility.id] = new FormControl<boolean>(false, {
                nonNullable: true,
            });
        });

        this.formFacilities = new FormGroup(facilityControls);
    }

    getSelectedFacilityIds(): string[] {
        if (!this.facilities || this.facilities.length === 0) {
            return [];
        }
        return this.facilities
            .filter(
                (facility) =>
                    this.formFacilities.get(facility.id)?.value === true
            )
            .map((facility) => facility.id);
    }

    protected openDialogCreate(content: PolymorpheusContent): void {
        this.formAccommodation.reset();
        this.formFacilities.reset();
        this.formAccommodation.patchValue({ country: 'Việt Nam' });
        this.formAccommodation.get('district')?.disable();

        this.dialogs
            .open(content, {
                label: 'Create Accommodation',
            })
            .subscribe({
                complete: () => {
                    this.formAccommodation.reset();
                },
            });
    }

    protected openDialogUpdate(
        content: PolymorpheusContent,
        accommodation: Accommodation
    ) {
        this.formAccommodation.reset();
        this.formFacilities.reset();

        this.formAccommodation.patchValue({
            name: accommodation.name,
            city: this.changeCitySlugToName(accommodation.city),
            country: 'Việt Nam',
            district: this.changeDistrictSlugToName(
                accommodation.city,
                accommodation.district
            ),
            description: accommodation.description,
            googleMap: accommodation.google_map,
            address: accommodation.address,
            rating: accommodation.rating,
        });

        const selectedCity = this.cities.find(
            (city) => city.name === accommodation.city
        );
        if (selectedCity) {
            this.districts = selectedCity.level2s;
            this.districtNames = this.districts.map((d) => d.name);
        }

        this.setFacilityValues(accommodation.facilities);

        this.idAccommodationUpdating = accommodation.id;

        this.dialogs
            .open(content, {
                label: 'Update Accommodation',
            })
            .subscribe({
                complete: () => {
                    this.formAccommodation.reset();
                },
            });
    }

    private setFacilityValues(accommodationFacilities: Facility[]) {
        const facilityValues: { [key: string]: boolean } = {};
        Object.keys(this.formFacilities.controls).forEach((facilityId) => {
            facilityValues[facilityId] = false;
        });

        accommodationFacilities.forEach((facilityId) => {
            if (facilityValues.hasOwnProperty(facilityId.id)) {
                facilityValues[facilityId.id] = true;
            }
        });

        this.formFacilities.patchValue(facilityValues);
    }

    protected getDescription(html: string): SafeHtml {
        return this.sanitizer.bypassSecurityTrustHtml(html);
    }

    protected createAccommodation() {
        const accommodation: CreateAccommodation = {
            name: this.formAccommodation.get('name')?.value || '',
            city: this.citySlug,
            country: 'Việt Nam',
            district: this.districtSlug,
            address: this.formAccommodation.get('address')?.value || '',
            description: this.formAccommodation.get('description')?.value || '',
            google_map: this.formAccommodation.get('googleMap')?.value || '',
            rating: this.formAccommodation.get('rating')?.value || 0,
            facilities: this.getSelectedFacilityIds(),
        };

        if (this.formAccommodation.invalid) {
            this.formAccommodation.markAllAsTouched();
            return;
        }

        this.accommodationService.createAccommodation(accommodation).subscribe({
            next: (response) => {
                this.accommodations.push(response.data);
                this.formAccommodation.reset();
                this.formFacilities.reset();
                this.accommodations = [...this.accommodations];
                this.checkDescriptionOverflow();
                this.showToast(
                    'success',
                    'Khách sạn đã được tạo thành công',
                    'Bạn có thể xem chi tiết khách sạn trong danh sách'
                );
            },
            error: (error) => {
                console.error('Error creating accommodation:', error);
                this.showToast(
                    'error',
                    'Tạo khách sạn thất bại',
                    'Vui lòng thử lại sau'
                );
            },
        });
    }

    protected updateAccommodation() {
        const accommodation: UpdateAccommodation = {
            id: this.idAccommodationUpdating,
            name: this.formAccommodation.get('name')?.value || '',
            city: this.citySlug,
            country: 'Việt Nam',
            district: this.districtSlug,
            address: this.formAccommodation.get('address')?.value || '',
            description: this.formAccommodation.get('description')?.value || '',
            google_map: this.formAccommodation.get('googleMap')?.value || '',
            rating: this.formAccommodation.get('rating')?.value || 0,
            facilities: this.getSelectedFacilityIds(),
        };

        this.accommodationService
            .updateAccommodation(accommodation)
            .subscribe((response) => {
                this.accommodations = this.accommodations.map(
                    (accommodation) => {
                        if (accommodation.id === response.data.id) {
                            this.showToast(
                                'success',
                                'Cập nhật khách sạn thành công',
                                'Bạn có thể xem chi tiết khách sạn trong danh sách'
                            );
                            return response.data;
                        } else {
                            this.showToast(
                                'error',
                                'Cập nhật khách sạn thất bại',
                                'Cập nhật khách sạn thất bại. Vui lòng thử lại sau'
                            );
                            return accommodation;
                        }
                    }
                );
            });
    }

    protected deleteAccommodation(id: string) {
        this.accommodationService.deleteAccommodation(id).subscribe((_) => {
            this.accommodations = this.accommodations.filter(
                (accommodation) => accommodation.id !== id
            );
            this.showToast('success', 'Xoá khách sạn thành công', 'Xoá khách sạn thành công');
        });
    }

    protected toggleDescription(id: string): void {
        this.showFullMap[id] = !this.showFullMap[id];
    }

    protected isDescriptionShown(id: string): boolean {
        return !!this.showFullMap[id];
    }

    private checkDescriptionOverflow() {
        setTimeout(() => {
            this.descEls.forEach((elRef) => {
                const el = elRef.nativeElement;
                const id = el.getAttribute('data-id');
                if (id) {
                    this.showButtonStates[id] = el.scrollHeight > 60;
                }
            });
        });
    }

    ngAfterViewInit(): void {
        this.descEls.changes.subscribe(() => {
            setTimeout(() => this.checkDescriptionOverflow(), 0);
        });
    }
}
