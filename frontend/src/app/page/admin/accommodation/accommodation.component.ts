import {
    AfterViewInit,
    Component,
    ElementRef,
    inject,
    Injector,
    OnInit,
    QueryList,
    ViewChild,
    ViewChildren,
} from '@angular/core';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import {
    DomSanitizer,
    SafeHtml,
    SafeResourceUrl,
} from '@angular/platform-browser';
import { TuiTable } from '@taiga-ui/addon-table';
import {
    TuiButton,
    TuiDialogService,
    TuiTextfield,
    TuiAppearance,
    TuiAlertService,
} from '@taiga-ui/core';
import { TuiInputModule, TuiSelectModule } from '@taiga-ui/legacy';
import {
    TuiConfirmService,
    TuiFiles,
    tuiCreateTimePeriods,
    TuiSelect,
    TuiDataListWrapper,
    TuiPagination,
} from '@taiga-ui/kit';
import { TuiResponsiveDialogService } from '@taiga-ui/addon-mobile';
import {
    TUI_EDITOR_DEFAULT_EXTENSIONS,
    TUI_EDITOR_DEFAULT_TOOLS,
    TUI_EDITOR_EXTENSIONS,
} from '@taiga-ui/editor';
import { ActivatedRoute, RouterLink } from '@angular/router';
import { TuiInputTimeModule } from '@taiga-ui/legacy';
import { AddressService } from '../../../services/address/address.service';
import { City, District } from '../../../models/address/address.model';
import { NavbarComponent } from '../../../components/navbar/navbar.component';
import { ManagerService } from '../../../services/admin/manager.service';
import {
    GetAccommodationsOfManagerByAdmin,
    SetDeletedAccommodationInput,
    VerifyAccommodationInput,
} from '../../../models/admin/manager.model';
import { Toast } from 'primeng/toast';
import { MessageService } from 'primeng/api';
import { finalize } from 'rxjs';
import { LoaderComponent } from '../../../components/loader/loader.component';
import { Pagination } from '../../../models/pagination/pagination.model';


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
        TuiFiles,
        TuiInputTimeModule,
        TuiSelect,
        TuiSelectModule,
        TuiDataListWrapper,
        NavbarComponent,
        Toast,
        LoaderComponent,
        TuiPagination,
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
export class AccommodationComponent implements OnInit, AfterViewInit {
    @ViewChildren('descEl') descEls!: QueryList<ElementRef<HTMLDivElement>>;
    @ViewChild('topList') topList!: ElementRef;

    protected columns: string[] = [
        'ID',
        'Name',
        'Country',
        'City',
        'District',
        'Address',
        'Description',
        'Rating',
        'Google Map',
        // 'Image',
        'Is Verified',
        'Is Deleted',
    ];
    protected readonly tools = TUI_EDITOR_DEFAULT_TOOLS;
    protected managerId: string = '';
    protected accommodations: GetAccommodationsOfManagerByAdmin[] = [];
    protected pagination: Pagination = {
        page: 1,
        limit: 10,
        total: 0,
        total_pages: 0,
    };
    protected cities: City[] = [];
    protected districts: District[] = [];
    protected cityNames: string[] = [];
    protected districtNames: string[] = [];
    protected cityName: string = '';
    protected citySlug: string = '';
    protected districtName: string = '';
    protected districtSlug: string = '';
    protected showFullMap: { [id: string]: boolean } = {};
    protected elList: { [id: string]: any } = {};
    protected showButtonStates: { [id: string]: boolean } = {};
    protected isUpdateVerified: boolean = false;
    protected isUpdateDeleted: boolean = false;
    protected updateId: string = '';
    protected isModalConfirmVerifyOpen: boolean = false;
    protected isModalConfirmDeleteOpen: boolean = false;
    isLoading: boolean = false;

    protected timePeriods = tuiCreateTimePeriods();
    private readonly alerts = inject(TuiAlertService);

    getSafeUrl(url: string): SafeResourceUrl {
        return this.sanitizer.bypassSecurityTrustResourceUrl(url);
    }

    showToast(
        severity: 'success' | 'info' | 'warn' | 'error',
        summary: string,
        detail: string
    ): void {
        this.messageService.add({ severity, summary, detail });
    }

    constructor(
        private accommodationService: ManagerService,
        private addressService: AddressService,
        private sanitizer: DomSanitizer,
        private route: ActivatedRoute,
        private messageService: MessageService
    ) {}

    ngOnInit() {
        // this.route.params.subscribe((params) => {
        //     this.managerId = params['id'];
        //     this.accommodationService
        //         .getAccommodationsOfManagerByAdmin(this.managerId)
        //         .subscribe((response) => {
        //             console.log(response);
        //             this.accommodations = response.data;
        //         });
        // });

        this.isLoading = true;

        this.route.params.subscribe((params) => {
            this.managerId = params['id'];
            this.accommodationService
                .getAccommodationsOfManagerByAdmin(this.managerId)
                .pipe(
                    finalize(() => {
                        this.isLoading = false;
                    })
                )
                .subscribe({
                    next: (response) => {
                        console.log(response);
                        this.accommodations = response.data;
                    },
                    error: (error) => {
                        console.error('Error loading accommodations:', error);
                        this.showToast(
                            'error',
                            'Lỗi tải dữ liệu',
                            'Không thể tải danh sách khách sạn. Vui lòng thử lại sau.'
                        );
                    },
                });
        });
        // this.addressService.getCities().subscribe({
        //     next: (res) => {
        //         this.cities = res.data;
        //         this.cityNames = res.data.map((city) => city.name);
        //     },
        //     error: (err) => {
        //         console.error('Error fetching cities:', err);
        //     },
        // });
        this.addressService
            .getCities()
            .pipe(
                finalize(() => {
                    this.isLoading = false;
                })
            )
            .subscribe({
                next: (res) => {
                    this.cities = res.data;
                    this.cityNames = this.cities.map((city) => city.name);
                },
                error: (err) => {
                    console.error('Error fetching cities:', err);
                    this.showToast(
                        'error',
                        'Lỗi tải dữ liệu',
                        'Không thể tải danh sách thành phố. Vui lòng thử lại sau.'
                    );
                },
            });
    }

