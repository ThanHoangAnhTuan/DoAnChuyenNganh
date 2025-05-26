import { Component, HostListener, OnInit } from '@angular/core';
import { AccommodationDetailService } from '../../../services/user/accommodation-detail.service';
import { ActivatedRoute } from '@angular/router';
import { NgClass, NgFor, NgIf } from '@angular/common';
import { TuiLike } from '@taiga-ui/kit';
import { ImageListModalComponent } from '../../../components/modals/image-list-modal/image-list-modal.component';
import { NavbarComponent } from "../../../components/navbar/navbar.component";
import SearchBoxComponent from "../../../components/search-box/search-box.component";
import { RoomService } from '../../../services/user/room.service';
import { RoomInformationModalComponent } from "../../../components/modals/room-information-modal/room-information-modal.component";

@Component({
  selector: 'app-accommodation-detail',
  imports: [NgIf, NgFor, NgClass, TuiLike, ImageListModalComponent, NavbarComponent, SearchBoxComponent, RoomInformationModalComponent],
  templateUrl: './accommodation-detail.component.html',
  styleUrl: './accommodation-detail.component.scss'
})
export class AccommodationDetailComponent implements OnInit {
  accommodationId: string = '';
  accommodation: any;
  rooms: any[] = [];
  isModalOpen: boolean = false;
  isRoomInformationModalOpen: boolean = false;
  roomInformationSelected: any = null;
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

  constructor(
    private accommodationDetailService: AccommodationDetailService,
    private route: ActivatedRoute,
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
    this.accommodationId = this.route.snapshot.paramMap.get('id') ?? ''; // Lấy giá trị name trong url

    if (this.accommodationId) {
      this.getAccommodationById(this.accommodationId);
      this.getRoomByAccommodationId(this.accommodationId);
      // this.getHotelByName(this.accommodationName);
    } else {
      console.error('Accommodation name is missing in URL');
    }
  }

  getAccommodationById(id: string) {
    this.accommodationDetailService.getAccommodationDetailById(id).subscribe((data: any) => {
      this.accommodation = data.data;
      console.log("accommodation: ", this.accommodation);
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
}
