import { Component, Inject, inject, Injector, OnInit } from '@angular/core';
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
    Facility,
    FacilityDetail,
} from '../../../models/facility/facility.model';
import { FacilityService } from '../../../services/facility/facility.service';
import { NavbarComponent } from '../../../components/navbar/navbar.component';
import { AsyncPipe, NgIf } from '@angular/common';

import type { TuiFileLike } from '@taiga-ui/kit';
import { finalize, of, Subject, switchMap } from 'rxjs';
import type { Observable } from 'rxjs';
import { FacilityDetailService } from '../../../services/facility-detail/facility-detail.service';
import { TUI_DIALOGS_CLOSE } from '@taiga-ui/core';
import { MessageService } from 'primeng/api';
import { Toast } from 'primeng/toast';
import { ButtonModule } from 'primeng/button';
import { Ripple } from 'primeng/ripple';
// import { TUI_DIALOGS_CLOSE } from '@taiga-ui/cdk';
@Component({
    selector: 'app-facility-detail',
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
        Toast,
        ButtonModule,
        Ripple,
    ],
    templateUrl: './facility-detail.component.html',
    styleUrl: './facility-detail.component.scss',
    providers: [MessageService],
})
export class FacilityDetailComponent implements OnInit {
    protected facilities!: FacilityDetail[];
    protected columns: string[] = ['Id', 'Name', 'Action'];
    protected idFacilityUpdating = '';

    private readonly dialogs = inject(TuiDialogService);

    protected formFacility = new FormGroup({
        name: new FormControl('', Validators.required),
    });

    protected timePeriods = tuiCreateTimePeriods();

    constructor(
        @Inject(TUI_DIALOGS_CLOSE) private readonly close$: Observable<void>,
        private facilityService: FacilityDetailService,
        private messageService: MessageService
    ) {}

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
    showToast(
        severity: 'success' | 'info' | 'warn' | 'error',
        summary: string,
        detail: string
    ): void {
        this.messageService.add({ severity, summary, detail });
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
        facility: FacilityDetail
    ) {
        this.formFacility.reset();

        this.formFacility.patchValue({
            name: facility.name,
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
        const nameValue = this.formFacility.get('name')?.value;

        if (!nameValue) {
            this.showToast('warn', 'Cơ sở Bắt buộc', 'Vui lòng nhập tên cơ sở');
            this.formFacility.markAllAsTouched();
            return;
        }

        const data = {
            name: nameValue,
        };

        this.facilityService.createFacility(data).subscribe({
            next: (response) => {
                if (Array.isArray(response.data)) {
                    this.facilities.push(...response.data);
                } else if (response.data) {
                    this.facilities.push(response.data);
                }

                this.formFacility.reset();
                this.close$.subscribe();
                this.showToast('success', 'Cơ sở Đã Được Tạo Thành Công!', '');
            },
            error: (error) => {
                this.showToast(
                    'error',
                    'Lỗi khi tạo cơ sở',
                    `${error.message || ''}`
                );
            },
        });
    }

    protected updateFacility(): void {
        // console.log('name', this.formFacility.get('name')?.value);
        if (this.formFacility.invalid) {
            this.showToast('warn', 'Cơ sở Bắt buộc', 'Vui lòng nhập tên cơ sở');
            this.formFacility.markAllAsTouched();
            return;
        }
        const data = {
            id: this.idFacilityUpdating,
            name: this.formFacility.get('name')?.value || '',
        };

        this.facilityService.updateFacility(data).subscribe({
            next: (response) => {
                // console.log(response);
                const updatedFacility = response.data as Facility;
                this.facilities = this.facilities.map((facility) => {
                    return facility.id === updatedFacility.id
                        ? updatedFacility
                        : facility;
                });

                // Show success message
                this.showToast('success', 'Cập nhật Cơ Sở Thành Công', '');
                // console.log('Cập nhật cơ sở thành công');
            },
            error: (error) => {
                this.showToast(
                    'error',
                    'Cập nhật Cơ Sở Thất Bại',
                    `${error.message || ''}`
                );
                // console.error('Error updating facility:', error);
                // console.warn(
                //     'Cập nhật cơ sở thất bại: ' +
                //         (error.message || 'Đã xảy ra lỗi')
                // );
            },
            complete: () => {
                this.showToast(
                    'info',
                    'Yêu cầu cập nhật cơ sở đã hoàn thành',
                    ''
                );
                // Hide loading indicator
                console.log('Update facility request completed');
            },
        });
    }
    protected deleteFacility(id: string) {
        this.facilityService.deleteFacility(id).subscribe((_) => {
            this.facilities = this.facilities.filter(
                (facility) => facility.id !== id
            );
            this.showToast('success', 'Cơ sở đã được xóa thành công', '');
        });
    }
}
