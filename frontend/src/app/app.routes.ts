import { Routes } from '@angular/router';
import { ManagerComponent } from './page/manager/manager.component';
import { HomeComponent } from './page/home/home.component';
import { AccommodationComponent } from './page/accommodation/accommodation.component';
import { AccommodationDetailComponent } from './page/accommodation-detail/accommodation-detail.component';
import { LoginComponent } from './page/login/login.component';
import { RoleGuard } from './shared/guards/role.guard';
import { MediaLibraryComponent } from './page/media-library/media-library.component';

export const routes: Routes = [
    {
        path: '',
        component: HomeComponent,
    },
    {
        path: 'manager',
        component: ManagerComponent,
        canActivate: [RoleGuard],
        data: { expectedRole: 'manager' },
    },
    {
        path: 'manager/login',
        component: LoginComponent,
    },
    {
        path: 'manager/accommodation',
        component: AccommodationComponent,
        canActivate: [RoleGuard],
        data: { expectedRole: 'manager' },
    },
    {
        path: 'manager/accommodation/:id/details',
        component: AccommodationDetailComponent,
        canActivate: [RoleGuard],
        data: { expectedRole: 'manager' },
    },
    {
        path: 'manager/accommodation/:id/images',
        component: MediaLibraryComponent,
        canActivate: [RoleGuard],
        data: { expectedRole: 'manager' },
    },
    {
        path: 'manager/accommodation/detail/:id/images',
        component: MediaLibraryComponent,
        canActivate: [RoleGuard],
        data: { expectedRole: 'manager' },
    },
    {
        path: '**',
        redirectTo: '',
        pathMatch: 'full',
    },
];
