import { Component, OnInit } from '@angular/core';
import { AccommodationDetailService } from '../../services/accommodation-detail.service';
import { ActivatedRoute, RouterModule } from '@angular/router';
import { NgFor } from '@angular/common';
import { TuiLike } from '@taiga-ui/kit';

@Component({
  selector: 'app-search',
  imports: [NgFor, TuiLike, RouterModule],
  templateUrl: './search.component.html',
  styleUrl: './search.component.scss'
})
export class SearchComponent implements OnInit {
  searchCity: string = "";
  accommodations: any[] = [];

  constructor(private accommodationService: AccommodationDetailService, private route: ActivatedRoute) { }

  ngOnInit(): void {
    const city = this.route.snapshot.paramMap.get('name');

    if (city !== null) {
      this.searchCity = city;
    } else {
      this.searchCity = 'Not Found';
    }

    this.getAccommodationByCity(this.searchCity);
  }

  getAccommodationByCity(city: string) {
    this.accommodationService.getAccommodationDetailByCity(city).subscribe((data: any) => {
      this.accommodations = data;
    })
  }

  goToLink(url: string) {
    window.open(url, '_blank');
  }
}
