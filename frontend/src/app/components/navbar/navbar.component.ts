import { Component, OnInit } from '@angular/core';
import { Router, RouterLink, RouterLinkActive } from '@angular/router';
import { TuiDataList, TuiDropdown, TuiIcon } from '@taiga-ui/core';
import { TuiAvatar, TuiChevron } from '@taiga-ui/kit';
import { UserService } from '../../services/user/user.service';
import { GetToken, GetUserRole } from '../../shared/token/token';

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
                    label: 'Saved',
                    routerLink: '/',
                    icon: '@tui.heart',
                },
                {
                    label: 'Reviews',
                    routerLink: '/',
                    icon: '@tui.message-circle-code',
                },
                {
                    label: 'Bookings & Trips',
                    routerLink: '/',
                    icon: '@tui.backpack',
                },
                // {
                //     label: 'Sign out',
                //     routerLink: '/',
                //     icon: '@tui.log-out',
                // },
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
        }
    }

    // loadUserInfo(): void {
    //     // Add this method to your UserService to get user profile info
    //     this.userService.getUserProfile().subscribe({
    //         next: (user) => {
    //             this.userName = user.name || user.email;
    //             this.userAvatar = user.avatar;
    //         },
    //         error: (error) => {
    //             console.error('Error loading user profile:', error);
    //             // Handle token expiration or other auth errors
    //             if (error.status === 401) {
    //                 this.logout();
    //             }
    //         },
    //     });
    // }

    logout(): void {
        // Lưu trạng thái trước khi reset
        const wasUserLoggedIn = this.isUserLoggedIn;
        const wasManagerLoggedIn = this.isManagerLoggedIn;

        // Xóa token/cookie nếu có
        document.cookie =
            'auth_token=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;';
        localStorage.removeItem('token');
        sessionStorage.removeItem('token');

        // Update UI
        this.isUserLoggedIn = false;
        this.isManagerLoggedIn = false;
        this.userName = null;
        this.userAvatar = null;

        // Navigate dựa vào trạng thái đã lưu
        if (wasManagerLoggedIn) {
            this.router.navigate(['/manager/login']);
        } else {
            this.router.navigate(['/']);
        }
    }
}
