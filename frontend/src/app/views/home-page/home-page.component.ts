import { NgFor } from '@angular/common';
import { ChangeDetectionStrategy, Component, HostListener, ViewChild } from '@angular/core';
import { TuiCarousel, TuiCarouselComponent } from '@taiga-ui/kit';
import { TuiButton } from '@taiga-ui/core';

@Component({
  selector: 'app-home-page',
  imports: [NgFor, TuiCarousel, TuiButton, TuiCarouselComponent],
  templateUrl: './home-page.component.html',
  styleUrl: './home-page.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})

export class HomePageComponent {
  trendinPlaces: any[] = [];
  explorePaces: any[] = [];
  windowWidth: number = 0;
  showPlaceTypeList: number = 0;
  showExplorePlacesList: number = 0;

  protected placeTypeIndex = 0;
  protected explorePlaceIndex = 0;

  @ViewChild('placeTypeCarousel') placeTypeCarousel!: TuiCarouselComponent;
  @ViewChild('explorePlacesCarousel') explorePlacesCarousel!: TuiCarouselComponent;

  protected readonly places = [
    { name: 'Khách sạn', image: 'khach-san.jpg' },
    { name: 'Căn hộ', image: 'can-ho.jpg' },
    { name: 'Các resort', image: 'resort.jpg' },
    { name: 'Các biệt thự', image: 'biet-thu.jpg' },
    { name: 'Cabin nghỉ dưỡng', image: 'cabin-nghi-duong.jpg' },
    { name: 'Các nhà nghỉ dưỡng', image: 'nha-nghi-duong.jpg' },
    { name: 'Các nhà khách', image: 'nha-khach.jpg' },
    { name: 'Các hostel', image: 'hostel.jpg' },
    { name: 'Các motel', image: 'motel.jpg' },
    { name: 'Nhà nghỉ B&B', image: 'nha-nghi-b&b.jpg' },
    { name: 'Các riad', image: 'riad.jpg' },
    { name: 'Các công viên nghỉ dưỡng', image: 'cong-vien-nghi-duong.jpg' },
    { name: 'Homestay', image: 'homestay.jpg' },
    { name: 'Các khu cắm trại', image: 'khu-cam-trai.jpg' },
    { name: 'Biệt thự đồng quê', image: 'biet-thu-dong-que.jpg' },
    { name: 'Các nhà nghỉ trang trại', image: 'nha-nghi-trang-trai.jpg' },
    { name: 'Lều trại sang trọng', image: 'leu-trai-sang-trong.jpg' }
  ];

  constructor() {
    this.trendinPlaces = [
      { name: 'Hồ Chí Minh', image: 'ho-chi-minh-city.jpg', alt: 'trending place 1' },
      { name: 'Hà Nội', image: 'ha-noi.jpg', alt: 'trending place 2' },
      { name: 'Đà Nẵng', image: 'da-nang.jpg', alt: 'trending place 3' },
      { name: 'Đà Lạt', image: 'da-lat.jpg', alt: 'trending place 4' },
      { name: 'Vũng Tàu', image: 'vung-tau.jpg', alt: 'trending place 5' },
    ];

    this.explorePaces = [
      { name: 'Hà Nội', image: 'ha-noi.png' },
      { name: 'Bình Thuận', image: 'binh-thuan.png' },
      { name: 'Hồ Chí Minh', image: 'ho-chi-minh-city.png' },
      { name: 'Vũng Tàu', image: 'vung-tau.png' },
      { name: 'Hưng Yên', image: 'hung-yen.png' },
      { name: 'Đà Lạt', image: 'da-lat.png' },
      { name: 'Đồng Nai', image: 'dong-nai.png' },
      { name: 'Bình Định', image: 'binh-dinh.png' },
      { name: 'Ninh Bình', image: 'ninh-binh.png' },
      { name: 'Nha Trang', image: 'nha-trang.png' },
      { name: 'Cần Thơ', image: 'can-tho.png' },
      { name: 'Huế', image: 'hue.png' },
      { name: 'Đà Nẵng', image: 'da-nang.png' },
      { name: 'Bắc Ninh', image: 'bac-ninh.png' },
      { name: 'Cao Bằng', image: 'cao-bang.png' },
    ]

    this.windowWidth = window.innerWidth;
    this.updateCarouselVisibility();
  }

  @HostListener('window:resize', ['$event'])
  onResize(event: any) {
    this.windowWidth = window.innerWidth;
    this.updateCarouselVisibility();
  }

  updateCarouselVisibility() {
    if (this.windowWidth >= 2500) {
      this.showPlaceTypeList = 7;
      this.showExplorePlacesList = 8;
    } else if (this.windowWidth >= 2300) {
      this.showPlaceTypeList = 6;
      this.showExplorePlacesList = 7;
    } else if (this.windowWidth >= 1800) {
      this.showPlaceTypeList = 5;
      this.showExplorePlacesList = 6;
    } else if (this.windowWidth >= 1025) {
      this.showPlaceTypeList = 4;
      this.showExplorePlacesList = 5;
    } else if (this.windowWidth >= 1000) {
      this.showPlaceTypeList = 5;
      this.showExplorePlacesList = 6;
    } else if (this.windowWidth >= 850) {
      this.showPlaceTypeList = 4;
      this.showExplorePlacesList = 5;
    } else if (this.windowWidth >= 750) {
      this.showPlaceTypeList = 3;
      this.showExplorePlacesList = 4;
    } else if (this.windowWidth >= 500) {
      this.showPlaceTypeList = 2;
      this.showExplorePlacesList = 3;
    } else {
      this.showPlaceTypeList = 1;
      this.showExplorePlacesList = 2;
    }
  }
}