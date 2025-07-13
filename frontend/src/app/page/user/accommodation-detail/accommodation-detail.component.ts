import {
    Component,
    ElementRef,
    HostListener,
    inject,
    OnInit,
    ViewChild,
} from '@angular/core';
import { AccommodationDetailService } from '../../../services/user/accommodation-detail.service';
import { ActivatedRoute, Router } from '@angular/router';
import { CommonModule, NgClass, NgFor, NgIf } from '@angular/common';
import { TuiLike } from '@taiga-ui/kit';
import { ImageListModalComponent } from '../../../components/modals/image-list-modal/image-list-modal.component';
import { NavbarComponent } from '../../../components/navbar/navbar.component';
import SearchBoxComponent from '../../../components/search-box/search-box.component';
import { RoomService } from '../../../services/user/room.service';
import { RoomInformationModalComponent } from '../../../components/modals/room-information-modal/room-information-modal.component';
import { ReviewService } from '../../../services/user/review.service';
import { GetAccommodationByIdResponse } from '../../../models/manager/accommodation.model';
import { GetReviewsByAccommodationIdResponse } from '../../../models/user/review.model';
import { ReviewListModalComponent } from '../../../components/modals/review-list-modal/review-list-modal.component';
import { PaymentService } from '../../../services/user/payment.service';
import { AddressService } from '../../../services/address/address.service';
import { City } from '../../../models/address/address.model';
import { GetToken } from '../../../shared/token/token';
import { MessageService } from 'primeng/api';
import { Toast } from 'primeng/toast';
import { ButtonModule } from 'primeng/button';
import { catchError, finalize, forkJoin, of } from 'rxjs';
import { LoaderComponent } from '../../../components/loader/loader.component';
import { DomSanitizer, SafeResourceUrl } from '@angular/platform-browser';

@Component({
    selector: 'app-accommodation-detail',
    imports: [
        NgIf,
        NgFor,
        NgClass,
        TuiLike,
        ImageListModalComponent,
        NavbarComponent,
        SearchBoxComponent,
        RoomInformationModalComponent,
        CommonModule,
        ReviewListModalComponent,
        Toast,
        ButtonModule,
        LoaderComponent,
    ],
    templateUrl: './accommodation-detail.component.html',
    styleUrl: './accommodation-detail.component.scss',
    providers: [MessageService],
})
export class AccommodationDetailComponent implements OnInit {
    @ViewChild('availablilityRoomTop') availablilityRoomTop!: ElementRef;
    accommodationId: string = '';
    accommodationCity: string = '';
    accommodationDistrict: string = '';
    checkIn: string = '';
    checkOut: string = '';
    accommodation: any;
    rooms: any[] = [];
    reviews: any[] = [];
    isModalOpen: boolean = false;
    isRoomInformationModalOpen: boolean = false;
    isReviewModalOpen: boolean = false;
    isReviewListModalOpen: boolean = false;
    roomInformationSelected: any = null;
    reviewSelected: any = null;
    windowWidth: number = 0;
    showFull: boolean = false;
    isMobile: boolean = false;
    isLoading: boolean = false;
    bedTypes = [
        {
            key: 'single_bed',
            label: 'giường đơn',
            icon: 'icons/accommodation-detail-icon/single-bed-icon.svg',
            containerClass: 'single-bed-icon-container',
        },
        {
            key: 'double_bed',
            label: 'giường đôi',
            icon: 'icons/accommodation-detail-icon/full-bed-icon.svg',
            containerClass: 'full-bed-icon-container',
        },
        {
            key: 'large_double_bed',
            label: 'giường đôi lớn',
            icon: 'icons/accommodation-detail-icon/full-bed-icon.svg',
            containerClass: 'full-bed-icon-container',
        },
        {
            key: 'extra_large_double_bed',
            label: 'giường đôi siêu lớn',
            icon: 'icons/accommodation-detail-icon/full-bed-icon.svg',
            containerClass: 'full-bed-icon-container',
        },
    ];
    selectedBedType: string = '';
    selectedRooms: {
        [roomId: number]: { quantity: number; total: number };
    } = {};
    avarageRating: number = 0;
    scrollY: number = 0;
    constructor(
        private accommodationDetailService: AccommodationDetailService,
        private route: ActivatedRoute,
        private roomService: RoomService,
        private reviewService: ReviewService,
        private paymentService: PaymentService,
        private addressService: AddressService,
        private messageService: MessageService,
        private sanitizer: DomSanitizer,
        private router: Router
    ) {
        this.windowWidth = window.innerWidth; // Gán giá trị của windowWidth bằng với width của màn hình
        this.updateDescription();
    }

    // Method để sanitize URL
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

    // Lắng nghe sự kiện mỗi khi thay đổi kích thước màn hình
    @HostListener('window:resize', ['$event'])
    onResize(event: any) {
        this.windowWidth = window.innerWidth; // Gán giá trị của windowWidth bằng với width của màn hình
        this.updateDescription();
    }