    private updateVerify(id: string, status: number) {
        this.isLoading = true;

        let newVerify: VerifyAccommodationInput = {
            accommodation_id: id,
            status: status,
        };
        this.accommodationService
            .updateVerified(newVerify)
            .pipe(
                finalize(() => {
                    this.isLoading = false;
                })
            )
            .subscribe({
                next: (response) => {
                    this.accommodationService
                        .getAccommodationsOfManagerByAdmin(this.managerId)
                        .subscribe((res) => {
                            this.accommodations = res.data;
                            this.showToast(
                                'success',
                                'Thành công',
                                `Đã cập nhật trạng thái xác minh thành công`
                            );
                        });
                },
                error: (error) => {
                    this.showToast(
                        'error',
                        'Lỗi',
                        `Không thể cập nhật trạng thái xác minh: ${error.message}`
                    );
                },
            });
    }

    private updateDelete(id: string, status: boolean) {
        let newDelete: SetDeletedAccommodationInput = {
            accommodation_id: id,
            status: status,
        };
        this.accommodationService
            .updateDeleted(newDelete)
            .subscribe((response) => {
                const message = response.message;
                this.getAlert('Notification', message);
            });
    }

    changeCitySlugToName(slug: string): string {
        const city = this.cities.find((city) => city.slug === slug);
        return city?.name ?? '';
    }

    changeDistrictSlugToName(citySlug: string, districtSlug: string): string {
        const city = this.cities.find((city) => city.slug === citySlug);
        let districts = city?.level2s ?? [];
        let district = districts.find(
            (district) => district.slug === districtSlug
        );
        return district?.name ?? '';
    }

    protected getDescription(html: string): SafeHtml {
        return this.sanitizer.bypassSecurityTrustHtml(html);
    }

    protected toggleDescription(id: string): void {
        this.showFullMap[id] = !this.showFullMap[id];
    }

    protected isDescriptionShown(id: string): boolean {
        return !!this.showFullMap[id];
    }

    private checkDescriptionOverflow() {
        setTimeout(() => {
            this.descEls.forEach((elRef) => {
                const el = elRef.nativeElement;
                const id = el.getAttribute('data-id');
                if (id) {
                    this.showButtonStates[id] = el.scrollHeight > 60;
                }
            });
        });
    }

    protected changeVerifiedApply(id: string) {
        this.updateId = id;
        this.isUpdateVerified = true;
        this.isUpdateDeleted = false;
    }

    protected changeVerifiedFinish() {
        const id: string = this.updateId;
        const accommodation = this.accommodations.find((a) => a.id === id);
        if (accommodation) {
            this.updateVerify(id, accommodation.is_verified);
        }
        this.updateId = '';
        this.isUpdateVerified = false;
    }

    protected resetInput() {
        this.isUpdateVerified = false;
        this.isUpdateDeleted = false;
    }

    protected openVerifyConfirmModal() {
        this.isModalConfirmVerifyOpen = true;
    }

    protected closeVerifyConfirmModal() {
        this.isModalConfirmVerifyOpen = false;
        this.isUpdateVerified = false;
    }

    protected changeDeletedApply(id: string) {
        this.updateId = id;
        this.isUpdateVerified = false;
        this.isUpdateDeleted = true;
    }

    protected changeDeleteFinish() {
        const id: string = this.updateId;
        const accommodation = this.accommodations.find((a) => a.id === id);
        if (accommodation) {
            this.showToast('success', 'Thành công', 'Đã xóa thành công');
            this.updateDelete(id, accommodation.is_deleted);
        }
        this.updateId = '';
        this.isUpdateDeleted = false;
    }

    protected openDeleteConfirmModal() {
        this.isModalConfirmDeleteOpen = true;
    }

    protected closeDeleteConfirmModal() {
        this.isModalConfirmDeleteOpen = false;
        this.isUpdateDeleted = false;
    }

    protected onPageChange(page: number) {
        console.log('Page changed to:', page + 1);

        this.accommodationService.getAccommodationsOfManagerByAdminWithPage(this.managerId, page + 1).subscribe((response) => {
            this.accommodations = response.data;
            this.pagination = response.pagination;
            this.pagination.page = page;
            this.scrollToTop();

            console.log(this.accommodations);
            console.log(this.pagination);
        })
    }

    protected scrollToTop() {
        if (this.topList) {
            this.topList.nativeElement.scrollIntoView({ behavior: 'smooth' });
        }
    }

    ngAfterViewInit(): void {
        this.descEls.changes.subscribe(() => {
            setTimeout(() => this.checkDescriptionOverflow(), 0);
        });
    }
}
