import { Component, inject, Injector, OnInit } from '@angular/core';
import {
    FormControl,
    FormGroup,
    FormsModule,
    ReactiveFormsModule,
    Validators,
} from '@angular/forms';
import { TuiTable } from '@taiga-ui/addon-table';
import {
    TuiButton,
    TuiDialogService,
    TuiTextfield,
    TuiAppearance,
} from '@taiga-ui/core';
import type { PolymorpheusContent } from '@taiga-ui/polymorpheus';
import { TuiInputModule } from '@taiga-ui/legacy';
import {
    TuiConfirmService,
    TuiFiles,
    tuiCreateTimePeriods,
    TuiSelect,
} from '@taiga-ui/kit';
import { TuiResponsiveDialogService } from '@taiga-ui/addon-mobile';
import { TuiCardLarge } from '@taiga-ui/layout';
import {
    TUI_EDITOR_DEFAULT_EXTENSIONS,
    TUI_EDITOR_EXTENSIONS,
} from '@taiga-ui/editor';
import { TuiInputTimeModule } from '@taiga-ui/legacy';
import { Facility } from '../../../models/facility/facility.model';
import { FacilityService } from '../../../services/facility/facility.service';
import { NavbarComponent } from '../../../components/navbar/navbar.component';
import { AsyncPipe, NgIf } from '@angular/common';
import type { TuiFileLike } from '@taiga-ui/kit';
import { finalize, of, Subject, switchMap } from 'rxjs';
import type { Observable } from 'rxjs';
import { MessageService } from 'primeng/api';
import { Toast } from 'primeng/toast';
import { ButtonModule } from 'primeng/button';
import { LoaderComponent } from "../../../components/loader/loader.component";

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
    TuiInputTimeModule,
    TuiSelect,
    NavbarComponent,
    AsyncPipe,
    FormsModule,
    TuiFiles,
    NgIf,
    Toast,
    ButtonModule,
    LoaderComponent
],
    templateUrl: './facility.component.html',
    styleUrl: './facility.component.scss',
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
export class FacilityComponent implements OnInit {
    protected facilities: Facility[] = [];
    protected columns: string[] = ['Id', 'Name', 'Image', 'Action'];
    protected idFacilityUpdating = '';
    isLoading: boolean = false;

    private readonly dialogs = inject(TuiDialogService);

    protected formFacility = new FormGroup({
        name: new FormControl('', Validators.required),
        image: new FormControl<TuiFileLike | null>(null),
    });

    protected timePeriods = tuiCreateTimePeriods();

    constructor(
        private facilityService: FacilityService,
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
        this.isLoading = true;
        this.facilityService
            .getFacilities()
            .pipe(
                finalize(() => {
                    this.isLoading = false;
                })
            )
            .subscribe({
                next: (response) => {
                    if (response.data) {
                        this.facilities = response.data;
                    } else {
                        this.facilities = [];
                    }
                },
                error: (error) => {
                    this.showToast(
                        'error',
                        'Lỗi khi lấy danh sách cơ sở:',
                        error.error.message || 'Unknown error'
                    );
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

        this.formFacility.patchValue({
            name: facility.name,
            image: null,
        });

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
        const nameValue = this.formFacility.get('name')?.value;
        const imageControlValue = this.formFacility.value.image;

        if (!nameValue) {
            this.showToast('warn', 'Cơ sở Bắt buộc', 'Vui lòng nhập tên cơ sở');
            this.formFacility.markAllAsTouched();
            return;
        }

        if (!imageControlValue || !(imageControlValue instanceof File)) {
            this.showToast(
                'warn',
                'Cơ sở Bắt buộc',
                'Vui lòng chọn hình ảnh cho cơ sở'
            );
            this.formFacility.markAllAsTouched();
            return;
        }

        this.isLoading = true;

        const formData = new FormData();

        formData.append('name', nameValue);
        formData.append('image', imageControlValue);

        this.facilityService
            .createFacility(formData)
            .pipe(
                finalize(() => {
                    this.isLoading = false;
                })
            )
            .subscribe({
                next: (response) => {
                    this.facilities.push(response.data);
                    this.formFacility.reset();
                    if (
                        this.control &&
                        this.control !== this.formFacility.get('image')
                    ) {
                        this.control.reset();
                    }
                    this.showToast(
                        'success',
                        'Cơ sở Đã Được Tạo Thành Công!',
                        ''
                    );
                },
                error: (error) => {
                    this.showToast(
                        'error',
                        'Lỗi khi tạo cơ sở:',
                        `${error.message || ''}`
                    );
                },
            });
    }

    protected updateFacility(): void {
        // Create form data for the update
        const formData = new FormData();
        formData.append('id', this.idFacilityUpdating);

        // Get the name value
        const nameValue = this.formFacility.get('name')?.value || '';

        // Validate name field only
        if (!nameValue.trim()) {
            // Show error message if name is empty
            this.formFacility.get('name')?.markAsTouched();
            return;
        }

        // Add name to form data
        formData.append('name', nameValue);

        // Only add image if a new one was selected
        const imageFile = this.formFacility.get('image')?.value;

        // SOLUTION 1: If no new image, fetch the current image file and send it
        if (imageFile instanceof File) {
            formData.append('image', imageFile, imageFile.name);
        }
        this.isLoading = true;
        this.facilityService
            .updateFacility(formData)
            .pipe(
                finalize(() => {
                    this.isLoading = false;
                })
            )
            .subscribe({
                next: (response) => {
                    const updatedFacility = response.data as Facility;
                    this.facilities = this.facilities.map((facility) => {
                        return facility.id === updatedFacility.id
                            ? updatedFacility
                            : facility;
                    });

                    // Show success message
                    this.showToast(
                        'success',
                        'Cập nhật Cơ Sở Thành Công',
                        'Cơ sở đã được cập nhật thành công'
                    );
                },
                error: (error) => {
                    // Handle error
                    this.showToast(
                        'error',
                        'Cập nhật Cơ Sở Thất Bại',
                        `Lỗi khi cập nhật cơ sở: ${error.message || ''}`
                    );
                },
            });
    }

    protected getCurrentFacilityImageUrl(): string {
        // Find the current facility being updated in the facilities array
        const currentFacility = this.facilities.find(
            (facility) => facility.id === this.idFacilityUpdating
        );

        // If facility is found and has an image, return the complete URL
        if (currentFacility && currentFacility.image) {
            return `http://localhost:8080/uploads/${currentFacility.image}`;
        }
        return 'assets/images/placeholder.png';
    }
    protected deleteFacility(id: string) {
        this.isLoading = true;
        this.facilityService
            .deleteFacility(id)
            .pipe(
                finalize(() => {
                    this.isLoading = false;
                })
            )
            .subscribe({
                next: () => {
                    this.facilities = this.facilities.filter(
                        (facility) => facility.id !== id
                    );
                },
                error: (error) => {
                    this.showToast(
                        'error',
                        'Lỗi khi xóa cơ sở:',
                        error.error.message || 'Unknown error'
                    );
                },
            });
        this.showToast('success', 'Cơ sở đã được xóa thành công', '');
    }
    protected get control(): FormControl<TuiFileLike | null> {
        return this.formFacility.get(
            'image'
        ) as FormControl<TuiFileLike | null>;
    }

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

    onChange(files: File[] | null): void {
        if (!files || files.length === 0) {
            return;
        }
        this.control.setValue(files[0]);
    }
}
