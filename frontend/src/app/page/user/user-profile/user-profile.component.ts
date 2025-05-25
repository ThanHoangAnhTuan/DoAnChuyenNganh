import { Component } from '@angular/core';
import { NavbarComponent } from '../../../components/navbar/navbar.component';
import { FormControl, ReactiveFormsModule } from '@angular/forms';
import { TuiInputModule } from '@taiga-ui/legacy';
import { TuiButton, TuiIcon } from '@taiga-ui/core';
import { NgIf } from '@angular/common';
import { TuiBadge } from '@taiga-ui/kit';

@Component({
    selector: 'app-user-profile',
    imports: [
        NavbarComponent,
        TuiInputModule,
        ReactiveFormsModule,
        TuiButton,
        NgIf,
        TuiBadge,
        TuiIcon
    ],
    templateUrl: './user-profile.component.html',
    standalone: true,
    styleUrl: './user-profile.component.scss',
})
export class UserProfileComponent {
    readonly name = new FormControl('Lưu Đình Quang Vinh');
    readonly displayName = new FormControl('Choose a display name');
    readonly email = new FormControl('vinh.ldq12464@sinhvien.hoasen.edu.vn');
    readonly phone = new FormControl('');
    readonly dob = new FormControl('');
    readonly nationality = new FormControl('');
    readonly gender = new FormControl('');
    readonly address = new FormControl('');
    readonly passport = new FormControl('Not provided');

    isEditing = false;
    currentEditingField: string | null = null;

    startEditing(field: string): void {
        this.isEditing = true;
        this.currentEditingField = field;
    }

    saveEdit(): void {
        this.isEditing = false;
        this.currentEditingField = null;
        // Add your save logic here
    }

    cancelEdit(): void {
        this.isEditing = false;
        this.currentEditingField = null;
    }
    getInitials(): string {
        const name = this.name.value || '';
        return name
            .split(' ')
            .map((n) => n[0])
            .join('')
            .toUpperCase()
            .substring(0, 2);
    }
}
