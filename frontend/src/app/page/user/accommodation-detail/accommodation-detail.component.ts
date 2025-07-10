import {
    Component,
    ElementRef,
    HostListener,
    inject,
    OnInit,
    ViewChild,
} from '@angular/core';
import { AccommodationDetailService } from '../../../services/user/accommodation-detail.service';
import { ActivatedRoute } from '@angular/router';
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
import { finalize } from 'rxjs';
import { LoaderComponent } from "../../../components/loader/loader.component";

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
    LoaderComponent
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
        private messageService: MessageService
    ) {
        this.windowWidth = window.innerWidth; // Gán giá trị của windowWidth bằng với width của màn hình
        this.updateDescription();
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
        this.accommodationId = this.route.snapshot.paramMap.get('id') ?? ''; // Lấy giá trị name trong url
        this.route.queryParams.subscribe((params) => {
            this.checkIn = params['checkIn'];
            this.checkOut = params['checkOut'];
        });

        if (this.accommodationId) {
            this.getAccommodationById(this.accommodationId);
            this.getRoomByAccommodationId(this.accommodationId);
            this.getReviewByAccommodationId(this.accommodationId);
        } else {
            this.showToast(
                'error',
                'Lỗi tải dữ liệu',
                'Không tìm thấy thông tin chỗ ở.'
            );
        }
    }

    getAccommodationById(id: string) {
        this.isLoading = true;
        this.accommodationDetailService
            .getAccommodationDetailById(id)
            .pipe(finalize(() => (this.isLoading = false)))
            .subscribe((data: GetAccommodationByIdResponse) => {
                if (data) {
                    this.accommodation = data.data;
                    this.getCityBySlug(this.accommodation.city);
                } else {
                    this.showToast(
                        'error',
                        'Lỗi tải dữ liệu',
                        'Không tìm thấy thông tin chỗ ở.'
                    );
                }
            });
    }

    getRoomByAccommodationId(id: string) {
        this.isLoading = true;
        this.roomService
            .getRoomDetailByAccommodationId(id, this.checkIn, this.checkOut)
            .pipe(finalize(() => (this.isLoading = false)))
            .subscribe({
                next: (value) => {
                    this.rooms = value.data;
                },
                error: (err) => {
                    this.showToast(
                        'error',
                        'Lỗi tải dữ liệu',
                        'Không tìm thấy thông tin phòng cho chỗ ở này.'
                    );
                },
            });
    }

    getReviewByAccommodationId(id: string) {
        this.isLoading = true;
        this.reviewService
            .getReviewsByAccommodationId(id)
            .pipe(finalize(() => (this.isLoading = false)))
            .subscribe((data: GetReviewsByAccommodationIdResponse) => {
                if (data) {
                    // const sortedReviews = data.data.sort((a, b) => {
                    //   return new Date(b.created_at).getTime() - new Date(a.created_at).getTime();
                    // });

                    this.reviews = data.data;

                    const totalRating = this.reviews.reduce(
                        (sum: number, review: any) => sum + review.rating,
                        0
                    );
                    this.avarageRating =
                        Math.floor((totalRating / this.reviews.length) * 10) /
                        10;
                } else {
                    this.reviews = [];
                    this.avarageRating = 0;
                    this.showToast(
                        'info',
                        'Cảnh báo',
                        'Không có đánh giá nào cho chỗ ở này.'
                    );
                }
            });
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
            this.showToast(
                'warn',
                'Cảnh báo',
                'Vui lòng đăng nhập trước khi thanh toán chỗ ở.'
            );
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
}
