import { AfterViewInit, Component, ElementRef, inject, Injector, OnInit, QueryList, ViewChildren } from '@angular/core';
import {
    FormControl,
    FormGroup,
    FormsModule,
    ReactiveFormsModule,
    Validators,
} from '@angular/forms';
import { DomSanitizer, SafeHtml } from '@angular/platform-browser';

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
import { ManagerService } from '../../../services/admin/manager.service';
import { GetAccommodationsOfManagerByAdmin } from '../../../models/admin/manager.model';

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
    ],
    templateUrl: './accommodation.component.html',
    styleUrl: './accommodation.component.scss',
    providers: [
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

    protected facilities!: Facility[];
    protected columns: string[] = [
        'Name',
        'Country',
        'City',
        'District',
        'Address',
        'Description',
        'Rating',
        'Google Map',
        'Image',
        'Is Verified',
        'Is Deleted',
    ];
    protected readonly tools = TUI_EDITOR_DEFAULT_TOOLS;
    protected managerId: string = '';
    protected accommodations: GetAccommodationsOfManagerByAdmin[] = [];
    protected cities: City[] = [];
    protected districts: District[] = [];
    protected cityNames: string[] = [];
    protected districtNames: string[] = [];
    protected cityName: string = '';
    protected citySlug: string = '';
    protected districtName: string = '';
    protected districtSlug: string = '';
    protected showFullMap: { [id: string]: boolean } = {};;
    protected elList: { [id: string]: any } = {};
    protected showButtonStates: { [id: string]: boolean } = {};

    private readonly dialogs = inject(TuiDialogService);

    // protected formAccommodation = new FormGroup({
    //     name: new FormControl('', Validators.required),
    //     country: new FormControl('Việt Nam'),
    //     city: new FormControl('', Validators.required),
    //     district: new FormControl({ value: '', disabled: true }, Validators.required),
    //     address: new FormControl('', Validators.required),
    //     rating: new FormControl(0, [
    //         Validators.required,
    //         Validators.min(1),
    //         Validators.max(5),
    //     ]),
    //     description: new FormControl('', Validators.required),
    //     googleMap: new FormControl('', Validators.required),
    // });

    protected timePeriods = tuiCreateTimePeriods();

    constructor(
        private accommodationService: ManagerService,
        private addressService: AddressService,
        private sanitizer: DomSanitizer,
        private route: ActivatedRoute,
    ) { }

    ngOnInit() {
        this.route.params.subscribe((params) => {
            this.managerId = params['id'];
            this.accommodationService.getAccommodationsOfManagerByAdmin(this.managerId).subscribe((response) => {
                this.accommodations = response.data;

                console.log(this.accommodations);
            })
        });

        // this.formAccommodation.get('city')?.valueChanges.subscribe((selectedCity: string | null) => {
        //     // console.log("selected city: ", selectedCity);

        //     this.onCitySelected(selectedCity);

        //     this.formAccommodation.get('district')?.reset();
        //     // this.onCityChange(selectedCity);
        // });

        // this.formAccommodation.get('district')?.valueChanges.subscribe((selectedDistrict: string | null) => {
        //     // console.log("selected district: ", selectedDistrict);

        //     this.onDistrictSelected(selectedDistrict);
        //     // this.onCityChange(selectedCity);
        // });

        this.addressService.getCities().subscribe((res) => {
            this.cities = res.data;
            this.cityNames = res.data.map(city => city.name);
            // this.initFormValueChanges();
        });
    }

    // private onCitySelected(selectedCityName: string | null): void {
    //     const selectedCity = this.cities.find(city => city.name === selectedCityName);

    //     if (selectedCity) {
    //         this.citySlug = selectedCity.slug;
    //         this.districts = selectedCity.level2s;
    //         this.districtNames = this.districts.map(d => d.name);
    //         this.formAccommodation.get('district')?.enable();
    //     } else {
    //         this.citySlug = '';
    //         this.districts = [];
    //         this.districtNames = [];
    //         this.formAccommodation.get('district')?.disable();
    //     }
    // }

    // private onDistrictSelected(selectedDistrictName: string | null): void {
    //     const selectedDistrict = this.districts.find(d => d.name === selectedDistrictName);
    //     this.districtSlug = selectedDistrict?.slug ?? '';
    // }

    changeCitySlugToName(slug: string): string {
        const city = this.cities.find(city => city.slug === slug);

        return city?.name ?? '';
    }

    changeDistrictSlugToName(citySlug: string, districtSlug: string): string {
        // console.log("cities: ", this.cities)
        const city = this.cities.find(city => city.slug === citySlug);
        let districts = city?.level2s ?? [];

        // console.log("districts", districts);
        let district = districts.find(district => district.slug === districtSlug);
        // console.log(district);

        return district?.name ?? '';
    }

    protected getDescription(html: string): SafeHtml {
        return this.sanitizer.bypassSecurityTrustHtml(html);
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

                console.log(this.showButtonStates);
            });
        });
    }

    ngAfterViewInit(): void {
        this.descEls.changes.subscribe(() => {
            setTimeout(() => this.checkDescriptionOverflow(), 0);
        });
    }
}
