import { Component, inject, Injector, OnInit } from '@angular/core';
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
} from '../../../models/accommodation.model';

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
import { TuiInputModule } from '@taiga-ui/legacy';
import { TuiConfirmService, TuiFiles, TuiCheckbox } from '@taiga-ui/kit';
import { TuiResponsiveDialogService } from '@taiga-ui/addon-mobile';
import { TuiCardLarge } from '@taiga-ui/layout';
import {
    TUI_EDITOR_DEFAULT_EXTENSIONS,
    TUI_EDITOR_DEFAULT_TOOLS,
    TUI_EDITOR_EXTENSIONS,
    TuiEditor,
} from '@taiga-ui/editor';
import { RouterLink } from '@angular/router';

@Component({
    selector: 'app-accommodation',
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
        TuiFiles,
        TuiEditor,
        TuiGroup,
        TuiCheckbox,
        RouterLink,
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
export class AccommodationComponent implements OnInit {
    protected accommodations!: Accommodation[];
    protected columns: string[] = [
        'ID',
        'Manager ID',
        'Name',
        'City',
        'Country',
        'District',
        'Image',
        'Description',
        'Rating',
        'Google Map',
        'Rules',
        'Wifi',
        'Air Condition',
        'TV',
        'Restaurant',
        'Bar',
        'Action',
        'Show Accommodation Detail',
    ];
    protected readonly tools = TUI_EDITOR_DEFAULT_TOOLS;
    protected idAccommodationUpdating = '';

    private readonly dialogs = inject(TuiDialogService);

    protected formAccommodation = new FormGroup({
        name: new FormControl('', Validators.required),
        city: new FormControl('', Validators.required),
        country: new FormControl('', Validators.required),
        district: new FormControl('', Validators.required),
        description: new FormControl('', Validators.required),
        wifi: new FormControl(false),
        airCondition: new FormControl(false),
        tv: new FormControl(false),
        googleMap: new FormControl('', Validators.required),
        restaurant: new FormControl(false),
        bar: new FormControl(false),
        rules: new FormControl('', Validators.required),
    });

    constructor(
        private accommodationService: AccommodationService,
        private sanitizer: DomSanitizer
    ) {}

    ngOnInit() {
        this.accommodationService.getAccommodations().subscribe((response) => {
            this.accommodations = response.data;
        });
    }

    protected openDialogCreate(content: PolymorpheusContent): void {
        this.formAccommodation.reset();
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

        this.formAccommodation.patchValue({
            name: accommodation.name,
            city: accommodation.city,
            country: accommodation.country,
            district: accommodation.district,
            description: accommodation.description,
            wifi: accommodation.facilities.wifi,
            airCondition: accommodation.facilities.air_condition,
            tv: accommodation.facilities.tv,
            googleMap: accommodation.google_map,
            restaurant: accommodation.property_surrounds.restaurant,
            bar: accommodation.property_surrounds.bar,
            rules: accommodation.rules,
        });

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

    protected getDescription(html: string): SafeHtml {
        return this.sanitizer.bypassSecurityTrustHtml(html);
    }

    protected createAccommodation() {
        const accommodation: CreateAccommodation = {
            name: this.formAccommodation.get('name')?.value || '',
            city: this.formAccommodation.get('city')?.value || '',
            country: this.formAccommodation.get('country')?.value || '',
            description: this.formAccommodation.get('description')?.value || '',
            district: this.formAccommodation.get('district')?.value || '',
            facilities: {
                air_condition:
                    this.formAccommodation.get('airCondition')?.value || false,
                tv: this.formAccommodation.get('tv')?.value || false,
                wifi: this.formAccommodation.get('wifi')?.value || false,
            },
            google_map: this.formAccommodation.get('googleMap')?.value || '',
            property_surrounds: {
                bar: this.formAccommodation.get('bar')?.value || false,
                restaurant:
                    this.formAccommodation.get('restaurant')?.value || false,
            },
            rules: this.formAccommodation.get('rules')?.value || '',
        };

        if (this.formAccommodation.invalid) {
            this.formAccommodation.markAllAsTouched();
            return;
        }

        this.accommodationService
            .createAccommodation(accommodation)
            .subscribe((response) => {
                this.accommodations.push(response.data);
            });
    }

    protected updateAccommodation() {
        const accommodation: UpdateAccommodation = {
            id: this.idAccommodationUpdating,
            name: this.formAccommodation.get('name')?.value || '',
            city: this.formAccommodation.get('city')?.value || '',
            country: this.formAccommodation.get('country')?.value || '',
            description: this.formAccommodation.get('description')?.value || '',
            district: this.formAccommodation.get('district')?.value || '',
            facilities: {
                air_condition:
                    this.formAccommodation.get('airCondition')?.value || false,
                tv: this.formAccommodation.get('tv')?.value || false,
                wifi: this.formAccommodation.get('wifi')?.value || false,
            },
            google_map: this.formAccommodation.get('googleMap')?.value || '',
            property_surrounds: {
                bar: this.formAccommodation.get('bar')?.value || false,
                restaurant:
                    this.formAccommodation.get('restaurant')?.value || false,
            },
            rules: this.formAccommodation.get('rules')?.value || '',
        };

        this.accommodationService
            .updateAccommodation(accommodation)
            .subscribe((response) => {
                this.accommodations = this.accommodations.map(
                    (accommodation) => {
                        if (accommodation.id === response.data.id) {
                            return response.data;
                        } else {
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
        });
    }
}
