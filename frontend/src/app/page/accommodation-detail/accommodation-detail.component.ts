import { Component, HostListener, OnInit } from '@angular/core';
import { AccommodationDetailService } from '../../services/accommodation-detail.service';
import { ActivatedRoute } from '@angular/router';
import { NgFor, NgIf } from '@angular/common';
import { TuiLike } from '@taiga-ui/kit';
import { TuiIcon } from '@taiga-ui/core';
import { ImageListModalComponent } from '../../components/image-list-modal/image-list-modal.component';
import { NavbarComponent } from "../../components/navbar/navbar.component";
import SearchBoxComponent from "../../components/search-box/search-box.component";

@Component({
  selector: 'app-accommodation-detail',
  imports: [NgIf, NgFor, TuiLike, TuiIcon, ImageListModalComponent, NavbarComponent, SearchBoxComponent],
  templateUrl: './accommodation-detail.component.html',
  styleUrl: './accommodation-detail.component.scss'
})
export class AccommodationDetailComponent implements OnInit {
  accommodation: any;
  accommodationId: string = '';
  accommodationName: string = '';
  isModalOpen: boolean = false;
  windowWidth: number = 0;
  showFull: boolean = false;
  isMobile: boolean = false;

  constructor(private accommodationDetailService: AccommodationDetailService, private route: ActivatedRoute) {
    this.windowWidth = window.innerWidth;
    this.updateDescription();
  }

  @HostListener('window:resize', ['$event'])
  onResize(event: any) {
    this.windowWidth = window.innerWidth;
    this.updateDescription();
  }

  ngOnInit(): void {
    this.route.queryParams.subscribe(params => {
      this.accommodationId = params['id'];
      this.accommodationName = params['name'];

      if (this.accommodationId != null && this.accommodationId !== '') {
        this.getAccommodationByName(this.accommodationName);
      } else {
        console.error('City param is missing in URL');
      }
    });
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

  updateDescription() {
    if (this.windowWidth <= 768) {
      this.showFull = false;
      this.isMobile = true
    } else {
      this.showFull = true;
      this.isMobile = false;
    }
  }

  toggleDescription() {
    this.showFull = !this.showFull;
  }
}
