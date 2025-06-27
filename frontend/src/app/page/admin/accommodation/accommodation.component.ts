import { AfterViewInit, Component, ElementRef, inject, Injector, OnInit, QueryList, ViewChildren } from '@angular/core';
import {
    FormsModule,
    ReactiveFormsModule,
} from '@angular/forms';
import { DomSanitizer, SafeHtml } from '@angular/platform-browser';
import { TuiTable } from '@taiga-ui/addon-table';
import {
    TuiIcon,
    TuiButton,
    TuiDialogService,
    TuiTextfield,
    TuiAppearance,
    TuiAlertService,
} from '@taiga-ui/core';
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
import { AddressService } from '../../../services/address/address.service';
import { City, District } from '../../../models/address/address.model';
import { NavbarComponent } from '../../../components/navbar/navbar.component';
import { ManagerService } from '../../../services/admin/manager.service';
import { GetAccommodationsOfManagerByAdmin, VerifyAccommodationInput } from '../../../models/admin/manager.model';
import { Accommodation } from '../../../models/manager/accommodation.model';

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
    protected isUpdateVerified: boolean = false;
    protected isUpdateDeleted: boolean = false;
    protected updateId: string = '';
    protected isModalConfirmVerifyOpen: boolean = false;
    protected isModalConfirmDeleteOpen: boolean = false;

    // private readonly dialogs = inject(TuiDialogService);

    protected timePeriods = tuiCreateTimePeriods();
    private readonly alerts = inject(TuiAlertService);

    protected getAlert(label: string, content: string): void {
        this.alerts
            .open(content, {
                label: label,
                appearance: 'success',
                autoClose: 5000,
            })
            .subscribe();
    }

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

        this.addressService.getCities().subscribe((res) => {
            this.cities = res.data;
            this.cityNames = res.data.map(city => city.name);
        });
    }

    private updateVerify(id: string, status: number) {
        let newVerify: VerifyAccommodationInput = {
            accommodation_id: id,
            status: Number(status)
        }

        this.accommodationService.updateVerified(newVerify).subscribe((response) => {
            const message = response.message;

            this.getAlert('Notification', message);
        })
    }

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

    protected changeVerifiedApply(id: string) {
        this.updateId = id;
        this.isUpdateVerified = true;
        this.isUpdateDeleted = false;
    }

    protected changeVerifiedFinish() {
        const id: string = this.updateId;
        const accommodation: any = this.accommodations.find(a => a.id === id);
        if (accommodation) {
            const status: number = accommodation.is_verified;

            this.updateVerify(id, status);
        }

        this.updateId = '';
        this.isUpdateVerified = false;
    }

    protected openVerifyConfirmModal() {
        this.isModalConfirmVerifyOpen = true;
    }

    protected closeVerifyConfirmModal() {
        this.isModalConfirmVerifyOpen = false;
        this.isUpdateVerified = false;
    }

    protected changeDeletedApply(id: string) {
        this.updateId = id;
        this.isUpdateVerified = false;
        this.isUpdateDeleted = true;
    }

    protected changeDeleteFinish() {
        const id: string = this.updateId;
        const accommodation: any = this.accommodations.find(a => a.id === id);
        if (accommodation) {
            const status: number = accommodation.is_deleted;

            // this.updateVerify(id, status);
            this.getAlert("Notification", "Update successfully");
        }

        this.updateId = '';
        this.isUpdateDeleted = false;
    }

    protected openDeleteConfirmModal() {
        this.isModalConfirmDeleteOpen = true;
    }

    protected closeDeleteConfirmModal() {
        this.isModalConfirmDeleteOpen = false;
        this.isUpdateDeleted = false;
    }

    ngAfterViewInit(): void {
        this.descEls.changes.subscribe(() => {
            setTimeout(() => this.checkDescriptionOverflow(), 0);
        });
    }
}
