import { Component, OnInit } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { OTP, UpdatePassword } from '../../../models/user/auth.model';
import { interval, Subscription, take } from 'rxjs';
import { CommonModule } from '@angular/common';
import { AuthService } from '../../../services/user/auth.service';
import { MessageService } from 'primeng/api';
import { Toast } from 'primeng/toast';
import { ButtonModule } from 'primeng/button';
import { Ripple } from 'primeng/ripple';

@Component({
    selector: 'app-verify-otp',
    imports: [FormsModule, CommonModule, Toast, ButtonModule, Ripple],
    templateUrl: './verify-otp.component.html',
    styleUrl: './verify-otp.component.scss',
    providers: [MessageService],
})
export class VerifyOtpComponent implements OnInit {
    otp: string = '';
    password: string = '';
    token: string = '';
    email: string = '';
    step = 1; // 1: OTP verification, 2: Password update
    confirmPassword = '';
    showPassword = false;
    showConfirmPassword = false;
    isVerifying = false;
    isUpdating = false;
    resendCountdown = 0;
    username: string = '';
    phone: string = '';
    gender: number = 0; // 0 = Male, 1 = Female
    birthday: string = ''; // ISO format (yyyy-mm-dd)
    goToStep(step: number) {
        this.step = step;
    }
    // Mảng chứa giá trị OTP từ các ô input
    otpValues: string[] = ['', '', '', '', '', ''];
    countdownInterval: any;
    // Store token after OTP verification
    verificationToken: string = '';

    // Error message
    errorMessage: string = '';

    private countdownSub: Subscription | null = null;

    constructor(
        private route: ActivatedRoute,
        private authService: AuthService,
        private router: Router,
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
        // Lấy email từ query params
        this.route.queryParams.subscribe((params) => {
            this.email =
                params['email'] || localStorage.getItem('resetEmail') || '';
            if (!this.email) {
                this.showToast(
                    'error',
                    'Lỗi xác thực',
                    'Không tìm thấy email để xác thực. Vui lòng thử lại.'
                );
                // Redirect to forgot password if no email found
                this.router.navigate(['/']);
                return;
            }
        });
    }
    ngOnDestroy() {
        if (this.countdownSub) {
            this.countdownSub.unsubscribe();
        }
    }

    onOtpInput(event: any, index: number, inputElement: HTMLInputElement) {
        const value = event.target.value;

        // Store the value
        this.otpValues[index] = value;

        // Clear error message when user types
        this.errorMessage = '';

        // Auto-focus next input if value is entered
        if (value.length === 1 && index < 5) {
            const nextInput = inputElement.parentElement?.querySelector(
                `input:nth-child(${index + 2})`
            ) as HTMLInputElement;
            if (nextInput) nextInput.focus();
        }
    }
    onOtpPaste(event: ClipboardEvent) {
        event.preventDefault();
        if (!event.clipboardData) return;

        const pastedText = event.clipboardData.getData('text');
        if (!pastedText) return;

        const otpInputs = document.querySelectorAll(
            '.otp-input'
        ) as NodeListOf<HTMLInputElement>;

        // Điền các ký tự vào các ô input
        for (
            let i = 0;
            i < Math.min(otpInputs.length, pastedText.length);
            i++
        ) {
            if (/^\d+$/.test(pastedText[i])) {
                otpInputs[i].value = pastedText[i];
            }
        }

        // Focus vào ô cuối cùng hoặc ô tiếp theo
        const lastFilledIndex = Math.min(
            otpInputs.length - 1,
            pastedText.length - 1
        );
        if (lastFilledIndex >= 0) {
            otpInputs[lastFilledIndex].focus();
        }

        this.updateOTPValue();
    }

    private updateOTPValue() {
        const otpInputs = document.querySelectorAll(
            '.otp-input'
        ) as NodeListOf<HTMLInputElement>;
        this.otp = Array.from(otpInputs)
            .map((input) => input.value)
            .join('');
    }
    resendOTP() {
        // Giả lập gọi API gửi lại OTP
        this.resendCountdown = 60;

        this.countdownSub = interval(1000)
            .pipe(take(60))
            .subscribe(() => {
                this.resendCountdown--;
                if (this.resendCountdown === 0 && this.countdownSub) {
                    this.countdownSub.unsubscribe();
                }
            });
    }

    togglePassword() {
        this.showPassword = !this.showPassword;
    }

    toggleConfirmPassword() {
        this.showConfirmPassword = !this.showConfirmPassword;
    }

