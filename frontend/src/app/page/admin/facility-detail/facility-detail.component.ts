import { Component, Inject, inject, OnInit } from '@angular/core';
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
import { TuiFiles, tuiCreateTimePeriods, TuiSelect } from '@taiga-ui/kit';
import { TuiCardLarge } from '@taiga-ui/layout';
import { RouterLink } from '@angular/router';
import { TuiInputTimeModule } from '@taiga-ui/legacy';
import {
    Facility,
    FacilityDetail,
} from '../../../models/facility/facility.model';
import { NavbarComponent } from '../../../components/navbar/navbar.component';
import { AsyncPipe, NgIf } from '@angular/common';

import type { Observable } from 'rxjs';
import { FacilityDetailService } from '../../../services/facility-detail/facility-detail.service';
import { TUI_DIALOGS_CLOSE } from '@taiga-ui/core';
import { MessageService } from 'primeng/api';
import { Toast } from 'primeng/toast';
import { ButtonModule } from 'primeng/button';
import { Ripple } from 'primeng/ripple';
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
            },
            error: (error) => {
                this.showToast(
                    'error',
                    'Tải dữ liệu',
                    `Không thể tải dữ liệu facility. Vui lòng thử lại sau`
                );
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
            this.showToast(
                'warn',
                'Facility Issue',
                'Vui lòng nhập tên facility'
            );
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
                this.showToast(
                    'success',
                    'Facility Created',
                    'Cơ sở đã được tạo thành công'
                );
            },
            error: (error) => {
                this.showToast(
                    'error',
                    'Facility Creation Error',
                    `Lỗi khi tạo cơ sở: ${error.message || 'Unknown error'}`
                );
            },
        });
    }

    protected updateFacility(): void {
        if (this.formFacility.invalid) {
            this.showToast(
                'warn',
                'Facility Update Issue',
                'Vui lòng nhập tên cơ sở'
            );
            this.formFacility.markAllAsTouched();
            return;
        }
        const data = {
            id: this.idFacilityUpdating,
            name: this.formFacility.get('name')?.value || '',
        };

        this.facilityService.updateFacility(data).subscribe({
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
                    'Facility Updated',
                    'Cơ sở đã được cập nhật thành công'
                );
            },
            error: (error) => {
                this.showToast(
                    'error',
                    'Facility Update Error',
                    `Lỗi khi cập nhật cơ sở: ${
                        error.message || 'Unknown error'
                    }`
                );
            },
            complete: () => {
                this.showToast(
                    'info',
                    'Update Facility Request Completed',
                    'Yêu cầu cập nhật cơ sở đã hoàn thành'
                );
            },
        });
    }
    protected deleteFacility(id: string) {
        this.facilityService.deleteFacility(id).subscribe((_) => {
            this.facilities = this.facilities.filter(
                (facility) => facility.id !== id
            );
            this.showToast(
                'success',
                'Facility Deleted',
                'Cơ sở đã được xóa thành công'
            );
        });
    }
}
