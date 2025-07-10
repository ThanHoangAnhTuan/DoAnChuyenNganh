import { Component, OnInit } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { OTP, UpdatePassword } from '../../../models/user/auth.model';
import { finalize, interval, Subscription, take } from 'rxjs';
import { CommonModule } from '@angular/common';
import { AuthService } from '../../../services/user/auth.service';
import { MessageService } from 'primeng/api';
import { Toast } from 'primeng/toast';
import { ButtonModule } from 'primeng/button';
import { LoaderComponent } from '../../../components/loader/loader.component';
import { InputOtpModule } from 'primeng/inputotp';
import { TuiInputNumber } from '@taiga-ui/kit';
import { TuiTextfield } from '@taiga-ui/core';

@Component({
    selector: 'app-verify-otp',
    imports: [
        FormsModule,
        CommonModule,
        Toast,
        ButtonModule,
        LoaderComponent,
        InputOtpModule,
        TuiInputNumber,
        TuiTextfield,
    ],
    templateUrl: './verify-otp.component.html',
    styleUrl: './verify-otp.component.scss',
    providers: [MessageService],
})
export class VerifyOtpComponent implements OnInit {
    isLoading: boolean = false;
    otp: string = '';
    password: string = '';
    token: string = '';
    email: string = '';
    step = 1; // 1: OTP verification, 2: Password update
    confirmPassword = '';
    showPassword = false;
    showConfirmPassword = false;
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
        this.route.queryParams
            .pipe(finalize(() => (this.isLoading = false)))
            .subscribe({
                next: (params) => {
                    this.email =
                        params['email'] ||
                        localStorage.getItem('resetEmail') ||
                        '';
                    if (!this.email) {
                        this.showToast(
                            'error',
                            'Lỗi xác thực',
                            'Không tìm thấy email để xác thực. Vui lòng thử lại.'
                        );
                        // Redirect to forgot password if no email found
                        this.router.navigate(['/']);
                    }
                },
                error: (error) => {
                    console.error('Error loading params:', error);
                    this.showToast(
                        'error',
                        'Lỗi kết nối',
                        'Không thể kết nối đến máy chủ. Vui lòng thử lại sau.'
                    );
                },
            });
    }
    ngOnDestroy() {
        if (this.countdownSub) {
            this.countdownSub.unsubscribe();
        }
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

        this.isLoading = true;
        // this.errorMessage = '';

        // Create otpData object
        const otpData: OTP = {
            verify_key: this.email,
            verify_code: otpCode,
        };

        // Call your authentication service to verify the OTP
        this.authService
            .verifyOTP(this.email, otpCode, otpData)
            .pipe(
                finalize(() => {
                    this.isLoading = false;
                })
            )
            .subscribe({
                next: (response) => {
                    // Store verification token if your API returns one
                    if (response.data?.token) {
                        this.verificationToken = response.data.token;
                        localStorage.setItem(
                            'resetToken',
                            this.verificationToken
                        );
                    }

                    // Success - move to step 2
                    this.step = 2;
                    this.showToast(
                        'success',
                        'Xác thực thành công',
                        'Mã OTP đã được xác thực thành công. Vui lòng cập nhật mật khẩu mới.'
                    );
                },
                error: (error) => {
                    this.showToast(
                        'error',
                        'Lỗi xác thực OTP',
                        error.error?.message ||
                            'Mã OTP không đúng. Vui lòng thử lại.'
                    );
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

        this.isLoading = true;

        const passwordData: UpdatePassword = {
            token: token, // Use the correct token from verification
            password: this.password,
        };

        this.authService
            .updatePassword(passwordData)
            .pipe(
                finalize(() => {
                    this.isLoading = false;
                })
            )
            .subscribe({
                next: (response) => {
                    // Clear localStorage items that are no longer needed
                    localStorage.removeItem('resetToken');
                    localStorage.removeItem('resetEmail');

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
                },
                complete: () => {
                    this.isLoading = false;
                },
            });
    }
    // In verify-otp.component.ts

    // Update the method to handle changes in the Taiga UI inputs
    onOtpChange(index: number, value: number | null): void {
        // Store the value (convert null to empty string)
        this.otpValues[index] = value === null ? '' : String(value);

        // Clear error message when user types
        this.errorMessage = '';

        // Auto-focus next input if value is entered and it's not the last input
        if (value !== null && value !== undefined && index < 5) {
            const inputs = document.querySelectorAll(
                '.tui-otp-input'
            ) as NodeListOf<HTMLInputElement>;
            if (inputs && inputs[index + 1]) {
                // Set timeout to ensure value is updated before focusing next field
                setTimeout(() => {
                    inputs[index + 1].focus();
                }, 10);
            }
        }

        // If backspace is pressed and current field is empty, go back to previous field
        if (value === null && index > 0) {
            const inputs = document.querySelectorAll(
                '.tui-otp-input'
            ) as NodeListOf<HTMLInputElement>;
            if (inputs && inputs[index - 1]) {
                inputs[index - 1].focus();
            }
        }
    }
    // Add this method to your component
    onKeyDown(event: KeyboardEvent, index: number): void {
        // Handle backspace - go to previous input when backspace is pressed in an empty field
        if (event.key === 'Backspace' && this.otpValues[index] === '') {
            if (index > 0) {
                const inputs = document.querySelectorAll(
                    '.tui-otp-input'
                ) as NodeListOf<HTMLInputElement>;
                inputs[index - 1].focus();
                // Optionally clear the previous input
                // this.otpValues[index - 1] = '';
            }
        }

        // Handle arrow left/right for navigation between inputs
        if (event.key === 'ArrowLeft' && index > 0) {
            const inputs = document.querySelectorAll(
                '.tui-otp-input'
            ) as NodeListOf<HTMLInputElement>;
            inputs[index - 1].focus();
        }

        if (event.key === 'ArrowRight' && index < 5) {
            const inputs = document.querySelectorAll(
                '.tui-otp-input'
            ) as NodeListOf<HTMLInputElement>;
            inputs[index + 1].focus();
        }
    }

    // Keep your existing onOtpPaste method but update the selector
    onOtpPaste(event: ClipboardEvent): void {
        event.preventDefault();
        if (!event.clipboardData) return;

        const pastedText = event.clipboardData.getData('text').trim();
        if (!pastedText) return;

        // Changed selector from '.tui-otp-input input' to '.tui-otp-input'
        const otpInputs = document.querySelectorAll(
            '.tui-otp-input'
        ) as NodeListOf<HTMLInputElement>;

        // Fill inputs with pasted characters
        for (
            let i = 0;
            i < Math.min(otpInputs.length, pastedText.length);
            i++
        ) {
            if (/^\d+$/.test(pastedText[i])) {
                this.otpValues[i] = pastedText[i];
                // Update the input value directly
                otpInputs[i].value = pastedText[i];
            }
        }

        // Trigger Angular change detection
        setTimeout(() => {
            // Submit automatically if all 6 digits are filled
            if (this.otpValues.filter((v) => v !== '').length === 6) {
                // Optional: auto-submit when all fields are filled
                // this.verifyOTP();
            } else {
                // Focus on the next empty field
                const nextEmptyIndex = this.otpValues.findIndex(
                    (val) => val === ''
                );
                if (nextEmptyIndex >= 0 && nextEmptyIndex < otpInputs.length) {
                    otpInputs[nextEmptyIndex].focus();
                } else {
                    // Or focus on the last field if all are filled
                    otpInputs[otpInputs.length - 1].focus();
                }
            }
        }, 10);
    }
}
