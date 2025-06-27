import { Component } from '@angular/core';
import {
    FormControl,
    FormGroup,
    ReactiveFormsModule,
    Validators,
} from '@angular/forms';
import { TuiIcon, TuiTextfield } from '@taiga-ui/core';
import { TuiPassword } from '@taiga-ui/kit';
import { Router } from '@angular/router';
import { AdminLoginInput } from '../../../models/admin/admin.model';
import { AuthService } from '../../../services/admin/auth.service';
import { MessageService } from 'primeng/api';
import { Toast } from 'primeng/toast';
import { ButtonModule } from 'primeng/button';
import { Ripple } from 'primeng/ripple';
@Component({
    selector: 'app-login',
    imports: [
        TuiTextfield,
        TuiIcon,
        ReactiveFormsModule,
        TuiPassword,
        Toast,
        ButtonModule,
        Ripple,
    ],
    templateUrl: './login.component.html',
    styleUrl: './login.component.scss',
    providers: [MessageService],
})
export class LoginComponent {
    protected formLogin = new FormGroup({
        account: new FormControl('', Validators.required),
        password: new FormControl('', Validators.required),
    });

    constructor(
        private authSerivce: AuthService,
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
    handleLogin() {
        if (this.formLogin.invalid) {
            this.showToast(
                'error',
                'Đăng nhập thất bại',
                'Vui lòng điền đầy đủ thông tin đăng nhập'
            );
            this.formLogin.markAllAsTouched();
            return;
        }

        let adminLogin: AdminLoginInput = {
            account: this.formLogin.value.account ?? '',
            password: this.formLogin.value.password ?? '',
        };

        this.authSerivce.login(adminLogin).subscribe({
            next: (response) => {
                this.saveTokenToCookie(response.data.token);
                this.router.navigate(['/admin/manager']);
            },
            error: (error) => {
                this.showToast(
                    'error',
                    'Đăng nhập thất bại',
                    error.error.message ||
                        'Vui lòng kiểm tra lại thông tin đăng nhập'
                );
            },
        });
    }

    private saveTokenToCookie(token: string) {
        // Tham số của document.cookie: name=value; expires=date; path=path; domain=domain; secure

        // Thiết lập thời gian hết hạn (1h)
        const expirationDate = new Date();
        expirationDate.setTime(expirationDate.getTime() + 1 * 60 * 60 * 1000);

        // Thiết lập cookie với các tùy chọn bảo mật
        document.cookie = `auth_token=${token}; expires=${expirationDate.toUTCString()}; path=/; SameSite=Strict`;

        // Nếu sử dụng HTTPS, bạn có thể thêm thuộc tính 'secure'
        // document.cookie = `auth_token=${token}; expires=${expirationDate.toUTCString()}; path=/; SameSite=Strict; secure`;
    }
}
