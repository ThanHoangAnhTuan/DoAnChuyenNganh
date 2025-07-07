import { Component, Input, OnInit } from '@angular/core';
import { GalleriaModule } from 'primeng/galleria';

@Component({
  selector: 'app-room-information-modal',
  imports: [GalleriaModule],
  templateUrl: './room-information-modal.component.html',
  styleUrl: './room-information-modal.component.scss'
})
export class RoomInformationModalComponent implements OnInit {
    activeIndex: number = 0;
    @Input() roomName: string = '';
    @Input() roomImages: string[] = [];

    ngOnInit(): void {
    }
    
    // Responsive của Galleria
    responsiveOptions = [
        {
            breakpoint: '1444px',
            numVisible: 4
        },
        {
            breakpoint: '1024px',
            numVisible: 4
        },
        {
            breakpoint: '950px',
            numVisible: 4
        },
        {
            breakpoint: '700px',
            numVisible: 3
        },
        {
            breakpoint: '560px',
            numVisible: 2
        },
        {
            breakpoint: '390px',
            numVisible: 1
        }
    ];
}
