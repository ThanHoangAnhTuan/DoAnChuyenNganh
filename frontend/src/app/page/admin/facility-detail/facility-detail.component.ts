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
    ],
    templateUrl: './facility-detail.component.html',
    styleUrl: './facility-detail.component.scss',
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
        private facilityService: FacilityDetailService
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
            alert('Vui lòng nhập tên facility');
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
        if (this.formFacility.invalid) {
            this.formFacility.markAllAsTouched();
            return;
        }
        const data = {
            id: this.idFacilityUpdating,
            name: this.formFacility.get('name')?.value || '',
        };

        this.facilityService.updateFacility(data).subscribe({
            next: (response) => {
                console.log(response);
                const updatedFacility = response.data as Facility;
                this.facilities = this.facilities.map((facility) => {
                    return facility.id === updatedFacility.id
                        ? updatedFacility
                        : facility;
                });

                // Show success message
                console.log('Cập nhật cơ sở thành công');
            },
            error: (error) => {
                console.error('Error updating facility:', error);
                console.warn(
                    'Cập nhật cơ sở thất bại: ' +
                        (error.message || 'Đã xảy ra lỗi')
                );
            },
            complete: () => {
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
        });
    }
}
