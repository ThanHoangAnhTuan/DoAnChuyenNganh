import { TuiRoot } from '@taiga-ui/core';
import { Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import {TuiButton} from '@taiga-ui/core';
import { Accommodation, CreateAccommodation, UpdateAccommodation } from "./models/accommodation.model";
import { FormControl, FormGroup } from "@angular/forms";
import { AccommodationService } from "./services/accommodation.service";
import { HomePageComponent } from "./views/home-page/home-page.component";

@Component({
    selector: 'app-root',
    imports: [RouterOutlet, TuiRoot],
    templateUrl: './app.component.html',
    styleUrl: './app.component.scss',
})
export class AppComponent {}
