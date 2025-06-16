import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Router, RouterLink } from '@angular/router';
import { AuthService } from '../../../services/user/auth.service';
import { ReactiveFormsModule } from '@angular/forms';
import { LoginModel } from '../../../models/user/auth.model';
import { SaveTokenToCookie } from '../../../shared/token/token';

@Component({
    selector: 'app-login',
    imports: [ReactiveFormsModule, RouterLink],
    templateUrl: './login.component.html',
    styleUrl: './login.component.scss',
})
export class LoginComponent implements OnInit {
    loginForm: FormGroup;
    isLoading: boolean = false;
    errorMessage: string = '';
    showPassword: boolean = false;

    constructor(
        private fb: FormBuilder,
        private router: Router,
        private authService: AuthService
    ) {
        this.loginForm = this.fb.group({
            email: ['', [Validators.required, Validators.email]],
            password: ['', Validators.required],
            rememberMe: [false],
        });
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
                console.log('Login successful:', response);

                SaveTokenToCookie(response.data.token);

                // Navigate to home or dashboard
                this.router.navigate(['/']);
            },
            error: (error) => {
                console.error('Login failed:', error);
                this.errorMessage =
                    error.error?.message ||
                    'Invalid email or password. Please try again.';
                this.isLoading = false;
            },
            complete: () => {
                this.isLoading = false;
            },
        });
    }
    loginWithGoogle(): void {
        // Implement Google login
        console.log('Google login clicked');
    }

    loginWithFacebook(): void {
        // Implement Facebook login
        console.log('Facebook login clicked');
    }

    loginWithApple(): void {
        // Implement Apple login
        console.log('Apple login clicked');
    }
}
