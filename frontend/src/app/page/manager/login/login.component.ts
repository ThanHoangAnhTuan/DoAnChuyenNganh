import { Component } from '@angular/core';
import {
    FormControl,
    FormGroup,
    ReactiveFormsModule,
    Validators,
} from '@angular/forms';
import { TuiIcon, TuiTextfield } from '@taiga-ui/core';
import { TuiPassword } from '@taiga-ui/kit';
import { AuthService } from '../../../services/manager/auth.service';
import { ManagerLoginInput } from '../../../models/manager/accommodation.model';
import { Router } from '@angular/router';
import { SaveTokenToCookie } from '../../../shared/token/token';
import { MessageService } from 'primeng/api';
import { Toast } from 'primeng/toast';
import { ButtonModule } from 'primeng/button';
import { LoaderComponent } from '../../../components/loader/loader.component';

@Component({
    selector: 'app-login',
    imports: [
        TuiTextfield,
        TuiIcon,
        ReactiveFormsModule,
        TuiPassword,
        Toast,
        ButtonModule,
        LoaderComponent,
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
    isLoading: boolean = false;

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
            this.formLogin.markAllAsTouched();
            return;
        }
        this.isLoading = true;

        let managerLogin: ManagerLoginInput = {
            account: this.formLogin.value.account ?? '',
            password: this.formLogin.value.password ?? '',
        };

        this.authSerivce.login(managerLogin).subscribe({
            next: (response) => {
                SaveTokenToCookie(response.data.token);
                this.router.navigate(['/manager/accommodation']);
            },
            error: (error) => {
                console.error('Login error:', error);
                this.showToast(
                    'error',
                    'Đăng nhập thất bại',
                    error.error.message ||
                        'Vui lòng kiểm tra lại thông tin đăng nhập.'
                );
                this.isLoading = false;
            },
            complete: () => {
                this.isLoading = false;
            },
        });
    }
}
