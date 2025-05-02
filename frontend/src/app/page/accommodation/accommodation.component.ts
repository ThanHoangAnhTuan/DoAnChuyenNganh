import { Component, inject, Injector, OnInit } from '@angular/core';
import { NgForOf, AsyncPipe, NgIf } from '@angular/common';
import {
    FormControl,
    FormGroup,
    FormsModule,
    ReactiveFormsModule,
    Validators,
} from '@angular/forms';
import { DomSanitizer, SafeHtml } from '@angular/platform-browser';

import { AccommodationService } from '../../services/accommodation.service';
import {
    Accommodation,
    CreateAccommodation,
    Facilities,
    PropertySurroundings,
    UpdateAccommodation,
} from '../../models/accommodation.model';

import {
    finalize,
    map,
    Observable,
    of,
    retry,
    Subject,
    switchMap,
    timer,
} from 'rxjs';

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
    TuiFileLike,
    TuiFiles,
    TuiCheckbox,
} from '@taiga-ui/kit';
import { TuiResponsiveDialogService } from '@taiga-ui/addon-mobile';
import { TuiCardLarge, TuiForm } from '@taiga-ui/layout';
import {
    TUI_EDITOR_DEFAULT_EXTENSIONS,
    TUI_EDITOR_DEFAULT_TOOLS,
    TUI_EDITOR_EXTENSIONS,
    TuiEditor,
} from '@taiga-ui/editor';

