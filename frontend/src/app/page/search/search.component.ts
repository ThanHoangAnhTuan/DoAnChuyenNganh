import { Component, OnInit } from '@angular/core';
import { AccommodationDetailService } from '../../services/accommodation-detail.service';
import { ActivatedRoute, RouterModule } from '@angular/router';
import { NgFor } from '@angular/common';
import { TuiLike } from '@taiga-ui/kit';
import { NavbarComponent } from "../../components/navbar/navbar.component";
import SearchBoxComponent from "../../components/search-box/search-box.component";

@Component({
  selector: 'app-search',
  imports: [NgFor, TuiLike, RouterModule, NavbarComponent, SearchBoxComponent],
  templateUrl: './search.component.html',
  styleUrl: './search.component.scss'
})
export class SearchComponent implements OnInit {
  cityId: string = '';
  cityName: string = '';
  accommodations: any[] = [];

  constructor(private accommodationService: AccommodationDetailService, private route: ActivatedRoute) { }

  ngOnInit(): void {
    this.route.queryParams.subscribe(params => {
      this.cityId = params['id'];
      this.cityName = this.normalizeCityName(params['name']);
    });

    this.getAccommodationByCity(this.cityName);
  }

  getAccommodationByCity(city: string) {
    this.accommodationService.getAccommodationDetailByCity(city).subscribe((data: any) => {
      this.accommodations = data;
    })
  }

  goToLink(url: string) {
    window.open(url, '_blank');
  }

  private normalizeCityName(name: string): string {
    const normalized = name.trim().toLowerCase();

    const cityMap: { [key: string]: string } = {
      'hồ chí minh': 'Hồ Chí Minh',
      'ho chi minh': 'Hồ Chí Minh',
      'ho chi minh city': 'Hồ Chí Minh',
      'hcm': 'Hồ Chí Minh',
      'hcmc': 'Hồ Chí Minh',
      'hochiminh': 'Hồ Chí Minh',
      'hochiminhcity': 'Hồ Chí Minh',
    };

    console.log('search city: ', cityMap[normalized]);

    return cityMap[normalized] || name;
  }
}
