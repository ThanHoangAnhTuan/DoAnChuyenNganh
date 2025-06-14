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
    TuiIcon,
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
import { RouterLink } from '@angular/router';
import { TuiInputTimeModule } from '@taiga-ui/legacy';
import {
    CreateFacilityInput,
    Facility,
} from '../../../models/facility/facility.model';
import { FacilityService } from '../../../services/facility/facility.service';
import { NavbarComponent } from '../../../components/navbar/navbar.component';
import { AsyncPipe, NgIf } from '@angular/common';

import type { TuiFileLike } from '@taiga-ui/kit';
import { finalize, of, Subject, switchMap } from 'rxjs';
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
        TuiIcon,
        RouterLink,
        TuiInputTimeModule,
        TuiSelect,
        NavbarComponent,
        AsyncPipe,
        FormsModule,
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
    protected columns: string[] = ['Id', 'Name', 'Image', 'Action'];
    protected idFacilityUpdating = '';

    private readonly dialogs = inject(TuiDialogService);

    protected formFacility = new FormGroup({
        name: new FormControl('', Validators.required),
        image: new FormControl<TuiFileLike | null>(null, Validators.required),
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

        this.formFacility.patchValue({
            name: facility.name,
            image: null, // or set to a File object if available
        });

        // console.log('facility: ', facility);

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
        // console.log('name', this.formFacility.get('name')?.value);
        // console.log(this.control.value);
        // console.log('Thực hiện create facility...');
        const nameControl = this.formFacility.get('name');
        const imageControl = this.formFacility.get('image');
        // const nameValid = nameControl?.valid;
        const nameValue = nameControl?.value;
        const imageValue = imageControl?.value;

        // console.log('Name:', nameValue, 'valid?', nameValid);

        // Thay vì kiểm tra this.formFacility.invalid, ta sẽ chỉ đảm bảo có tên facility

        if (!nameValue) {
            alert('Vui lòng nhập tên facility');
            this.formFacility.markAllAsTouched();
            return;
        }

        if (!imageValue) {
            alert('Vui lòng chọn hình ảnh cho facility');
            this.formFacility.markAllAsTouched();
            return;
        }
        if (!nameValue) {
            alert('Vui lòng nhập tên facility');
            this.formFacility.markAllAsTouched();
            return;
        }

        const formData = new FormData();

        formData.append('name', nameValue);

        let image: File | null = null;

        if (this.control?.value instanceof File) {
            image = this.control.value;
        } else if (this.formFacility.get('image')?.value instanceof File) {
            image = this.formFacility.get('image')?.value as File;
        } else if (this.control?.value) {
            const tuiFile = this.control.value as any;

            if (tuiFile.file instanceof File) {
                image = tuiFile.file;
            }
        }

        // console.log(
        //     'File được tìm thấy:',
        //     image ? image.name : 'KHÔNG CÓ FILE'
        // );

        if (image) {
            formData.append('image', image, image.name);
            // console.log('Đã thêm file vào request:', image.name);
        } else {
            const fileName = this.control?.value?.name || '';
            formData.append('image', fileName);
            // console.log('Không có file, sử dụng filename:', fileName);
        }
        // formData.forEach((value, key) => {
        //     console.log(
        //         `FormData: ${key} = ${
        //             value instanceof File ? value.name : value
        //         }`
        //     );
        // });
        // console.log('Gửi request đến server...');
        this.facilityService.createFacility(formData).subscribe({
            next: (response) => {
                // console.log('Server trả về thành công:', response);

                // Cập nhật danh sách facility
                if (Array.isArray(response.data)) {
                    this.facilities.push(...response.data);
                } else if (response.data) {
                    this.facilities.push(response.data);
                }

                // Reset form và đóng dialog
                this.formFacility.reset();
                if (
                    this.control &&
                    this.control !== this.formFacility.get('image')
                ) {
                    this.control.reset();
                }
            },
            error: (error) => {
                console.error('Lỗi từ server:', error);
                alert(
                    `Lỗi khi tạo facility: ${error.message || 'Unknown error'}`
                );
            },
        });
    }

    protected updateFacility(): void {
        console.log('name', this.formFacility.get('name')?.value);
        // if (this.formFacility.invalid) {
        //     this.formFacility.markAllAsTouched();
        //     return;
        // }

        // // Show loading indicator

        // // Create FormData for multipart data submission
        // const formData = new FormData();
        // formData.append('id', this.idFacilityUpdating);
        // formData.append('name', this.formFacility.get('name')?.value || '');

        // // Handle image file if present
        // const imageFile = this.formFacility.get('image')?.value;
        // if (imageFile instanceof File) {
        //     formData.append('image', imageFile, imageFile.name);
        // }

        // this.facilityService.updateFacility(formData).subscribe({
        //     next: (response) => {
        //         // Update facility in the list
        //         const updatedFacility = response.data[0] as Facility;
        //         this.facilities = this.facilities.map((facility) => {
        //             return facility.id === updatedFacility.id
        //                 ? updatedFacility
        //                 : facility;
        //         });

        //         // Show success message
        //         this.showMessage('Cập nhật cơ sở thành công');

        //         // Reset the form or close the dialog if needed
        //         this.resetFormOrCloseDialog();
        //     },
        //     error: (error) => {
        //         console.error('Error updating facility:', error);
        //         this.showMessage(
        //             'Cập nhật cơ sở thất bại: ' +
        //                 (error.message || 'Đã xảy ra lỗi')
        //         );
        //     },
        //     complete: () => {
        //         // Hide loading indicator
        //         console.log('Update facility request completed');
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

    onChange(files: File[] | null): void {
        if (!files || files.length === 0) {
            return;
        }
        this.control.setValue(files[0]);
    }
}
