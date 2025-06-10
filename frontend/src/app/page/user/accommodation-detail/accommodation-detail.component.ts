import { Component, HostListener, inject, OnInit } from '@angular/core';
import { AccommodationDetailService } from '../../../services/user/accommodation-detail.service';
import { ActivatedRoute } from '@angular/router';
import { CommonModule, NgClass, NgFor, NgIf } from '@angular/common';
import { TuiLike } from '@taiga-ui/kit';
import { ImageListModalComponent } from '../../../components/modals/image-list-modal/image-list-modal.component';
import { NavbarComponent } from "../../../components/navbar/navbar.component";
import SearchBoxComponent from "../../../components/search-box/search-box.component";
import { RoomService } from '../../../services/user/room.service';
import { RoomInformationModalComponent } from "../../../components/modals/room-information-modal/room-information-modal.component";
import { ReviewService } from '../../../services/user/review.service';
import { GetAccommodationByIdResponse } from '../../../models/manager/accommodation.model';
import { GetAccommodationDetailsResponse } from '../../../models/manager/accommodation-detail.model';
import { Review } from '../../../models/user/review.model';
import { ReviewListModalComponent } from "../../../components/modals/review-list-modal/review-list-modal.component";
import { PaymentService } from '../../../services/user/payment.service';
import { TuiAlertService } from '@taiga-ui/core';
import { AddressService } from '../../../services/address/address.service';
import { City } from '../../../models/address/address.model';

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
    ReviewListModalComponent
  ],
  templateUrl: './accommodation-detail.component.html',
  styleUrl: './accommodation-detail.component.scss'
})
export class AccommodationDetailComponent implements OnInit {
  accommodationId: string = '';
  accommodationCity: string = '';
  accommodationDistrict: string = '';
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
  bedTypes = [
    {
      key: 'single_bed',
      label: 'giường đơn',
      icon: 'icons/accommodation-detail-icon/single-bed-icon.svg',
      containerClass: 'single-bed-icon-container'
    },
    {
      key: 'double_bed',
      label: 'giường đôi',
      icon: 'icons/accommodation-detail-icon/full-bed-icon.svg',
      containerClass: 'full-bed-icon-container'
    },
    {
      key: 'large_double_bed',
      label: 'giường đôi lớn',
      icon: 'icons/accommodation-detail-icon/full-bed-icon.svg',
      containerClass: 'full-bed-icon-container'
    },
    {
      key: 'extra_large_double_bed',
      label: 'giường đôi siêu lớn',
      icon: 'icons/accommodation-detail-icon/full-bed-icon.svg',
      containerClass: 'full-bed-icon-container'
    }
  ];
  selectedBedType: string = '';
  selectedRooms: {
    [roomId: number]: { quantity: number; total: number };
  } = {};
  avarageRating: number = 0;
  scrollY: number = 0;

  private readonly alerts = inject(TuiAlertService);

  protected getAlert(label: string, content: string): void {
    this.alerts
      .open(content, {
        label: label,
        appearance: 'negative',
        autoClose: 5000,
      })
      .subscribe();
  }

  constructor(
    private accommodationDetailService: AccommodationDetailService,
    private route: ActivatedRoute,
    private roomService: RoomService,
    private reviewService: ReviewService,
    private paymentService: PaymentService,
    private addressService: AddressService,
  ) {
    this.windowWidth = window.innerWidth; // Gán giá trị của windowWidth bằng với width của màn hình
    this.updateDescription();
  }

  // Lắng nghe sự kiện mỗi khi thay đổi kích thước màn hình
  @HostListener('window:resize', ['$event'])
  onResize(event: any) {
    this.windowWidth = window.innerWidth; // Gán giá trị của windowWidth bằng với width của màn hình
    this.updateDescription();
  }

  ngOnInit(): void {
    this.accommodationId = this.route.snapshot.paramMap.get('id') ?? ''; // Lấy giá trị name trong url

    if (this.accommodationId) {
      this.getAccommodationById(this.accommodationId);
      this.getRoomByAccommodationId(this.accommodationId);
      // this.getReviewByAccommodationId(this.accommodationId);
    } else {
      console.error('Accommodation name is missing in URL');
    }
  }

