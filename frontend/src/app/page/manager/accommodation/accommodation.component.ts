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
import { TuiInputModule } from '@taiga-ui/legacy';
import {
    TuiConfirmService,
    TuiFiles,
    TuiCheckbox,
    tuiCreateTimePeriods,
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

@Component({
    standalone: true,
    selector: 'app-accommodation',
    imports: [
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
        'Name',
        'Country',
        'City',
        'District',
        'Address',
        'Description',
        'Rating',
        'Google Map',
        'Image',
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
        address: new FormControl('', Validators.required),
        description: new FormControl('', Validators.required),
        googleMap: new FormControl('', Validators.required),
    });

    protected timePeriods = tuiCreateTimePeriods();

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
            googleMap: accommodation.google_map,
            address: accommodation.address,
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
            district: this.formAccommodation.get('district')?.value || '',
            address: this.formAccommodation.get('address')?.value || '',
            description: this.formAccommodation.get('description')?.value || '',
            google_map: this.formAccommodation.get('googleMap')?.value || '',
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
            district: this.formAccommodation.get('district')?.value || '',
            address: this.formAccommodation.get('address')?.value || '',
            description: this.formAccommodation.get('description')?.value || '',
            google_map: this.formAccommodation.get('googleMap')?.value || '',
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
