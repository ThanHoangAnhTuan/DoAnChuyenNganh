import { Component, HostListener, OnInit } from '@angular/core';
import { AccommodationDetailService } from '../../../services/user/accommodation-detail.service';
import { ActivatedRoute } from '@angular/router';
import { NgFor, NgIf } from '@angular/common';
import { TuiLike } from '@taiga-ui/kit';
import { TuiIcon } from '@taiga-ui/core';
import { ImageListModalComponent } from '../../../components/image-list-modal/image-list-modal.component';
import { NavbarComponent } from "../../../components/navbar/navbar.component";
import SearchBoxComponent from "../../../components/search-box/search-box.component";

@Component({
  selector: 'app-accommodation-detail',
  imports: [NgIf, NgFor, TuiLike, TuiIcon, ImageListModalComponent, NavbarComponent, SearchBoxComponent],
  templateUrl: './accommodation-detail.component.html',
  styleUrl: './accommodation-detail.component.scss'
})
export class AccommodationDetailComponent implements OnInit {
  accommodation: any;
  isModalOpen: boolean = false;
  windowWidth: number = 0;
  showFull: boolean = false;
  isMobile: boolean = false;

  constructor(private accommodationDetailService: AccommodationDetailService, private route: ActivatedRoute) {
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
    const accommodationName = this.route.snapshot.paramMap.get('name'); // Lấy giá trị name trong url

    if (accommodationName) {
      this.getAccommodationByName(accommodationName);
    } else {
      console.error('City name is missing in URL');
    }
  };

  getAccommodationByName(name: string) {
    this.accommodationDetailService.getAccommodationDetailByName(name).subscribe((data: any) => {
      this.accommodation = data[0];
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
}
