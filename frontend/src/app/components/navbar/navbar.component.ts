import { NgForOf } from '@angular/common';
import { Component } from '@angular/core';
import { RouterLink, RouterLinkActive } from '@angular/router';
import { TuiDataList, TuiDropdown, TuiIcon } from '@taiga-ui/core';
import { TuiChevron } from '@taiga-ui/kit';

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
    ],
    standalone: true,
    templateUrl: './navbar.component.html',
    styleUrl: './navbar.component.scss',
})
export class NavbarComponent {
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
                {
                    label: 'Sign out',
                    routerLink: '/',
                    icon: '@tui.log-out',
                },
            ],
        },
    ];
}
