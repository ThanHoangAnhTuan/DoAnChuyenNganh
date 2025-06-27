import { Component } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { Router } from '@angular/router';
import { TuiIcon } from '@taiga-ui/core';
import { AuthService } from '../../../services/user/auth.service';
import {
    OTP,
    RegisterModel,
    RegisterResponse,
} from '../../../models/user/auth.model';
import { IsLoggedIn } from '../../../shared/token/token';
import { MessageService } from 'primeng/api';
import { Toast } from 'primeng/toast';
import { ButtonModule } from 'primeng/button';
import { Ripple } from 'primeng/ripple';

@Component({
    selector: 'app-register',
    imports: [FormsModule, TuiIcon, Toast, ButtonModule, Ripple],
    templateUrl: './register.component.html',
    styleUrl: './register.component.scss',
    providers: [MessageService],
})
export class RegisterComponent {
    email: string = '';

    constructor(
        private router: Router,
        private authService: AuthService,
        private messageService: MessageService
    ) {
        // Kiểm tra nếu đã đăng nhập thì chuyển hướng đến trang chính
        if (IsLoggedIn()) {
            this.router.navigate(['/']);
        }
    }
    showToast(
        severity: 'success' | 'info' | 'warn' | 'error',
        summary: string,
        detail: string
    ): void {
        this.messageService.add({ severity, summary, detail });
    }

    continueWithEmail() {
        if (this.email) {
            // Create an OTP object with the email
            const otpData: OTP = {
                verify_key: this.email,
                // We're sending the initial request, so no code is needed yet
                verify_code: '', // or remove this field if your OTP interface doesn't require it for sending
            };

            this.authService.verifyOTP(otpData).subscribe({
                next: (response) => {
                    this.showToast(
                        'success',
                        'OTP đã được gửi',
                        'Vui lòng kiểm tra email của bạn để xác nhận đăng ký.'
                    );
                    console.log('OTP sent successfully: ', response);
                    // Navigate to OTP verification page with email in query params
                    this.router.navigate(['/register/verify-otp'], {
                        queryParams: { email: this.email },
                    });
                },
                error: (error) => {
                    this.showToast(
                        'error',
                        'Lỗi gửi OTP',
                        error.message ||
                            'Đã xảy ra lỗi khi gửi OTP. Vui lòng thử lại sau.'
                    );
                    console.error('Error sending OTP:', error);
                    // Handle error (show error message to user)
                },
            });
        }
    }

    registerByEmail() {
        if (!this.email || this.email.trim() === '') {
            this.showToast(
                'warn',
                'Thông tin không hợp lệ',
                'Vui lòng nhập địa chỉ email của bạn.'
            );
            // console.warn('Please enter your email address.');
            return;
        }

        const newUser: RegisterModel = {
            verify_key: this.email,
            verify_type: 1,
            verify_purpose: 'TEST_USER',
        };

        console.log('Email gửi đi:', this.email);

        this.authService.registerUser(newUser).subscribe({
            next: (response) => {
                this.showToast(
                    'success',
                    'Đăng ký thành công',
                    'Vui lòng kiểm tra email của bạn để xác nhận đăng ký.'
                );
                console.log('Thông báo:', response.message);

                this.router.navigate(['/verify-otp'], {
                    queryParams: { email: this.email },
                });
            },
            error: (error) => {
                this.showToast(
                    'error',
                    'Lỗi đăng ký',
                    error.error?.message ||
                        'Đã xảy ra lỗi khi tạo người dùng. Vui lòng thử lại sau.'
                );
                console.error('Lỗi khi tạo người dùng:', error);
            },
        });
    }

    loginWithGoogle() {}

    loginWithApple() {}

    loginWithFacebook() {}
}
