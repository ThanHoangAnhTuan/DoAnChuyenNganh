import { Component } from '@angular/core';
import {
    FormControl,
    FormGroup,
    ReactiveFormsModule,
    Validators,
} from '@angular/forms';
import { TuiIcon, TuiTextfield } from '@taiga-ui/core';
import { TuiPassword } from '@taiga-ui/kit';
import { ManagerService } from '../../services/manager/manager.service';
import { ManagerLoginInput } from '../../models/accommodation.model';
import { Router } from '@angular/router';

@Component({
    selector: 'app-login',
    imports: [TuiTextfield, TuiIcon, ReactiveFormsModule, TuiPassword],
    templateUrl: './login.component.html',
    styleUrl: './login.component.scss',
})
export class LoginComponent {
    protected formLogin = new FormGroup({
        account: new FormControl('', Validators.required),
        password: new FormControl('', Validators.required),
    });

    constructor(private managerSerivce: ManagerService, private router: Router) {}

    handleLogin() {
        if (this.formLogin.invalid) {
            this.formLogin.markAllAsTouched();
            return;
        }

        let managerLogin: ManagerLoginInput = {
            account: this.formLogin.value.account ?? '',
            password: this.formLogin.value.password ?? '',
        };

        this.managerSerivce.login(managerLogin).subscribe((response) => {
            this.saveTokenToCookie(response.data.token);
            this.router.navigate(['/manager/accommodation']);
        });
    }

    private saveTokenToCookie(token: string) {
        // Tham số của document.cookie: name=value; expires=date; path=path; domain=domain; secure

        // Thiết lập thời gian hết hạn (1h)
        const expirationDate = new Date();
        expirationDate.setTime(expirationDate.getTime() + (1 * 60 * 60 * 1000));

        // Thiết lập cookie với các tùy chọn bảo mật
        document.cookie = `auth_token=${token}; expires=${expirationDate.toUTCString()}; path=/; SameSite=Strict`;

        // Nếu sử dụng HTTPS, bạn có thể thêm thuộc tính 'secure'
        // document.cookie = `auth_token=${token}; expires=${expirationDate.toUTCString()}; path=/; SameSite=Strict; secure`;
    }
}
