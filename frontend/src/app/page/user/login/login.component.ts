import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Router, RouterLink } from '@angular/router';
import { AuthService } from '../../../services/user/auth.service';
import { ReactiveFormsModule } from '@angular/forms';
import { LoginModel } from '../../../models/user/auth.model';
import { SaveTokenToCookie } from '../../../shared/token/token';
import { MessageService } from 'primeng/api';
import { Toast } from 'primeng/toast';
import { ButtonModule } from 'primeng/button';
import { Ripple } from 'primeng/ripple';
import { provideAnimations } from '@angular/platform-browser/animations';

@Component({
    selector: 'app-login',
    imports: [ReactiveFormsModule, RouterLink, Toast, ButtonModule, Ripple],
    templateUrl: './login.component.html',
    styleUrl: './login.component.scss',
    providers: [MessageService, provideAnimations()],
})
export class LoginComponent implements OnInit {
    loginForm: FormGroup;
    isLoading: boolean = false;
    errorMessage: string = '';
    showPassword: boolean = false;

    constructor(
        private fb: FormBuilder,
        private router: Router,
        private authService: AuthService,
        private messageService: MessageService
    ) {
        this.loginForm = this.fb.group({
            email: ['', [Validators.required, Validators.email]],
            password: ['', Validators.required],
            rememberMe: [false],
        });
    }
    showToast(
        severity: 'success' | 'info' | 'warn' | 'error',
        summary: string,
        detail: string
    ): void {
        this.messageService.add({ severity, summary, detail });
    }

    ngOnInit(): void {
        // Check if user is already logged in
        // if (this.userService.isLoggedIn()) {
        //   this.router.navigate(['/']);
        // }
    }

    get email() {
        return this.loginForm.get('email')!;
    }
    get password() {
        return this.loginForm.get('password')!;
    }

    togglePasswordVisibility(): void {
        this.showPassword = !this.showPassword;
    }

    onSubmit(): void {
        if (this.loginForm.invalid) {
            this.showToast(
                'error',
                'Form không hợp lệ',
                'Vui lòng kiểm tra lại thông tin.'
            );
            return;
        }

        this.isLoading = true;
        this.errorMessage = '';

        const loginData: LoginModel = {
            account: this.email.value,
            password: this.password.value,
        };

        this.authService.loginUser(loginData).subscribe({
            next: (response) => {
                // console.log('Login successful:', response);
                SaveTokenToCookie(response.data.token);
                // Navigate to home or dashboard
                this.router.navigate(['/']);
            },
            error: (error) => {
                this.showToast(
                    'error',
                    'Đăng nhập thất bại',
                    error.error?.message ||
                        'Email hoặc mật khẩu không đúng. Vui lòng thử lại.'
                );
                // console.error('Login failed:', error);
                // this.errorMessage =
                //     error.error?.message ||
                //     'Invalid email or password. Please try again.';
                this.isLoading = false;
            },
            complete: () => {
                this.isLoading = false;
            },
        });
    }
    // loginWithGoogle(): void {
    //     // Implement Google login
    //     console.log('Google login clicked');
    // }

    // loginWithFacebook(): void {
    //     // Implement Facebook login
    //     console.log('Facebook login clicked');
    // }

    // loginWithApple(): void {
    //     // Implement Apple login
    //     console.log('Apple login clicked');
    // }
}