@Component({
    selector: 'app-accommodation',
    imports: [
        NgForOf,
        TuiTable,
        TuiIcon,
        TuiButton,
        TuiInputModule,
        FormsModule,
        ReactiveFormsModule,
        TuiTextfield,
        TuiAppearance,
        TuiForm,
        TuiCardLarge,
        AsyncPipe,
        NgIf,
        TuiFiles,
        TuiEditor,
        TuiGroup,
        TuiCheckbox,
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
    ];
    protected readonly baseUrl: string = 'http://localhost:8080/uploads/';
    protected value = '';
    readonly tools = TUI_EDITOR_DEFAULT_TOOLS;
    protected idAccommodationUpdating = '';

    private readonly confirm = inject(TuiConfirmService);
    private readonly dialogs = inject(TuiDialogService);

    protected readonly formCreate = new FormGroup({
        name: new FormControl(''),
        city: new FormControl(''),
        country: new FormControl(''),
        district: new FormControl(''),
        image: new FormControl<TuiFileLike | null>(null, Validators.required),
        description: new FormControl(''),
        wifi: new FormControl(false),
        airCondition: new FormControl(false),
        tv: new FormControl(false),
        googleMap: new FormControl(''),
        restaurant: new FormControl(false),
        bar: new FormControl(false),
        rules: new FormControl(''),
    });

    protected readonly formUpdate = new FormGroup({
        name: new FormControl(''),
        city: new FormControl(''),
        country: new FormControl(''),
        district: new FormControl(''),
        image: new FormControl<TuiFileLike | null>(null, Validators.required),
        description: new FormControl(''),
        rating: new FormControl(''),
        wifi: new FormControl(false),
        airCondition: new FormControl(false),
        tv: new FormControl(false),
        googleMap: new FormControl(''),
        restaurant: new FormControl(false),
        bar: new FormControl(false),
        rules: new FormControl(''),
    });

    protected readonly failedFiles$ = new Subject<TuiFileLike | null>();
    protected readonly loadingFiles$ = new Subject<TuiFileLike | null>();
    protected readonly loadedFiles$ = this.formCreate
        .get('image')!
        .valueChanges.pipe(switchMap((file) => this.processFile(file)));

    protected readonly previewUrl$ = this.loadedFiles$.pipe(
        map((file) => {
            if (!file) {
                return null;
            }

            const objectUrl = URL.createObjectURL(file as File);
            return this.sanitizer.bypassSecurityTrustUrl(objectUrl);
        })
    );

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
        this.dialogs
            .open(content, {
                label: 'Create Accommodation',
            })
            .subscribe({
                complete: () => {
                    this.formCreate.reset();
                },
            });
    }

    protected openDialogUpdate(
        content: PolymorpheusContent,
        accommodation: Accommodation
    ) {
        this.formUpdate.patchValue({
            name: accommodation.name,
            city: accommodation.city,
            country: accommodation.country,
            district: accommodation.district,
            description: accommodation.description,
            rating: accommodation.rating,
            wifi: accommodation.facilities.wifi,
            airCondition: accommodation.facilities.airCondition,
            tv: accommodation.facilities.tv,
            googleMap: accommodation.googleMap,
            restaurant: accommodation.propertySurrounds.restaurant,
            bar: accommodation.propertySurrounds.bar,
            rules: accommodation.rules,
        });

        this.idAccommodationUpdating = accommodation.id;

        this.dialogs
            .open(content, {
                label: 'Update Accommodation',
            })
            .subscribe({
                complete: () => {
                    this.formUpdate.reset();
                },
            });
    }

    protected getDescription(html: string): SafeHtml {
        return this.sanitizer.bypassSecurityTrustHtml(html);
    }

    protected removeFile(): void {
        this.formCreate.get('image')!.setValue(null);
    }

    protected processFile(
        file: TuiFileLike | null
    ): Observable<TuiFileLike | null> {
        this.failedFiles$.next(null);

        if (this.formCreate.get('image')!.invalid || !file) {
            return of(null);
        }

        this.loadingFiles$.next(file);

        return timer(1000).pipe(
            map(() => {
                return file;
            }),
            finalize(() => this.loadingFiles$.next(null))
        );
    }

    protected get imageControl(): FormControl {
        return this.formCreate.get('image') as FormControl;
    }

    async addDefaultImage(url: string): Promise<TuiFileLike> {
        const response = await fetch(url);
        const blob = await response.blob();
        const file = new File([blob], url, { type: blob.type });
        return {
            name: file.name,
            size: file.size,
            type: file.type,
            src: URL.createObjectURL(file),
        };
    }

    protected createAccommodation() {
        const facilities: Facilities = {
            airCondition: this.formCreate.get('airCondition')?.value || false,
            tv: this.formCreate.get('tv')?.value || false,
            wifi: this.formCreate.get('wifi')?.value || false,
        };

        const propertySurrounds: PropertySurroundings = {
            bar: this.formCreate.get('bar')?.value || false,
            restaurant: this.formCreate.get('restaurant')?.value || false,
        };

        const file = this.formCreate.get('image')?.value;

        const accommodation: CreateAccommodation = {
            name: this.formCreate.get('name')?.value || '',
            city: this.formCreate.get('city')?.value || '',
            country: this.formCreate.get('country')?.value || '',
            description: this.formCreate.get('description')?.value || '',
            district: this.formCreate.get('district')?.value || '',
            facilities: facilities,
            googleMap: this.formCreate.get('googleMap')?.value || '',
            image: file ? [file as File] : [],
            propertySurrounds: propertySurrounds,
            rules: this.formCreate.get('rules')?.value || '',
        };

        this.accommodationService
            .createAccommodation(accommodation)
            .subscribe((response) => {
                this.accommodations.push(response.data);
            });
    }

    protected updateAccommodation() {
        const facilities: Facilities = {
            airCondition: this.formUpdate.get('airCondition')?.value || false,
            tv: this.formUpdate.get('tv')?.value || false,
            wifi: this.formUpdate.get('wifi')?.value || false,
        };

        const propertySurrounds: PropertySurroundings = {
            bar: this.formUpdate.get('bar')?.value || false,
            restaurant: this.formUpdate.get('restaurant')?.value || false,
        };

        const file = this.formUpdate.get('image')?.value;

        const accommodation: UpdateAccommodation = {
            id: this.idAccommodationUpdating,
            name: this.formUpdate.get('name')?.value || '',
            city: this.formUpdate.get('city')?.value || '',
            country: this.formUpdate.get('country')?.value || '',
            description: this.formUpdate.get('description')?.value || '',
            district: this.formUpdate.get('district')?.value || '',
            facilities: facilities,
            googleMap: this.formUpdate.get('googleMap')?.value || '',
            image: file ? [file as File] : [],
            propertySurrounds: propertySurrounds,
            rules: this.formUpdate.get('rules')?.value || '',
        };

        console.log(accommodation);

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
        this.accommodationService
            .deleteAccommodation(id)
            .subscribe((response) => {
                this.accommodations = this.accommodations.filter(
                    (accommodation) => accommodation.id !== id
                );
            });
    }
}
