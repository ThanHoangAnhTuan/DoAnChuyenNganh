import { Component } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { Router } from '@angular/router';
import { TuiIcon } from '@taiga-ui/core';
import { UserService } from '../../../services/user/user.service';
import {
    OTP,
    RegisterModel,
    RegisterResponse,
} from '../../../models/user/user.model';

@Component({
    selector: 'app-register',
    imports: [FormsModule, TuiIcon],
    templateUrl: './register.component.html',
    styleUrl: './register.component.scss',
})
export class RegisterComponent {
    email: string = '';

    constructor(private router: Router, private userService: UserService) {
        // Kiểm tra nếu đã đăng nhập thì chuyển hướng đến trang chính
        if (this.userService.isLoggedIn()) {
            this.router.navigate(['/']);
        }
    }

    continueWithEmail() {
        if (this.email) {
            // Create an OTP object with the email
            const otpData: OTP = {
                verify_key: this.email,
                // We're sending the initial request, so no code is needed yet
                verify_code: '', // or remove this field if your OTP interface doesn't require it for sending
            };

            this.userService.verifyOTP(otpData).subscribe({
                next: (response) => {
                    console.log('OTP sent successfully: ', response);
                    // Navigate to OTP verification page with email in query params
                    this.router.navigate(['/register/verify-otp'], {
                        queryParams: { email: this.email },
                    });
                },
                error: (error) => {
                    console.error('Error sending OTP:', error);
                    // Handle error (show error message to user)
                },
            });
        }
    }

    registerByEmail() {
        if (!this.email || this.email.trim() === '') {
            console.warn('Please enter your email address.');
            return;
        }

        const newUser: RegisterModel = {
            verify_key: this.email,
            verify_type: 1,
            verify_purpose: 'TEST_USER',
        };

        console.log('Email gửi đi:', this.email);

        this.userService.registerUser(newUser).subscribe({
            next: (response) => {
                console.log('Thông báo:', response.message);

                this.router.navigate(['/verify-otp'], {
                    queryParams: { email: this.email },
                });
            },
            error: (error) => {
                console.error('Lỗi khi tạo người dùng:', error);
            },
        });
    }

    loginWithGoogle() {}

    loginWithApple() {}

    loginWithFacebook() {}
}
