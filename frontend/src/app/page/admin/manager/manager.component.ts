import { Component, inject, OnInit } from '@angular/core';
import {
    AbstractControl,
    FormControl,
    FormGroup,
    FormsModule,
    ReactiveFormsModule,
    ValidationErrors,
    Validators,
} from '@angular/forms';
import { TuiTable } from '@taiga-ui/addon-table';
import {
    TuiButton,
    TuiDialogContext,
    TuiDialogService,
    TuiTextfield,
} from '@taiga-ui/core';
import { TuiInputModule, TuiSelectModule } from '@taiga-ui/legacy';
import { PolymorpheusContent } from '@taiga-ui/polymorpheus';
import { CreateManager, Manager } from '../../../models/admin/manager.model';
import { ManagerService } from '../../../services/admin/manager.service';
import { NavbarComponent } from '../../../components/navbar/navbar.component';
import { MessageService } from 'primeng/api';
import { Toast } from 'primeng/toast';
import { ButtonModule } from 'primeng/button';
import { Ripple } from 'primeng/ripple';

@Component({
    selector: 'app-manager',
    imports: [
        TuiTable,
        TuiButton,
        TuiInputModule,
        TuiSelectModule,
        FormsModule,
        ReactiveFormsModule,
        TuiTextfield,
        NavbarComponent,
        Toast,
        ButtonModule,
        Ripple,
    ],
    templateUrl: './manager.component.html',
    styleUrl: './manager.component.scss',
    providers: [MessageService],
})
export class ManagerComponent implements OnInit {
    protected managers!: Manager[];
    protected errorMessage: string = '';
    protected columns: string[] = [
        'Account',
        'User Name',
        'Is Deleted',
        'Created At',
        'Updated At',
        'Action',
        'Show Accommodation',
    ];

    protected formCreateManger = new FormGroup(
        {
            account: new FormControl('', Validators.required),
            password: new FormControl('', Validators.required),
            confirm: new FormControl('', Validators.required),
        },
        { validators: this.passwordsMatchValidator }
    );

    // protected formManager = new FormGroup({
    //   account: new FormControl('', Validators.required),
    //   user_name: new FormControl('', Validators.required),
    //   login_time: new FormControl('', Validators.required),
    //   logout_time: new FormControl('', Validators.required),
    //   is_deleted: new FormControl('', Validators.required),
    //   created_at: new FormControl('', Validators.required),
    //   updated_at: new FormControl('', Validators.required),
    // });

    private readonly dialogs = inject(TuiDialogService);
    protected openDialogCreate(
        content: PolymorpheusContent<TuiDialogContext<string, void>>
    ): void {
        this.formCreateManger.reset();

        this.dialogs
            .open<string>(content, {
                label: 'Create Manager',
            })
            .subscribe({
                next: (result) => {
                    console.log('Dialog result:', result);
                },
                complete: () => {
                    console.log('Dialog closed');
                },
                error: (err) => {
                    console.error('Dialog error:', err);
                },
            });
    }

    constructor(
        private managerService: ManagerService,
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
        // TODO: get managers by admin
        this.managerService.getManagers().subscribe({
            next: (value) => {
                this.managers = value.data;
                // console.log(this.managers);
            },
            error: (err) => {
                console.error(err);
            },
            complete: () => {
                console.log('Manager fetch complete.');
            },
        });
    }

    protected createManager() {
        this.errorMessage = '';

        const manager: CreateManager = {
            account: this.formCreateManger.get('account')?.value || '',
            password: this.formCreateManger.get('password')?.value || '',
        };

        if (this.formCreateManger.invalid) {
            this.formCreateManger.markAllAsTouched();
            return;
        }

        this.managerService.createNewManager(manager).subscribe({
            next: (response) => {
                // this.managers.push(response.data);

                // console.log("Message: ", response.message);
                // console.log("Status code: ", response.code);

                this.formCreateManger.reset();
                this.showToast(
                    'success',
                    'Tài khoản Quản Lý Đã Được Tạo Thành Công',
                    response.message
                );
            },
            error: (err) => {
                this.showToast(
                    'error',
                    'Lỗi khi tạo tài khoản quản lý',
                    err.error.message
                );
                // console.log('Message:', err.error.message);
                // this.errorMessage = err.error.message;
            },
        });
    }

    protected passwordsMatchValidator(
        group: AbstractControl
    ): ValidationErrors | null {
        const password = group.get('password')?.value;
        const confirm = group.get('confirm')?.value;
        return password === confirm ? null : { passwordMismatch: true };
    }
}
