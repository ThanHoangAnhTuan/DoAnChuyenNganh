import { Component, OnInit } from '@angular/core';
import { AccommodationDetailService } from '../../services/accommodation-detail.service';
import { ActivatedRoute } from '@angular/router';
import { NgFor } from '@angular/common';
import { TuiLike } from '@taiga-ui/kit';
import { TuiIcon } from '@taiga-ui/core';

@Component({
  selector: 'app-accommodation-detail',
  imports: [NgFor, TuiLike, TuiIcon],
  templateUrl: './accommodation-detail.component.html',
  styleUrl: './accommodation-detail.component.scss'
})
export class AccommodationDetailComponent implements OnInit {
  accommodationDetail: any[] = [];

  constructor(private accommodationDetailService: AccommodationDetailService, private route: ActivatedRoute) { }

  ngOnInit(): void {
    const city = this.route.snapshot.paramMap.get('name');
    console.log('City param:', city);

    // if (city) {
    //   this.getAccommodationDetail(city);
    // } else {
    //   console.error('City param is missing in URL');
    // }

    this.getGet();
  };

  getGet() {
    this.accommodationDetailService.getAccommodationDetail().subscribe((data: any) => {
      this.accommodationDetail = data;
    });
  }

  getAccommodationDetail(city: string) {
    this.accommodationDetailService.getAccommodationDetailByCity(city).subscribe((data: any) => {
      this.accommodationDetail = data;
    });
  }
}
