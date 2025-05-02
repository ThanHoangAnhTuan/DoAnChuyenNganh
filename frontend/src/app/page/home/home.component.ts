import { Component } from '@angular/core';
import SearchBoxComponent from '../../components/search-box/search-box.component';
import { NavbarComponent } from '../../components/navbar/navbar.component';

@Component({
  selector: 'app-home',
  imports: [SearchBoxComponent, NavbarComponent],
  templateUrl: './home.component.html',
  styleUrl: './home.component.scss'
})
export class HomeComponent {
    
}
