import { Component, HostListener, OnInit } from '@angular/core';
import { AccommodationDetailService } from '../../../services/user/accommodation-detail.service';
import { ActivatedRoute } from '@angular/router';
import { NgClass, NgFor, NgIf } from '@angular/common';
import { TuiLike } from '@taiga-ui/kit';
import { ImageListModalComponent } from '../../../components/image-list-modal/image-list-modal.component';
import { NavbarComponent } from "../../../components/navbar/navbar.component";
import SearchBoxComponent from "../../../components/search-box/search-box.component";
import { HotelService } from '../../../services/user/hotel.service';
import { RoomService } from '../../../services/user/room.service';

@Component({
  selector: 'app-accommodation-detail',
  imports: [NgIf, NgFor, NgClass, TuiLike, ImageListModalComponent, NavbarComponent, SearchBoxComponent],
  templateUrl: './accommodation-detail.component.html',
  styleUrl: './accommodation-detail.component.scss'
})
export class AccommodationDetailComponent implements OnInit {
  accommodationName: string = '';
  accommodationId: string = '';
  accommodation: any;
  hotel: any;
  rooms: any[] = [];
  isModalOpen: boolean = false;
  isRoomInformationModalOpen: boolean = false;
  windowWidth: number = 0;
  showFull: boolean = false;
  isMobile: boolean = false;
  bedTypes = [
    {
      key: 'single_bed',
      label: 'single bed',
      icon: 'icons/accommodation-detail-icon/single-bed-icon.svg',
      
      containerClass: 'single-bed-icon-container'
    },
    {
      key: 'double_bed',
      label: 'full bed',
      icon: 'icons/accommodation-detail-icon/full-bed-icon.svg',
      containerClass: 'full-bed-icon-container'
    },
    {
      key: 'large_double_bed',
      label: 'large full bed',
      icon: 'icons/accommodation-detail-icon/full-bed-icon.svg',
      containerClass: 'full-bed-icon-container'
    },
    {
      key: 'extra_large_double_bed',
      label: 'extra large full bed',
      icon: 'icons/accommodation-detail-icon/full-bed-icon.svg',
      containerClass: 'full-bed-icon-container'
    }
  ];
  totalPrice: number = 0;
  numberOfRooms: number = 0;
  selectedBedType: string = '';

  get allAvailableRooms(): any[] {
    return this.rooms.flatMap(room =>
      this.getRoomByAvailableRooms(room.available_rooms).map(() => room)
    );
  }

  constructor(
    private accommodationDetailService: AccommodationDetailService,
    private route: ActivatedRoute,
    private hotelService: HotelService,
    private roomService: RoomService
  ) {
    this.windowWidth = window.innerWidth; // Gán giá trị của windowWidth bằng với width của màn hình
    this.updateDescription();
    console.log(this.isRoomInformationModalOpen);
  }

  // Lắng nghe sự kiện mỗi khi thay đổi kích thước màn hình
  @HostListener('window:resize', ['$event'])
  onResize(event: any) {
    this.windowWidth = window.innerWidth; // Gán giá trị của windowWidth bằng với width của màn hình
    this.updateDescription();
  }

  ngOnInit(): void {
    this.accommodationName = this.route.snapshot.paramMap.get('name') ?? ''; // Lấy giá trị name trong url
    this.route.queryParams.subscribe(params => {
      this.accommodationId = params['id'];
    });

    if (this.accommodationName) {
      this.getAccommodationByName(this.accommodationName);
      this.getHotelByName(this.accommodationName);
    } else {
      console.error('Accommodation name is missing in URL');
    }

    if (this.accommodationId) {
      this.getRoomByAccommodationId(this.accommodationId);
    } else {
      console.error('Accommodation id is missing in URL');
    }
  };

  getAccommodationByName(name: string) {
    this.accommodationDetailService.getAccommodationDetailByName(name).subscribe((data: any) => {
      this.accommodation = data.data[0];
      // console.log("accommodation: ", this.accommodation);
    })
  }

  getHotelByName(name: string) {
    this.hotelService.getHotelDetailByName(name).subscribe((data: any) => {
      this.hotel = data[0];
    })
  }

  getRoomByAccommodationId(id: string) {
    this.roomService.getRoomDetailByAccommodationId(id).subscribe((data: any) => {
      this.rooms = data.data;
      console.log("room: ", this.rooms);
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

  getRoomByAvailableRooms(count: number): number[] {
    return Array.from({ length: count }, (_, i) => i + 1);
  }

  addPriceToTotal(value: string) {
    const [priceStr, numberOfRoomsStr] = value.split(',');
    const price = Number(priceStr);
    const numberRooms = Number(numberOfRoomsStr);

    this.totalPrice += price;
    this.totalPrice = Math.round(this.totalPrice * 100) / 100;

    this.numberOfRooms += numberRooms;
  }

  getTotalBeds(beds: any): number {
    return Object.values(beds || {}).reduce((total: number, count: any) => total + (+count || 0), 0);
  }

  toggleOpenModal() {
    this.isRoomInformationModalOpen = true;
    document.body.style.overflow = 'hidden';
  }

  toggleCloseModal() {
    this.isRoomInformationModalOpen = false;
    document.body.style.overflow = 'auto';
  }
}
