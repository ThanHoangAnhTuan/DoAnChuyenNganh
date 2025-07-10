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
    TuiButton,
    TuiDialogService,
    TuiTextfield,
    TuiAppearance,
} from '@taiga-ui/core';
import type { PolymorpheusContent } from '@taiga-ui/polymorpheus';
import { TuiInputModule } from '@taiga-ui/legacy';
import { TuiFiles, tuiCreateTimePeriods, TuiSelect } from '@taiga-ui/kit';
import { TuiCardLarge } from '@taiga-ui/layout';
import { TuiInputTimeModule } from '@taiga-ui/legacy';
import {
    Facility,
    FacilityDetail,
} from '../../../models/facility/facility.model';
import { NavbarComponent } from '../../../components/navbar/navbar.component';
import { finalize, type Observable } from 'rxjs';
import { FacilityDetailService } from '../../../services/facility-detail/facility-detail.service';
import { TUI_DIALOGS_CLOSE } from '@taiga-ui/core';
import { MessageService } from 'primeng/api';
import { Toast } from 'primeng/toast';
import { ButtonModule } from 'primeng/button';
import { LoaderComponent } from '../../../components/loader/loader.component';
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
        TuiInputTimeModule,
        TuiSelect,
        NavbarComponent,
        FormsModule,
        TuiFiles,
        Toast,
        ButtonModule,
        LoaderComponent,
    ],
    templateUrl: './facility-detail.component.html',
    styleUrl: './facility-detail.component.scss',
    providers: [MessageService],
})
export class FacilityDetailComponent implements OnInit {
    protected facilities!: FacilityDetail[];
    protected columns: string[] = ['Id', 'Name', 'Action'];
    protected idFacilityUpdating = '';
    isLoading: boolean = false;

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
        this.facilityService
            .getFacilities()
            .pipe(finalize(() => (this.isLoading = false)))
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
                        `Không thể tải dữ liệu facility. Vui lòng thử lại sau`,
                        `${error.message || ''}`
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
            this.showToast('warn', 'Cơ sở Bắt buộc', 'Vui lòng nhập tên cơ sở');
            this.formFacility.markAllAsTouched();
            return;
        }

        const data = {
            name: nameValue,
        };
        this.isLoading = true;
        this.facilityService
            .createFacility(data)
            .pipe(finalize(() => (this.isLoading = false)))
            .subscribe({
                next: (response) => {
                    this.facilities.push(response.data);
                    this.formFacility.reset();
                    this.close$.subscribe();
                    this.showToast(
                        'success',
                        'Cơ sở Đã Được Tạo Thành Công!',
                        ''
                    );
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
        if (this.formFacility.invalid) {
            this.showToast('warn', 'Cơ sở Bắt buộc', 'Vui lòng nhập tên cơ sở');
            this.formFacility.markAllAsTouched();
            return;
        }
        const data = {
            id: this.idFacilityUpdating,
            name: this.formFacility.get('name')?.value || '',
        };
        this.isLoading = true;
        this.facilityService
            .updateFacility(data)
            .pipe(finalize(() => (this.isLoading = false)))
            .subscribe({
                next: (response) => {
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
                },
                complete: () => {
                    this.showToast(
                        'info',
                        'Yêu cầu cập nhật cơ sở đã hoàn thành',
                        ''
                    );
                },
            });
    }
    protected deleteFacility(id: string) {
        // this.facilityService.deleteFacility(id).subscribe((_) => {
        //     this.facilities = this.facilities.filter(
        //         (facility) => facility.id !== id
        //     );
        //     this.showToast('success', 'Cơ sở đã được xóa thành công', '');
        // });
        this.isLoading = true;
        this.facilityService
            .deleteFacility(id)
            .pipe(finalize(() => (this.isLoading = false)))
            .subscribe({
                next: () => {
                    this.facilities = this.facilities.filter(
                        (facility) => facility.id !== id
                    );
                    this.showToast(
                        'success',
                        'Cơ sở đã được xóa thành công',
                        ''
                    );
                },
                error: (error) => {
                    this.showToast(
                        'error',
                        'Xóa Cơ Sở Thất Bại',
                        `${error.message || ''}`
                    );
                },
            });
    }
}
