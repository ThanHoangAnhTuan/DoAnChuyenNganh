import { Component, OnInit } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { ActivatedRoute } from '@angular/router';
import { UserService } from '../../../services/user/user.service';
import { OTP, UpdatePassword } from '../../../models/user/register.model';


@Component({
    selector: 'app-verify-otp',
    imports: [FormsModule],
    templateUrl: './verify-otp.component.html',
    styleUrl: './verify-otp.component.scss',
})
export class VerifyOtpComponent implements OnInit {
    otp: string = '';
    password: string = '';
    token: string = '';
    email: string = '';

    constructor(private route: ActivatedRoute, private userService: UserService) {}

    ngOnInit(): void {
        // Lấy email từ query params
        this.route.queryParams.subscribe((params) => {
            this.email = params['email'] || '';
            console.log('Email from query param:', this.email);
        });
    }

    verifyOTP() {
        if (!this.otp || this.otp.trim() === '') {
            console.warn('Please enter otp.');
            return;
        }

        const otpData: OTP = {
            verify_key: this.email,
            verify_code: this.otp,
        };

        if (this.otp && this.email) {
            this.userService.verifyOTP(otpData).subscribe({
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

        this.userService.updatePassword(passwordData).subscribe({
            next: (response) => {
                console.log('Password updated successfully:', response);
                // Navigate to success page or show success message
            },
            error: (error) => {
                console.error('Error updating password:', error);
                // Handle error (show error message to user)
            },
        });
    }
}
