import { Routes } from '@angular/router';
import { ManagerComponent } from './page/manager/manager.component';
import { HomeComponent } from './page/home/home.component';
import { AccommodationComponent } from './page/accommodation/accommodation.component';
import { AccommodationDetailComponent } from './page/accommodation-detail/accommodation-detail.component';

export const routes: Routes = [
    {
        path: '',
        component: HomeComponent,
    },
    {
        path: 'manager',
        component: ManagerComponent,
    },
    {
        path: 'manager/accommodation',
        component: AccommodationComponent,
    },
    {
        path: 'manager/accommodation/detail',
        component: AccommodationDetailComponent,
    },
    {
        path: 'search/accommodation/detail/:name',
        component: AccommodationDetailComponent,
    },
    {
        path: '**',
        redirectTo: '',
        pathMatch: 'full',
    },
];