    getPasswordStrengthClass() {
        if (!this.password) return '';

        const hasLetter = /[a-zA-Z]/.test(this.password);
        const hasNumber = /\d/.test(this.password);
        const hasSpecial = /[!@#$%^&*(),.?":{}|<>]/.test(this.password);

        if (this.password.length < 6) return 'weak';
        if (this.password.length >= 8 && hasLetter && hasNumber && hasSpecial)
            return 'strong';
        return 'medium';
    }

    getPasswordStrengthText() {
        const strength = this.getPasswordStrengthClass();
        switch (strength) {
            case 'weak':
                return 'Yếu - Mật khẩu quá ngắn';
            case 'medium':
                return 'Trung bình - Thêm ký tự đặc biệt';
            case 'strong':
                return 'Mạnh - Mật khẩu an toàn';
            default:
                return '';
        }
    }

    canUpdatePassword() {
        return (
            this.password &&
            this.confirmPassword &&
            this.password === this.confirmPassword &&
            this.getPasswordStrengthClass() !== 'weak'
        );
    }

    canUpdateUserInfo() {
        return (
            this.username &&
            this.phone &&
            this.gender !== null &&
            this.birthday &&
            this.birthday.trim() !== ''
        );
    }

    verifyOTP() {
        // Combine OTP values
        const otpCode = this.otpValues.join('');

        // Validate OTP length
        if (otpCode.length !== 6) {
            this.showToast(
                'warn',
                'Thông tin không hợp lệ',
                'Vui lòng nhập đầy đủ mã OTP 6 số.'
            );
            // this.errorMessage = 'Vui lòng nhập đầy đủ mã OTP 6 số';
            return;
        }

        this.isVerifying = true;
        // this.errorMessage = '';

        // Create otpData object
        const otpData: OTP = {
            verify_key: this.email,
            verify_code: otpCode,
        };

        // Call your authentication service to verify the OTP
        this.authService.verifyOTP(this.email, otpCode, otpData).subscribe({
            next: (response) => {
                // Store verification token if your API returns one
                if (response.data?.token) {
                    this.verificationToken = response.data.token;
                    localStorage.setItem('resetToken', this.verificationToken);
                }

                // Success - move to step 2
                this.step = 2;
                this.showToast(
                    'success',
                    'Xác thực thành công',
                    'Mã OTP đã được xác thực thành công. Vui lòng cập nhật mật khẩu mới.'
                );
                this.isVerifying = false;
            },
            error: (error) => {
                this.showToast(
                    'error',
                    'Lỗi xác thực OTP',
                    error.error?.message ||
                        'Mã OTP không đúng. Vui lòng thử lại.'
                );
                // console.error('OTP verification error:', error);
                this.isVerifying = false;
            },
        });
    }

    updatePassword() {
        if (!this.password || this.password.trim() === '') {
            this.showToast(
                'warn',
                'Thông tin không hợp lệ',
                'Vui lòng nhập mật khẩu mới.'
            );
            return;
        }

        // Password validation
        if (this.password !== this.confirmPassword) {
            this.showToast(
                'warn',
                'Thông tin không hợp lệ',
                'Mật khẩu xác nhận không khớp. Vui lòng kiểm tra lại.'
            );
            return;
        }

        // Get token from the verification response or localStorage
        const token =
            this.verificationToken || localStorage.getItem('resetToken') || '';

        if (!token) {
            this.showToast(
                'error',
                'Lỗi xác thực',
                'Không tìm thấy token xác thực. Vui lòng thử lại.'
            );
            return;
        }

        this.isUpdating = true;

        const passwordData: UpdatePassword = {
            token: token, // Use the correct token from verification
            password: this.password,
        };

        this.authService.updatePassword(passwordData).subscribe({
            next: (response) => {
                this.showToast(
                    'success',
                    'Cập nhật mật khẩu thành công',
                    'Mật khẩu của bạn đã được cập nhật thành công.'
                );
                // Clear localStorage items that are no longer needed
                localStorage.removeItem('resetToken');
                localStorage.removeItem('resetEmail');
                // Navigate to login page or show success message
                this.router.navigate(['/login'], {
                    queryParams: { passwordUpdated: 'success' },
                });
            },
            error: (error) => {
                this.showToast(
                    'error',
                    'Lỗi cập nhật mật khẩu',
                    error.error?.message ||
                        'Đã xảy ra lỗi khi cập nhật mật khẩu. Vui lòng thử lại.'
                );
                console.error('Error updating password:', error);
                this.isUpdating = false;
            },
            complete: () => {
                this.isUpdating = false;
            },
        });
    }
}
