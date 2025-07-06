import {
    TuiInputDateModule,
    TuiInputModule,
    TuiInputPhoneModule,
} from '@taiga-ui/legacy';
import { Component, OnInit } from '@angular/core';
import { UserService } from '../../../services/user/user.service';
import {
    Gender,
    UpdateUser,
    User,
    UserResponse,
} from '../../../models/user/user.model';
import { NavbarComponent } from '../../../components/navbar/navbar.component';
import {
    FormBuilder,
    FormControl,
    FormGroup,
    FormsModule,
    ReactiveFormsModule,
    Validators,
} from '@angular/forms';
import { CommonModule, NgIf } from '@angular/common';
import { finalize } from 'rxjs';
import { TuiButton, TuiLabel, TuiLoader } from '@taiga-ui/core';
import { TuiRadio } from '@taiga-ui/kit';
import { TuiTextfield } from '@taiga-ui/core';
import { DatePicker } from 'primeng/datepicker';
import { MessageService } from 'primeng/api';
import { Toast } from 'primeng/toast';
import { ButtonModule } from 'primeng/button';

@Component({
    selector: 'app-user-profile',
    imports: [
        NavbarComponent,
        ReactiveFormsModule,
        NgIf,
        TuiInputModule,
        TuiInputPhoneModule,
        TuiLoader,
        TuiRadio,
        TuiLabel,
        CommonModule,
        TuiButton,
        FormsModule,
        TuiInputDateModule,
        TuiTextfield,
        DatePicker,
        Toast,
        ButtonModule,
    ],
    templateUrl: './user-profile.component.html',
    standalone: true,
    styleUrl: './user-profile.component.scss',
    providers: [MessageService],
})
export class UserProfileComponent implements OnInit {
    profileForm!: FormGroup;
    Gender = Gender;

    currentUser: User = {
        // id: '',
        username: '',
        phone: '',
        gender: Gender.Male,
        birthday: '',
    };

    isLoading = false;
    isSubmitting = false;
    notification: { message: string; status: 'success' | 'error' } | null =
        null;

    constructor(
        private fb: FormBuilder,
        private userService: UserService,
        private messageService: MessageService
    ) {}
    showToast(
        severity: 'success' | 'info' | 'warn' | 'error',
        summary: string,
        detail: string
    ): void {
        this.messageService.add({ severity, summary, detail });
    }
    ngOnInit(): void {
        this.profileForm = this.fb.group({
            username: ['', Validators.required],
            phone: [''],
            gender: [Gender.Male],
            birthday: new FormControl(null, []),
        });
        this.loadUserData();
    }

    // loadUserData(): void {
    //     this.isLoading = true;
    //     this.userService
    //         .getUserInfo()
    //         .pipe(finalize(() => (this.isLoading = false)))
    //         .subscribe({
    //             next: (response: UserResponse) => {
    //                 this.currentUser = response.data;

    //                 console.log('Current User:', this.currentUser);

    //                 const [dayStr, monthStr, yearStr] =
    //                     this.currentUser.birthday.split('-'); // ["21", "06", "2025"]

    //                 const checkInDay = Number(dayStr); // 21
    //                 const checkInMonth = Number(monthStr); // 6
    //                 const checkInYear = Number(yearStr); // 2025

    //                 const birthdayDate = new Date(
    //                     checkInYear,
    //                     checkInMonth - 1,
    //                     checkInDay
    //                 );

    //                 this.profileForm.patchValue({
    //                     username: this.currentUser.username,
    //                     phone: this.currentUser.phone,
    //                     gender: this.currentUser.gender,
    //                     birthday: birthdayDate,
    //                 });
    //             },
    //             error: (error) => {
    //                 this.showToast(
    //                     'error',
    //                     'Lỗi tải thông tin người dùng',
    //                     error.message ||
    //                         'Đã xảy ra lỗi khi tải thông tin người dùng. Vui lòng thử lại sau.'
    //                 );
    //                 // console.error('Error loading user data:', error);
    //             },
    //         });
    // }
    loadUserData(): void {
        this.isLoading = true;
        this.userService
            .getUserInfo()
            .pipe(finalize(() => (this.isLoading = false)))
            .subscribe({
                next: (response: UserResponse) => {
                    this.currentUser = response.data;
                    console.log('Current User:', this.currentUser);
                    let birthdayDate: Date;
                    if (this.currentUser.birthday) {
                        // parse từ chuỗi "21-06-2025"
                        const [dayStr, monthStr, yearStr] =
                            this.currentUser.birthday.split('-');
                        const checkInDay = Number(dayStr);
                        const checkInMonth = Number(monthStr);
                        const checkInYear = Number(yearStr);
                        birthdayDate = new Date(
                            checkInYear,
                            checkInMonth - 1,
                            checkInDay
                        );
                    } else {
                        // Không có ngày sinh => dùng ngày hiện tại
                        birthdayDate = new Date();
                    }
                    this.profileForm.patchValue({
                        username: this.currentUser.username,
                        phone: this.currentUser.phone,
                        gender: this.currentUser.gender,
                        birthday: birthdayDate,
                    });
                },
                error: (error) => {
                    this.showToast(
                        'error',
                        'Lỗi tải thông tin người dùng',
                        error.message ||
                            'Đã xảy ra lỗi khi tải thông tin người dùng. Vui lòng thử lại sau.'
                    );
                },
            });
    }

    updateUserProfile(): void {
        if (this.profileForm.invalid) {
            this.showToast(
                'error',
                'Thông tin không hợp lệ',
                'Vui lòng kiểm tra lại thông tin.'
            );
            this.profileForm.markAllAsTouched();
            return;
        }
        this.isSubmitting = true;
        const day = String(this.profileForm.value?.birthday.getDate()).padStart(
            2,
            '0'
        ); // "21"
        const month = String(
            this.profileForm.value?.birthday.getMonth() + 1
        ).padStart(2, '0'); // "06" (vì tháng JS bắt đầu từ 0)
        const year = this.profileForm.value?.birthday.getFullYear(); // 2025

        const formattedBirthday = `${day}-${month}-${year}`; // "21-06-2025"

        const userData: UpdateUser = {
            // id: this.currentUser.id,
            username: this.profileForm.value.username,
            phone: this.profileForm.value.phone,
            gender: this.profileForm.value.gender === 'male' ? 0 : 1,
            birthday: formattedBirthday,
        };

        this.userService
            .updateUserInfo(userData)
            .pipe(finalize(() => (this.isSubmitting = false)))
            .subscribe({
                next: (response) => {
                    this.currentUser = response.data;
                    this.showToast(
                        'success',
                        'Cập nhật thành công',
                        'Thông tin cá nhân đã được cập nhật thành công.'
                    );
                },
                error: (error) => {
                    this.showToast(
                        'error',
                        'Cập nhật thất bại',
                        error.error?.message ||
                            'Đã xảy ra lỗi khi cập nhật thông tin. Vui lòng thử lại sau.'
                    );
                    // console.error('Update failed:', error);
                },
            });
    }
}
