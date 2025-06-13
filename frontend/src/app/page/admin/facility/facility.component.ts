import { CreateFacilityInput } from './../../../models/facility/facility.model';
import {
    ChangeDetectionStrategy,
    Component,
    inject,
    Injector,
    OnInit,
} from '@angular/core';
import {
    FormControl,
    FormGroup,
    FormsModule,
    ReactiveFormsModule,
    Validators,
} from '@angular/forms';
import { DomSanitizer, SafeHtml } from '@angular/platform-browser';

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
    TuiRating,
    TuiSelect,
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
import {
    Facility,
    UpdateFacility,
} from '../../../models/facility/facility.model';
import { FacilityService } from '../../../services/facility/facility.service';
import { NavbarComponent } from '../../../components/navbar/navbar.component';
import { AsyncPipe, NgForOf, NgIf } from '@angular/common';

import type { TuiFileLike } from '@taiga-ui/kit';
import { finalize, map, of, Subject, switchMap, timer } from 'rxjs';
import type { Observable } from 'rxjs';

@Component({
    selector: 'app-facility',
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
        RouterLink,
        TuiInputTimeModule,
        TuiSelect,
        NavbarComponent,
        AsyncPipe,
        FormsModule,
        NgForOf,
        TuiFiles,
        NgIf,
    ],
    templateUrl: './facility.component.html',
    styleUrl: './facility.component.scss',
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
export class FacilityComponent implements OnInit {
    protected facilities!: Facility[];
    protected columns: string[] = ['Id', 'Name', 'Image', 'Actions'];
    protected idFacilityUpdating = '';

    private readonly dialogs = inject(TuiDialogService);

    protected formFacility = new FormGroup({
        name: new FormControl('', Validators.required),
    });

    protected timePeriods = tuiCreateTimePeriods();

    constructor(private facilityService: FacilityService) {}

    ngOnInit() {
        this.facilityService.getFacilities().subscribe({
            next: (response) => {
                this.facilities = response.data;
                // this.createFacilityControls();
                // console.log('Facilities:', this.facilities);
            },
            error: (error) => {
                console.error('Error fetching facilities:', error);
            },
        });
    }

    protected openDialogCreate(content: PolymorpheusContent): void {
        this.formFacility.reset();
        this.dialogs
            .open(content, {
                label: 'Create Facility',
            })
            .subscribe({
                complete: () => {
                    this.formFacility.reset();
                },
            });
    }

    protected openDialogUpdate(
        content: PolymorpheusContent,
        facility: Facility
    ) {
        this.formFacility.reset();
        // this.formFacilities.reset();

        this.formFacility.patchValue({
            name: facility.name,
            // image: null, // or set to a File object if available
        });

        console.log('facility: ', facility);

        this.idFacilityUpdating = facility.id;

        this.dialogs
            .open(content, {
                label: 'Update Facility',
            })
            .subscribe({
                complete: () => {
                    this.formFacility.reset();
                },
            });
    }

    protected CreateFacilityInput(): void {
        // if (this.formFacility.invalid) {
        //     this.formFacility.markAllAsTouched();
        //     console.log('Form invalid:', this.formFacility.errors);
        //     console.log('Image value:', this.formFacility.get('image')?.value);
        //     return;
        // }
        console.log('name', this.formFacility.get('name')?.value);
        // console.log(this.loadedFiles$);
        console.log(this.control.value);

        // Tạo FormData để gửi dữ liệu multipart (bao gồm file)
        // const formData = new FormData();
        // formData.append('name', this.formFacility.get('name')?.value || '');

        // // Lấy file từ form control
        // const imageFile = this.formFacility.get('image')?.value;
        // if (imageFile instanceof File) {
        //     const isSvg =
        //         imageFile.type === 'image/svg+xml' ||
        //         imageFile.type === 'image/svg';

        //     console.log(
        //         'Uploading SVG file:',
        //         imageFile.name,
        //         'Type:',
        //         imageFile.type
        //     );
        //     formData.append('image', imageFile, imageFile.name);
        // } else {
        //     console.error('Image is not a File object:', imageFile);
        //     alert('Vui lòng chọn một hình ảnh');
        //     return;
        // }

        // this.facilityService.createFacility(formData).subscribe({
        //     next: (response) => {
        //         this.facilities.push(...response.data);
        //         // this.createFacilityControls();
        //     },
        //     error: (error) => {
        //         console.error('Error creating facility:', error);
        //     },
        // });
    }

    protected updateFacility(): void {
        console.log('name', this.formFacility.get('name')?.value);
        // if (this.formFacility.invalid) {
        //     this.formFacility.markAllAsTouched();
        //     return;
        // }

        // Tạo FormData để gửi dữ liệu multipart (bao gồm file)
        // const formData = new FormData();
        // formData.append('id', this.idFacilityUpdating);
        // formData.append('name', this.formFacility.get('name')?.value || '');

        // // Lấy file từ form control
        // const imageFile = this.formFacility.get('image')?.value;
        // if (imageFile instanceof File) {
        //     formData.append('image', imageFile, imageFile.name);
        // }

        // this.facilityService.updateFacility(formData).subscribe({
        //     next: (response) => {
        //         // Cập nhật facility trong danh sách
        //         const updatedFacility = response.data[0] as Facility;
        //         this.facilities = this.facilities.map((facility) => {
        //             if (facility.id === updatedFacility.id) {
        //                 return updatedFacility;
        //             } else {
        //                 return facility;
        //             }
        //         });
        //     },
        //     error: (error) => {
        //         console.error('Error updating facility:', error);
        //     },
        // });
    }

    protected deleteFacility(id: string) {
        this.facilityService.deleteFacility(id).subscribe((_) => {
            this.facilities = this.facilities.filter(
                (facility) => facility.id !== id
            );
        });
    }

    protected readonly control = new FormControl<TuiFileLike | null>(
        null,
        Validators.required
    );

    protected readonly failedFiles$ = new Subject<TuiFileLike | null>();
    protected readonly loadingFiles$ = new Subject<TuiFileLike | null>();
    protected readonly loadedFiles$ = this.control.valueChanges.pipe(
        switchMap((file) => this.processFile(file))
    );

    protected removeFile(): void {
        this.control.setValue(null);
    }

    protected processFile(
        file: TuiFileLike | null
    ): Observable<TuiFileLike | null> {
        this.failedFiles$.next(null);

        if (this.control.invalid || !file) {
            return of(null);
        }

        this.loadingFiles$.next(file);

        return of(file).pipe(finalize(() => this.loadingFiles$.next(null)));
    }
}
