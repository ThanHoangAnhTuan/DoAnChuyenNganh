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
    TuiGroup,
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
import { ActivatedRoute, RouterLink } from '@angular/router';
import { TuiInputTimeModule } from '@taiga-ui/legacy';
import { Facility } from '../../../models/facility/facility.model';
import { FacilityService } from '../../../services/facility/facility.service';
import { AddressService } from '../../../services/address/address.service';
import { City, District } from '../../../models/address/address.model';
import { NavbarComponent } from '../../../components/navbar/navbar.component';
import { MessageService } from 'primeng/api';
import { Toast } from 'primeng/toast';
import { ButtonModule } from 'primeng/button';
import { Ripple } from 'primeng/ripple';
import { RoomService } from '../../../services/manager/room.service';
import { CreateRoom, Room } from '../../../models/manager/room.model';

@Component({
    selector: 'app-room',
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
        Ripple,
    ],
    templateUrl: './room.component.html',
    styleUrl: './room.component.scss',
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
export class RoomComponent implements OnInit, AfterViewInit {
    // @ViewChildren('descEl') descEls!: QueryList<ElementRef<HTMLDivElement>>;

    // protected accommodations!: Accommodation[];
    // protected facilities!: Facility[];
    protected columns: string[] = [
        'ID',
        'Name',
        'Status',
        // 'City',
        // 'District',
        // 'Address',
        // 'Description',
        // 'Rating',
        // 'Google Map',
        // 'Image',
        'Action',
        // 'Show Accommodation Detail',
    ];
    // protected readonly tools = TUI_EDITOR_DEFAULT_TOOLS;
    // protected idAccommodationUpdating = '';
    // protected cities: City[] = [];
    // protected districts: District[] = [];
    // protected cityNames: string[] = [];
    // protected districtNames: string[] = [];
    // protected cityName: string = '';
    // protected citySlug: string = '';
    // protected districtName: string = '';
    // protected districtSlug: string = '';
    // protected showFullMap: { [id: string]: boolean } = {};
    // protected elList: { [id: string]: any } = {};
    // protected showButtonStates: { [id: string]: boolean } = {};

    protected status: string[] = ['available', 'unavailable', 'occupied'];

    private readonly dialogs = inject(TuiDialogService);

    protected formCreateRoom = new FormGroup({
        prefix: new FormControl('', Validators.required),
        quantity: new FormControl<number | null>(null, Validators.required),
    });

    protected formUpdateRoom = new FormGroup({
        name: new FormControl('', Validators.required),
        status: new FormControl('', Validators.required),
    });
    // protected formFacilities = new FormGroup({});

    // protected timePeriods = tuiCreateTimePeriods();

    protected accommodationDetailId: string = '';
    protected accommodationRoomId: string = '';
    protected rooms: Room[] = [];

    constructor(
        private accommodationService: AccommodationService,
        private facilityService: FacilityService,
        private addressService: AddressService,
        private roomService: RoomService,
        private messageService: MessageService,
        private sanitizer: DomSanitizer,
        private route: ActivatedRoute
    ) {}

    // // Method để sanitize URL
    // getSafeUrl(url: string): SafeResourceUrl {
    //     return this.sanitizer.bypassSecurityTrustResourceUrl(url);
    // }

    showToast(
        severity: 'success' | 'info' | 'warn' | 'error',
        summary: string,
        detail: string
    ): void {
        this.messageService.add({ severity, summary, detail });
    }

    ngAfterViewInit(): void {
        //     this.descEls.changes.subscribe(() => {
        //         setTimeout(() => this.checkDescriptionOverflow(), 0);
        //     });
    }

    ngOnInit() {
        this.route.params.subscribe((params) => {
            this.accommodationDetailId = params['id'];
            this.roomService
                .getAccommodationRooms(params['id'])
                .subscribe((response) => {
                    this.rooms = response.data;
                });
        });
        //     // this.addressService.getCities().subscribe((data: City[]) => {
        //     //     this.cities = data;

        //     //     this.cityNames = this.cities.map(city => city.name);
        //     //     console.log("data:", this.cityNames);
        //     // })

        //     // this.formAccommodation.get('city')?.valueChanges.subscribe((selectedCityName: string | null) => {
        //     //     const selectedCity = this.cities.find(city => city.name === selectedCityName);

        //     //     if (selectedCity) {
        //     //         // console.log("selected city: ", selectedCity);

        //     //         this.citySlug = selectedCity.slug;
        //     //         this.districts = selectedCity.level2s;
        //     //         // console.log("districts: ", this.districts);

        //     //         this.districtNames = this.districts.map(d => d.name);
        //     //         // console.log("districts: ", this.districtNames);

        //     //         this.formAccommodation.get('district')?.enable();
        //     //     } else {
        //     //         this.citySlug = '';
        //     //         this.districts = [];
        //     //         this.districtNames = [];
        //     //         this.formAccommodation.get('district')?.disable();
        //     //     }
        //     // })

        //     // this.formAccommodation.get('district')?.valueChanges.subscribe((selectedDistrictName: string | null) => {
        //     //     const selectedDistrict = this.districts.find(district => district.name === selectedDistrictName);

        //     //     if (selectedDistrict) {
        //     //         // console.log("selected district: ", selectedDistrict);

        //     //         this.districtSlug = selectedDistrict.slug;
        //     //     } else {
        //     //         this.districtSlug = '';
        //     //     }
        //     // })

        //     this.formAccommodation
        //         .get('city')
        //         ?.valueChanges.subscribe((selectedCity: string | null) => {
        //             // console.log("selected city: ", selectedCity);

        //             this.onCitySelected(selectedCity);

        //             this.formAccommodation.get('district')?.reset();
        //             // this.onCityChange(selectedCity);
        //         });

        //     this.formAccommodation
        //         .get('district')
        //         ?.valueChanges.subscribe((selectedDistrict: string | null) => {
        //             // console.log("selected district: ", selectedDistrict);

        //             this.onDistrictSelected(selectedDistrict);
        //             // this.onCityChange(selectedCity);
        //         });

        //     this.addressService.getCities().subscribe((res) => {
        //         this.cities = res.data;
        //         this.cityNames = res.data.map((city) => city.name);
        //         // this.initFormValueChanges();
        //     });

        //     this.accommodationService.getAccommodations().subscribe((response) => {
        //         this.accommodations = response.data;
        //     });

        //     this.facilityService.getFacilities().subscribe((response) => {
        //         this.facilities = response.data;
        //         this.createFacilityControls();
        //     });
        // }

        // // private initFormValueChanges(): void {
        // //     this.formAccommodation.get('city')?.valueChanges
        // //         .subscribe(
        // //             this.onCitySelected.bind(this)
        // //         );

        // //     this.formAccommodation.get('district')?.valueChanges
        // //         .subscribe(this.onDistrictSelected.bind(this));
    }

    // private onCitySelected(selectedCityName: string | null): void {
    //     const selectedCity = this.cities.find(
    //         (city) => city.name === selectedCityName
    //     );

    //     if (selectedCity) {
    //         this.citySlug = selectedCity.slug;
    //         this.districts = selectedCity.level2s;
    //         this.districtNames = this.districts.map((d) => d.name);
    //         this.formAccommodation.get('district')?.enable();
    //     } else {
    //         this.citySlug = '';
    //         this.districts = [];
    //         this.districtNames = [];
    //         this.formAccommodation.get('district')?.disable();
    //     }
    // }

    // private onDistrictSelected(selectedDistrictName: string | null): void {
    //     const selectedDistrict = this.districts.find(
    //         (d) => d.name === selectedDistrictName
    //     );
    //     this.districtSlug = selectedDistrict?.slug ?? '';
    // }

    // changeCitySlugToName(slug: string): string {
    //     const city = this.cities.find((city) => city.slug === slug);

    //     return city?.name ?? '';
    // }

    // changeDistrictSlugToName(citySlug: string, districtSlug: string): string {
    //     // console.log("cities: ", this.cities)
    //     const city = this.cities.find((city) => city.slug === citySlug);
    //     let districts = city?.level2s ?? [];

    //     // console.log("districts", districts);
    //     let district = districts.find(
    //         (district) => district.slug === districtSlug
    //     );
    //     // console.log(district);

    //     return district?.name ?? '';
    // }

    // private createFacilityControls() {
    //     const facilityControls: { [key: string]: FormControl<boolean> } = {};

    //     if (!this.facilities || this.facilities.length === 0) {
    //         this.formFacilities = new FormGroup(facilityControls);
    //         return;
    //     }

    //     this.facilities.forEach((facility) => {
    //         facilityControls[facility.id] = new FormControl<boolean>(false, {
    //             nonNullable: true,
    //         });
    //     });

    //     // Tạo FormGroup mới với các controls
    //     this.formFacilities = new FormGroup(facilityControls);
    // }

    // getSelectedFacilityIds(): string[] {
    //     if (!this.facilities || this.facilities.length === 0) {
    //         return [];
    //     }
    //     return this.facilities
    //         .filter(
    //             (facility) =>
    //                 this.formFacilities.get(facility.id)?.value === true
    //         )
    //         .map((facility) => facility.id);
    // }

    protected openDialogCreate(content: PolymorpheusContent): void {
        this.formCreateRoom.reset();
        //     this.formFacilities.reset();
        //     this.formAccommodation.patchValue({ country: 'Việt Nam' });
        //     this.formAccommodation.get('district')?.disable();
        this.dialogs
            .open(content, {
                label: 'Create Room',
            })
            .subscribe({
                complete: () => {
                    this.formCreateRoom.reset();
                },
            });
    }

    protected openDialogUpdate(content: PolymorpheusContent, room: Room) {
        this.formUpdateRoom.reset();
        console.log(room);
        this.accommodationRoomId = room.id;
        this.formUpdateRoom.patchValue({
            name: room.name,
            status: room.status,
        });
        //     const selectedCity = this.cities.find(
        //         (city) => city.name === accommodation.city
        //     );
        //     if (selectedCity) {
        //         this.districts = selectedCity.level2s;
        //         this.districtNames = this.districts.map((d) => d.name);
        //     }
        //     console.log('accommodation: ', accommodation);
        //     this.setFacilityValues(accommodation.facilities);
        //     this.idAccommodationUpdating = accommodation.id;
        this.dialogs
            .open(content, {
                label: 'Update Room',
            })
            .subscribe({
                complete: () => {
                    this.formUpdateRoom.reset();
                },
            });
    }

    // private setFacilityValues(accommodationFacilities: Facility[]) {
    //     const facilityValues: { [key: string]: boolean } = {};
    //     Object.keys(this.formFacilities.controls).forEach((facilityId) => {
    //         facilityValues[facilityId] = false;
    //     });

    //     accommodationFacilities.forEach((facilityId) => {
    //         if (facilityValues.hasOwnProperty(facilityId.id)) {
    //             facilityValues[facilityId.id] = true;
    //         }
    //     });

    //     this.formFacilities.patchValue(facilityValues);
    // }

    // protected getDescription(html: string): SafeHtml {
    //     return this.sanitizer.bypassSecurityTrustHtml(html);
    // }

    protected createRoom() {
        // const accommodation: CreateAccommodation = {
        //         name: this.formAccommodation.get('name')?.value || '',
        //         city: this.citySlug,
        //         country: 'Việt Nam',
        //         district: this.districtSlug,
        //         address: this.formAccommodation.get('address')?.value || '',
        //         description: this.formAccommodation.get('description')?.value || '',
        //         google_map: this.formAccommodation.get('googleMap')?.value || '',
        //         rating: this.formAccommodation.get('rating')?.value || 0,
        //         facilities: this.getSelectedFacilityIds(),
        //     };
        console.log(this.formCreateRoom);
        if (this.formCreateRoom.invalid) {
            this.formCreateRoom.markAllAsTouched();
            return;
        }
        const room: CreateRoom = {
            prefix: this.formCreateRoom.get('prefix')?.value || '',
            quantity: Number(this.formCreateRoom.get('quantity')?.value) || 0,
            accommodation_type_id: this.accommodationDetailId,
        };
        // this.roomService
        //     .createAccommodationRoom(room)
        //     .subscribe((response) => {
        //     //         this.accommodations.push(response.data);
        //     //         this.formAccommodation.reset();
        //     //         this.formFacilities.reset();
        //     //         this.accommodations = [...this.accommodations]; // force trigger DOM update
        //     //         this.checkDescriptionOverflow();
        //     //         this.showToast(
        //     //             'success',
        //     //             'Khách sạn đã được tạo thành công',
        //     //             'Bạn có thể xem chi tiết khách sạn trong danh sách'
        //     //         );
        //     //     });
        this.roomService.createAccommodationRoom(room).subscribe({
            next: (response) => {
                console.log(response);
                this.rooms.push(...response.data);
                this.formCreateRoom.reset();
                this.showToast(
                    'success',
                    'Phòng đã được tạo thành công',
                    'Bạn có thể xem chi tiết phòng trong danh sách'
                );
            },
            error: (error) => {
                console.error('Error creating room:', error);
                this.showToast(
                    'error',
                    'Tạo phòng thất bại',
                    'Vui lòng thử lại sau'
                );
            },
        });
    }

    protected updateRoom() {
        if (this.formUpdateRoom.invalid) {
            this.formUpdateRoom.markAllAsTouched();
            return;
        }
        const room: Room = {
            id: this.accommodationRoomId,
            name: this.formUpdateRoom.get('name')?.value || '',
            status: this.formUpdateRoom.get('status')?.value || '',
        };
        this.roomService.updateAccommodationRoom(room).subscribe({
            next: (response) => {
                this.rooms = this.rooms.map((room) => {
                    if (room.id === response.data.id) {
                        return response.data;
                    } else {
                        return room;
                    }
                });
                this.showToast(
                    'success',
                    'Cập nhật phòng thành công',
                    'Bạn có thể xem chi tiết phòng trong danh sách'
                );
            },
            error: (error) => {
                console.error('Lỗi khi thêm đánh giá:', error);
                this.showToast(
                    'error',
                    'Cập nhật phòng thất bại',
                    'Cập nhật phòng thất bại, vui lòng thử lại sau'
                );
            },
        });
    }

    protected deleteRoom(id: string) {
        this.roomService.deleteAccommodationRoom(id).subscribe({
            next: (value) => {
                this.rooms = this.rooms.filter((room) => room.id !== id);
                this.showToast(
                    'success',
                    'Xoá phòng thành công',
                    'Bạn có thể xem thông tin trong danh sách'
                );
            },
            error: (err) => {
                this.showToast(
                    'error',
                    'Xoá phòng thất bại',
                    'Xoá phòng thất bại, vui lòng thử lại sau'
                );
            },
            complete() {},
        });
    }

    // protected toggleDescription(id: string): void {
    //     this.showFullMap[id] = !this.showFullMap[id];
    // }

    // protected isDescriptionShown(id: string): boolean {
    //     return !!this.showFullMap[id];
    // }

    // private checkDescriptionOverflow() {
    //     setTimeout(() => {
    //         this.descEls.forEach((elRef) => {
    //             const el = elRef.nativeElement;
    //             const id = el.getAttribute('data-id');

    //             if (id) {
    //                 this.showButtonStates[id] = el.scrollHeight > 60;
    //             }

    //             console.log(this.showButtonStates);
    //         });
    //     });
    // }

    // ngAfterViewInit(): void {
    //     this.descEls.changes.subscribe(() => {
    //         setTimeout(() => this.checkDescriptionOverflow(), 0);
    //     });
    // }
}
