import { Component, OnInit } from '@angular/core';
import { Router, RouterLink, RouterLinkActive } from '@angular/router';
import { TuiDataList, TuiDropdown, TuiIcon } from '@taiga-ui/core';
import { TuiAvatar, TuiChevron } from '@taiga-ui/kit';
import { GetUserRole } from '../../shared/token/token';

@Component({
    selector: 'app-navbar',
    imports: [
        RouterLink,
        RouterLinkActive,
        TuiChevron,
        TuiDataList,
        TuiDropdown,
        TuiIcon,
        TuiAvatar,
    ],
    standalone: true,
    templateUrl: './navbar.component.html',
    styleUrl: './navbar.component.scss',
})
export class NavbarComponent implements OnInit {
    isUserLoggedIn = false;
    isManagerLoggedIn = false;
    isAdminLoggedIn = false;
    userName: string | null = null;
    userAvatar: string | null = null;

    protected readonly groups = [
        {
            label: '',
            items: [
                {
                    label: 'My account',
                    routerLink: '/user-profile',
                    icon: '@tui.user',
                },
                {
                    label: 'Bookings & Trips',
                    routerLink: '/',
                    icon: '@tui.backpack',
                },
            ],
        },
    ];
    constructor(private router: Router) {}

    ngOnInit(): void {
        this.checkLoginStatus();
    }

    checkLoginStatus(): void {
        const user = GetUserRole();
        if (user === 'user') {
            this.isUserLoggedIn = true;
        } else if (user === 'manager') {
            this.isManagerLoggedIn = true;
        } else if (user === 'admin') {
            this.isAdminLoggedIn = true;
        }
    }

    logout(): void {
        // Lưu trạng thái trước khi reset
        const wasUserLoggedIn = this.isUserLoggedIn;
        const wasManagerLoggedIn = this.isManagerLoggedIn;
        const wasAdminLoggedIn = this.isAdminLoggedIn;

        // Xóa token/cookie nếu có
        document.cookie =
            'auth_token=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;';
        localStorage.removeItem('token');

        // Update UI
        this.isUserLoggedIn = false;
        this.isManagerLoggedIn = false;
        this.isAdminLoggedIn = false;
        this.userName = null;
        this.userAvatar = null;

        // Navigate dựa vào trạng thái đã lưu
        if (wasManagerLoggedIn) {
            this.router.navigate(['/manager/login']);
        } else if (wasUserLoggedIn) {
            this.router.navigate(['/']);
        } else if (wasAdminLoggedIn) {
            this.router.navigate(['/admin/login']);
        }
    }
}