    ngOnInit(): void {
        this.isLoading = true;
        this.accommodationId = this.route.snapshot.paramMap.get('id') ?? ''; // Lấy giá trị name trong url
        this.route.queryParams.subscribe((params) => {
            this.checkIn = params['checkIn'];
            this.checkOut = params['checkOut'];
        });

        if (this.accommodationId) {
            // this.getAccommodationById(this.accommodationId);
            // this.getRoomByAccommodationId(this.accommodationId);
            // this.getReviewByAccommodationId(this.accommodationId);
            this.loadAllDataAdvanced();
        } else {
            this.showToast(
                'error',
                'Lỗi tải dữ liệu',
                'Không tìm thấy thông tin chỗ ở.'
            );
        }
    }

    loadAllDataAdvanced() {
        forkJoin({
            accommodation: this.accommodationDetailService
                .getAccommodationDetailById(this.accommodationId)
                .pipe(
                    catchError((error) => {
                        console.error('Lỗi API accommodation:', error);
                        return of(null); // Trả về null thay vì throw error
                    })
                ),
            rooms: this.roomService
                .getRoomDetailByAccommodationId(
                    this.accommodationId,
                    this.checkIn,
                    this.checkOut
                )
                .pipe(
                    catchError((error) => {
                        console.error('Lỗi API rooms:', error);
                        return of(null);
                    })
                ),
            reviews: this.reviewService
                .getReviewsByAccommodationId(this.accommodationId)
                .pipe(
                    catchError((error) => {
                        console.error('Lỗi API reviews:', error);
                        return of(null);
                    })
                ),
        })
            .pipe(finalize(() => (this.isLoading = false)))
            .subscribe({
                next: (results) => {
                    // Xử lý từng kết quả riêng biệt
                    if (results.accommodation) {
                        this.handleAccommodationData(results.accommodation);
                    } else {
                        this.showToast(
                            'error',
                            'Lỗi',
                            'Không tải được thông tin chỗ ở.'
                        );
                    }

                    if (results.rooms) {
                        this.handleRoomsData(results.rooms);
                    } else {
                        this.showToast(
                            'error',
                            'Lỗi',
                            'Không tải được thông tin phòng.'
                        );
                    }

                    if (results.reviews) {
                        this.handleReviewsData(results.reviews);
                    } else {
                        this.showToast(
                            'info',
                            'Cảnh báo',
                            'Không có đánh giá nào.'
                        );
                    }
                },
                error: (error) => {
                    // Trường hợp này hiếm khi xảy ra vì đã catchError ở trên
                    console.error('Lỗi không mong muốn:', error);
                },
            });
    }

    private handleAccommodationData(data: any) {
        if (data?.data) {
            this.accommodation = data.data;
            this.getCityBySlug(this.accommodation.city);
        } else {
            this.showToast(
                'error',
                'Lỗi tải dữ liệu',
                'Không tìm thấy thông tin chỗ ở.'
            );
        }
    }

    private handleRoomsData(data: any) {
        if (data?.data) {
            this.rooms = data.data;
        } else {
            this.showToast(
                'error',
                'Lỗi tải dữ liệu',
                'Không tìm thấy thông tin phòng.'
            );
        }
    }

    private handleReviewsData(data: any) {
        if (data?.data) {
            this.reviews = data.data;
            const totalRating = this.reviews.reduce(
                (sum: number, review: any) => sum + review.rating,
                0
            );
            this.avarageRating =
                Math.floor((totalRating / this.reviews.length) * 10) / 10;
        } else {
            this.reviews = [];
            this.avarageRating = 0;
        }
    }

    createPayment() {
        if (this.numberRoomSelected() === 0) {
            this.showToast(
                'warn',
                'Cảnh báo',
                'Vui lòng chọn ít nhất một phòng trước khi thanh toán.'
            );
            return;
        }

        const roomSelected = Object.entries(this.selectedRooms).map(
            ([roomId, { quantity, total }]) => ({
                id: String(roomId),
                quantity,
            })
        );

        const payment = {
            check_in: this.checkIn,
            check_out: this.checkOut,
            accommodation_id: this.accommodation.id,
            room_selected: roomSelected,
        };

        const token = GetToken();

        if (token == null) {
            this.router.navigate(['/login']);
            return;
        }

        this.paymentService.createPayment(payment).subscribe({
            next: (response) => {
                this.showToast(
                    'success',
                    'Thành công',
                    'Tạo liên kết thanh toán thành công. Vui lòng kiểm tra liên kết.'
                );
                if (!response.body.data.url) {
                    this.showToast(
                        'error',
                        'Lỗi',
                        'Liên kết thanh toán không có trong phản hồi.'
                    );
                    return;
                }

                // Open payment page in new tab
                window.open(response.body.data.url, '_blank');
            },
            error: (error) => {
                this.showToast(
                    'error',
                    'Lỗi',
                    error.error.message ||
                        'Đã xảy ra lỗi khi tạo liên kết thanh toán. Vui lòng thử lại sau.'
                );
                // console.error('Error creating payment URL:', error);
            },
        });
    }