  getAccommodationById(id: string) {
    this.accommodationDetailService.getAccommodationDetailById(id).subscribe((data: GetAccommodationByIdResponse) => {
      if (data) {
        this.accommodation = data.data;
        // console.log("accommodation: ", this.accommodation);

        this.getCityById(this.accommodation.city);
      } else {
        console.log("Can't get accommodation");
      }
    })
  }

  getRoomByAccommodationId(id: string) {
    this.roomService.getRoomDetailByAccommodationId(id).subscribe((data: GetAccommodationDetailsResponse) => {
      if (data) {
        this.rooms = data.data;
        // console.log("room: ", this.rooms);
      } else {
        console.log("Can't get accommodation room");
      }
    })
  }

  getReviewByAccommodationId(id: string) {
    this.reviewService.getReviewsByAccommodationId(id).subscribe((data: Review[]) => {
      if (data && data.length > 0) {
        const sortedReviews = data.sort((a, b) => {
          return new Date(b.created_at).getTime() - new Date(a.created_at).getTime();
        });

        this.reviews = sortedReviews;

        const totalRating = this.reviews.reduce((sum: number, review: any) => sum + review.rating, 0);
        this.avarageRating = Math.floor((totalRating / this.reviews.length) * 10) / 10;

        // console.log("reviews: ", this.reviews);
      } else {
        this.reviews = [];
        this.avarageRating = 0;
        console.log("No reviews found for this accommodation.");
      }
    });
  }

  createPayment() {
    if (this.numberRoomSelected() === 0) {
      this.getAlert('Notification', 'Please select at least one room before proceeding to payment.');
      return;
    }

    const roomSelected = Object.entries(this.selectedRooms).map(([roomId, { quantity, total }]) => ({
      id: String(roomId),
      quantity,
    }));

    const payment = {
      check_in: "09-06-2025",
      check_out: "12-06-2025",
      accommodation_id: this.accommodation.id,
      room_selected: roomSelected,
    };

    const token = sessionStorage.getItem('token');

    if (token == null) {
      this.getAlert('Notification', 'Please log in before pay for accommodation');
      return;
    }

    this.paymentService.createPayment(payment).subscribe({
      next: (response) => {
        console.log('Payment URL created successfully:', response.body);

        if (!response.body.data.url) {
          this.getAlert('Notification', 'Payment URL is missing in the response.');
          return;
        }

        // Mở link thanh toán trong tab mới
        window.open(response.body.data.url, '_blank');
      },
      error: (error) => {
        console.error('Error creating payment URL:', error);
        this.getAlert('Notification', 'An error occurred while creating the payment URL. Please try again later.');
      }
    });
  }

  getCityById(id: string) {
    this.addressService.getCityByLevel1id(id).subscribe((data: City[]) => {
      if (data) {
        this.accommodationCity = data[0].name;

        const district = data[0].level2s.find(d => d.level2_id === this.accommodation.district);
        this.accommodationDistrict = district?.name ?? '';
        
        // console.log("City: ", this.accommodationCity);
        // console.log("District", this.accommodationDistrict);
      } else {
        console.log("Can't get city by id");
      }
    })
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
      this.isMobile = true
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

    console.log("Selected Rooms: ", this.selectedRooms);
    console.log('values:', Object.values(this.selectedRooms));
  }

  numberRoomSelected(): number {
    return Object.values(this.selectedRooms)
      .filter(room => room && typeof room.quantity === 'number')
      .reduce((sum, room) => sum + room.quantity, 0);
  }

  getTotalPrice(): number {
    return Object.values(this.selectedRooms)
      .reduce((acc, room) => acc + room.total, 0);
  }

  getTotalBeds(beds: any): number {
    return Object.values(beds || {}).reduce((total: number, count: any) => total + (+count || 0), 0);
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
}