import { Component, HostListener, OnInit } from '@angular/core';
import { AccommodationDetailService } from '../../../services/user/accommodation-detail.service';
import { ActivatedRoute } from '@angular/router';
import { CommonModule, NgFor, NgIf } from '@angular/common';
import { TuiLike } from '@taiga-ui/kit';
import { TuiIcon } from '@taiga-ui/core';
import { ImageListModalComponent } from '../../../components/image-list-modal/image-list-modal.component';
import { NavbarComponent } from '../../../components/navbar/navbar.component';
import SearchBoxComponent from '../../../components/search-box/search-box.component';
import { AccommodationService } from '../../../services/user/accommodation.service';
import { Accommodation } from '../../../models/manager/accommodation.model';
import { DomSanitizer, SafeHtml } from '@angular/platform-browser';

@Component({
    selector: 'app-accommodation-detail',
    imports: [
        CommonModule,
        NgIf,
        NgFor,
        TuiLike,
        TuiIcon,
        ImageListModalComponent,
        NavbarComponent,
        SearchBoxComponent,
    ],
    templateUrl: './accommodation-detail.component.html',
    styleUrl: './accommodation-detail.component.scss',
})
export class AccommodationDetailComponent implements OnInit {
    accommodation !: Accommodation;
    isModalOpen: boolean = false;
    windowWidth: number = 0;
    showFull: boolean = false;
    isMobile: boolean = false;
    protected readonly apiUpload = "http://localhost:8080/uploads/"


    constructor(
        private accommodationService: AccommodationService,
        private route: ActivatedRoute,
        private sanitizer: DomSanitizer
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
        const accommodationId = this.route.snapshot.paramMap.get('id'); // Lấy giá trị Id trong url

        if (accommodationId) {
            this.getAccommodationById(accommodationId);
        } else {
            console.error('City Id is missing in URL');
        }
    }

    protected getDescription(html: string): SafeHtml {
        return this.sanitizer.bypassSecurityTrustHtml(html);
    }

    getAccommodationById(id: string) {
        this.accommodationService
            .getAccommodationDetailById(id)
            .subscribe((response) => {
                console.log(response.data);
                
                this.accommodation = response.data;
                console.log(this.accommodation);
                
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
}
