import { Component, HostListener, OnInit } from '@angular/core';
import { AccommodationDetailService } from '../../../services/user/accommodation-detail.service';
import { ActivatedRoute } from '@angular/router';
import { CommonModule, DatePipe, NgClass, NgFor, NgIf } from '@angular/common';
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
  providers: [DatePipe],
  templateUrl: './accommodation-detail.component.html',
  styleUrl: './accommodation-detail.component.scss'
})
export class AccommodationDetailComponent implements OnInit {
  accommodationId: string = '';
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
  numberOfRooms: number = 0;
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
    private datePipe: DatePipe,
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
      this.getReviewByAccommodationId(this.accommodationId);
    } else {
      console.error('Accommodation name is missing in URL');
    }
  }

  getAccommodationById(id: string) {
    this.accommodationDetailService.getAccommodationDetailById(id).subscribe((data: GetAccommodationByIdResponse) => {
      this.accommodation = data.data;
      // console.log("accommodation: ", this.accommodation);
    })
  }

  getRoomByAccommodationId(id: string) {
    this.roomService.getRoomDetailByAccommodationId(id).subscribe((data: GetAccommodationDetailsResponse) => {
      this.rooms = data.data;
      // console.log("room: ", this.rooms);
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
  }

  getTotalPrice(): number {
    return Object.values(this.selectedRooms).reduce((acc, room) => acc + room.total, 0);
  }

  getTotalBeds(beds: any): number {
    return Object.values(beds || {}).reduce((total: number, count: any) => total + (+count || 0), 0);
  }

  toggleOpenModal(room: any) {
    this.roomInformationSelected = room;
    this.isRoomInformationModalOpen = true;
    document.body.style.overflow = 'hidden';
  }

  toggleCloseModal() {
    this.roomInformationSelected = null;
    this.isRoomInformationModalOpen = false;
    document.body.style.overflow = 'auto';
  }

  formatDate(dateStr: string): string {
    const date = new Date(dateStr);
    const hoursMinutes = this.datePipe.transform(date, 'HH:mm');
    const day = date.getUTCDate();
    const month = date.getUTCMonth() + 1;
    const year = date.getUTCFullYear();
    return `${hoursMinutes} ngày ${day} tháng ${month} năm ${year}`;
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
