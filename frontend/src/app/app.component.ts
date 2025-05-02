import { TuiRoot } from "@taiga-ui/core";
import { ChangeDetectionStrategy, Component, OnInit, ViewChild } from '@angular/core';
import { Accommodation, CreateAccommodation, UpdateAccommodation } from "./models/accommodation.model";
import { RouterOutlet } from '@angular/router';
import { FormControl, FormGroup } from "@angular/forms";
import { AccommodationService } from "./services/accommodation.service";
import {NavbarComponent} from './components/navbar/navbar.component';
import {HomepageComponent} from './pages/homepage/homepage.component';
import SearchBoxComponent from './components/search-box/search-box.component';

@Component({
  selector: 'app-root',
  imports: [RouterOutlet, TuiRoot, NavbarComponent, SearchBoxComponent, HomepageComponent],
  templateUrl: './app.component.html',
  styleUrl: './app.component.scss',
})
export class AppComponent {}
