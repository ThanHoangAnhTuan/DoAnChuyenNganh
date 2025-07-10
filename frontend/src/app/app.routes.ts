import { Routes } from '@angular/router';
import { HomeComponent } from './page/user/home/home.component';
import { AccommodationComponent } from './page/manager/accommodation/accommodation.component';
import { AccommodationDetailComponent } from './page/user/accommodation-detail/accommodation-detail.component';
import { AccommodationDetailComponent as ManagerAccommodationDetailComponent } from './page/manager/accommodation-detail/accommodation-detail.component';
import { LoginComponent as ManagerLoginComponent } from './page/manager/login/login.component';
import { LoginComponent as AdminLoginComponent } from './page/admin/login/login.component';
import { RoleGuard } from './shared/guards/role.guard';
import { MediaLibraryComponent } from './page/manager/media-library/media-library.component';
import { SearchPageComponent } from './page/user/search-page/search-page.component';
import { UserProfileComponent } from './page/user/user-profile/user-profile.component';
import { FacilityComponent } from './page/admin/facility/facility.component';
import { RegisterComponent } from './page/user/register/register.component';
import { LoginComponent } from './page/user/login/login.component';
import { VerifyOtpComponent } from './page/user/verify-otp/verify-otp.component';
import { PaymentComponent } from './page/payment/payment.component';
import { FacilityDetailComponent } from './page/admin/facility-detail/facility-detail.component';
import { StatsComponent } from './page/manager/stats/stats.component';
import { ManagerComponent } from './page/admin/manager/manager.component';
import { RoomComponent } from './page/manager/room/room.component';
import { AccommodationComponent as AdminAccommodationComponent } from './page/admin/accommodation/accommodation.component';
import { OrderComponent as ManagerOrderComponent} from './page/manager/order/order.component';

export const routes: Routes = [
    {
        path: '',
        component: HomeComponent,
    },
    {
        path: 'manager/login',
        component: ManagerLoginComponent,
    },
    {
        path: 'manager/accommodation',
        component: AccommodationComponent,
        canActivate: [RoleGuard],
        data: { expectedRole: 'manager' },
    },
    {
        path: 'manager/order',
        component: ManagerOrderComponent,
        canActivate: [RoleGuard],
        data: { expectedRole: 'manager' },
    },
    {
        path: 'manager/accommodation/:id/details',
        component: ManagerAccommodationDetailComponent,
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
        path: 'manager/accommodation/stats',
        component: StatsComponent,
        canActivate: [RoleGuard],
        data: { expectedRole: 'manager' },
    },
    {
        path: 'manager/accommodation/detail/:id/rooms',
        component: RoomComponent,
        canActivate: [RoleGuard],
        data: { expectedRole: 'manager' },
    },
    {
        path: 'admin/login',
        component: AdminLoginComponent,
    },
    {
        path: 'admin/facility',
        component: FacilityComponent,
        canActivate: [RoleGuard],
        data: { expectedRole: 'admin' },
    },
    {
        path: 'admin/manager',
        component: ManagerComponent,
        canActivate: [RoleGuard],
        data: { expectedRole: 'admin' },
    },
    {
        path: 'admin/manager/:id/accommodations',
        component: AdminAccommodationComponent,
        canActivate: [RoleGuard],
        data: { expectedRole: 'admin' },
    },
    {
        path: 'admin/facility-detail',
        component: FacilityDetailComponent,
        canActivate: [RoleGuard],
        data: { expectedRole: 'admin' },
    },
    {
        path: 'search/accommodation/detail/:id',
        component: AccommodationDetailComponent,
    },
    {
        path: 'search/:city',
        component: SearchPageComponent,
    },
    {
        path: 'register',
        component: RegisterComponent,
    },
    {
        path: 'login',
        component: LoginComponent,
    },
    {
        path: 'user-profile',
        component: UserProfileComponent,
    },
    {
        path: 'verify-otp',
        component: VerifyOtpComponent,
    },
    {
        path: 'payment/result',
        component: PaymentComponent,
    },
    {
        path: '**',
        redirectTo: '',
        pathMatch: 'full',
    },
];
