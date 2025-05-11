import {Routes} from '@angular/router';
import {ManagerComponent} from './page/manager/manager.component';
import {HomeComponent} from './page/home/home.component';
import {AccommodationComponent} from './page/accommodation/accommodation.component';
import {AccommodationDetailComponent} from './page/accommodation-detail/accommodation-detail.component';
import {SearchPageComponent} from './page/search-page/search-page.component';

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
        path: 'search/:city',
        component: SearchPageComponent,
    },
    {
        path: '**',
        redirectTo: '',
        pathMatch: 'full',
    },
];