    getCityBySlug(slug: string) {
        this.addressService.getCityBySlug(slug).subscribe((data: City[]) => {
            if (data) {
                this.accommodationCity = data[0].name;

                const district = data[0].level2s.find(
                    (d) => d.slug === this.accommodation.district
                );
                this.accommodationDistrict = district?.name ?? '';
            } else {
                this.showToast(
                    'error',
                    'Lỗi tải dữ liệu',
                    'Không tìm thấy thông tin thành phố cho chỗ ở này.'
                );
            }
        });
    }

    goToLink(url: string) {
        window.open(url, '_blank');
    }

    openModal() {
        this.isModalOpen = true;
    }

    closeModal() {
        this.isModalOpen = false;
    }

    // Dựa vào giá trị của windowWidth để kiểm tra có phải mobile không
    updateDescription() {
        if (this.windowWidth <= 768) {
            this.showFull = false;
            this.isMobile = true;
        } else {
            this.showFull = true;
            this.isMobile = false;
        }
    }

    // Hiện thêm hay thu gọn description
    toggleDescription() {
        this.showFull = !this.showFull;
    }

    onRoomSelect(room: any, value: string) {
        const [priceStr, quantityStr] = value.split(',');
        const price = Number(priceStr);
        const quantity = Number(quantityStr);

        this.selectedRooms[room.id] = {
            quantity,
            total: price,
        };
    }

    numberRoomSelected(): number {
        return Object.values(this.selectedRooms)
            .filter((room) => room && typeof room.quantity === 'number')
            .reduce((sum, room) => sum + room.quantity, 0);
    }

    getTotalPrice(): number {
        return Object.values(this.selectedRooms).reduce(
            (acc, room) => acc + room.total,
            0
        );
    }

    getTotalBeds(beds: any): number {
        return Object.values(beds || {}).reduce(
            (total: number, count: any) => total + (+count || 0),
            0
        );
    }

    toggleOpenModal(room: any) {
        this.roomInformationSelected = room;
        this.isRoomInformationModalOpen = true;
        this.scrollY = window.scrollY;
        document.body.style.position = 'fixed';
        document.body.style.top = `-${this.scrollY}px`;
        document.body.style.left = '0';
        document.body.style.right = '0';
        document.body.style.width = '100%';
    }

    toggleCloseModal() {
        this.roomInformationSelected = null;
        this.isRoomInformationModalOpen = false;
        document.body.style.position = '';
        document.body.style.top = '';
        document.body.style.left = '';
        document.body.style.right = '';
        document.body.style.width = '';
        // Trả lại vị trí scroll cũ
        window.scrollTo(0, this.scrollY);
    }

    toggleOpenReviewModal(review: any) {
        this.reviewSelected = review;
        this.isReviewModalOpen = true;
        this.scrollY = window.scrollY;
        document.body.style.position = 'fixed';
        document.body.style.top = `-${this.scrollY}px`;
        document.body.style.left = '0';
        document.body.style.right = '0';
        document.body.style.width = '100%';
    }

    toggleCloseReviewModal() {
        this.reviewSelected = null;
        this.isReviewModalOpen = false;
        document.body.style.position = '';
        document.body.style.top = '';
        document.body.style.left = '';
        document.body.style.right = '';
        document.body.style.width = '';
        // Trả lại vị trí scroll cũ
        window.scrollTo(0, this.scrollY);
    }

    toggleOpenReviewListModal() {
        this.isReviewListModalOpen = true;
        this.scrollY = window.scrollY;
        document.body.style.position = 'fixed';
        document.body.style.top = `-${this.scrollY}px`;
        document.body.style.left = '0';
        document.body.style.right = '0';
        document.body.style.width = '100%';
    }

    toggleCloseReviewListModal() {
        this.isReviewListModalOpen = false;
        document.body.style.position = '';
        document.body.style.top = '';
        document.body.style.left = '';
        document.body.style.right = '';
        document.body.style.width = '';
        // Trả lại vị trí scroll cũ
        window.scrollTo(0, this.scrollY);
    }

    goToSelectRoom() {
        // Cuộn đến phần tử
        this.availablilityRoomTop.nativeElement.scrollIntoView({
            behavior: 'smooth',
        });
    }

    changeDateDDMMYYYYToYYYYMMDD(date: string): string {
        const [day, month, year] = date.split('-');
        return `${year}-${month}-${day}`; // yyyy-MM-dd
    }

    getNumberOfNights(): number {
        if (!this.checkIn || !this.checkOut) return 0;

        // Nếu ngày có dạng DD-MM-YYYY thì chuyển sang YYYY-MM-DD
        const formatDate = (dateStr: string) => {
            if (/^\d{2}-\d{2}-\d{4}$/.test(dateStr)) {
                const [day, month, year] = dateStr.split('-');
                return `${year}-${month}-${day}`;
            }
            return dateStr;
        };

        const checkInDate = new Date(formatDate(this.checkIn));
        const checkOutDate = new Date(formatDate(this.checkOut));

        if (isNaN(checkInDate.getTime()) || isNaN(checkOutDate.getTime()))
            return 0;

        const diffTime = checkOutDate.getTime() - checkInDate.getTime();
        const diffDays = Math.ceil(diffTime / (1000 * 60 * 60 * 24));
        return diffDays > 0 ? diffDays : 1;
    }
}
