import { NgForOf } from '@angular/common';
import { Component, OnInit } from '@angular/core';
import { Router, RouterLink, RouterLinkActive } from '@angular/router';
import { TuiDataList, TuiDropdown, TuiIcon } from '@taiga-ui/core';
import { TuiAvatar, TuiChevron } from '@taiga-ui/kit';
import { UserService } from '../../services/user/user.service';

@Component({
    selector: 'app-navbar',
    imports: [
        NgForOf,
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
    isLoggedIn = false;
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
    constructor(private userService: UserService, private router: Router) {}

    ngOnInit(): void {
        this.checkLoginStatus();
    }

    checkLoginStatus(): void {
        // You'll need to add an isLoggedIn method to your UserService
        const token =
            localStorage.getItem('token') || sessionStorage.getItem('token');
        this.isLoggedIn = !!token;

        if (this.isLoggedIn) {
            // this.loadUserInfo();
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
        // Clear authentication data
        localStorage.removeItem('token');
        sessionStorage.removeItem('token');

        // Update UI
        this.isLoggedIn = false;
        this.userName = null;
        this.userAvatar = null;

        // Navigate to home page
        this.router.navigate(['/']);
    }
}
