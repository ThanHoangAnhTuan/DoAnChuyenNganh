import { Component, OnInit } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { OTP, UpdatePassword } from '../../../models/user/auth.model';
import { interval, Subscription, take } from 'rxjs';
import { CommonModule } from '@angular/common';
import { UserService } from '../../../services/user/user.service';
import { AuthService } from '../../../services/user/auth.service';

@Component({
    selector: 'app-verify-otp',
    imports: [FormsModule, CommonModule],
    templateUrl: './verify-otp.component.html',
    styleUrl: './verify-otp.component.scss',
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
    private countdownSub: Subscription | null = null;

    constructor(
        private route: ActivatedRoute,
        private userService: UserService,
        private authService: AuthService,
        private router: Router
    ) {}

    ngOnInit(): void {
        // Lấy email từ query params
        this.route.queryParams.subscribe((params) => {
            this.email = params['email'] || '';
            console.log('Email from query param:', this.email);
        });
    }
    ngOnDestroy() {
        if (this.countdownSub) {
            this.countdownSub.unsubscribe();
        }
    }
    onOtpInput(event: any, index: number, inputElement: HTMLInputElement) {
        // Lấy giá trị đã nhập
        const value = inputElement.value;

        // Auto focus tới ô tiếp theo nếu đã nhập
        if (value.length === 1 && index < 5) {
            const nextInput = inputElement.parentElement?.children[
                index + 1
            ] as HTMLInputElement;
            if (nextInput) {
                nextInput.focus();
            }
        }

        // Quay lại ô trước nếu xóa
        if (value.length === 0 && event.key === 'Backspace' && index > 0) {
            const prevInput = inputElement.parentElement?.children[
                index - 1
            ] as HTMLInputElement;
            if (prevInput) {
                prevInput.focus();
            }
        }

        // Cập nhật giá trị OTP tổng
        setTimeout(() => {
            this.updateOTPValue();
        });
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
        if (!this.otp || this.otp.trim() === '') {
            console.warn('Please enter otp.');
            return;
        }
        this.isVerifying = true;
        // Giả lập gọi API xác thực OTP
        setTimeout(() => {
            this.isVerifying = false;
            this.step = 2; // Chuyển sang bước cập nhật mật khẩu
        }, 1500);

        const otpData: OTP = {
            verify_key: this.email,
            verify_code: this.otp,
        };

        if (this.otp && this.email) {
            this.authService.verifyOTP(otpData).subscribe({
                next: (response) => {
                    console.log('OTP verified successfully:', response);
                    this.token = response.data.token; // Assuming the token is returned in the response
                    // Navigate to success page or show success message
                },
                error: (error) => {
                    console.error('Error verifying OTP:', error);
                    // Handle error (show error message to user)
                },
            });
        }
    }

    updatePassword() {
        if (!this.password || this.password.trim() === '') {
            console.warn('Please enter your new password.');
            return;
        }
        const passwordData: UpdatePassword = {
            token: this.token,
            password: this.password,
        };

        console.log('Updating password with token:', this.token);
        console.log('New password:', this.password);

        this.authService.updatePassword(passwordData).subscribe({
            next: (response) => {
                console.log('Password updated successfully:', response);
                // Navigate to success page or show success message
                this.router.navigate(['/login']);
            },
            error: (error) => {
                console.error('Error updating password:', error);
                // Handle error (show error message to user)
            },
        });
    }
}
