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
import { RouterModule } from '@angular/router';
import { MessageService } from 'primeng/api';
import { Toast } from 'primeng/toast';
import { ButtonModule } from 'primeng/button';

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
        RouterModule,
        Toast,
        ButtonModule,
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
        'Show Accommodation',
    ];

    protected formCreateManger = new FormGroup(
        {
            account: new FormControl('', Validators.required),
            username: new FormControl('', Validators.required),
            password: new FormControl('', Validators.required),
            confirm: new FormControl('', Validators.required),
        },
        { validators: this.passwordsMatchValidator }
    );

    private readonly dialogs = inject(TuiDialogService);
    protected openDialogCreate(
        content: PolymorpheusContent<TuiDialogContext<string, void>>
    ): void {
        this.formCreateManger.reset();

        this.dialogs
            .open<string>(content, {
                label: 'Tạo manager',
            })
            .subscribe({
                complete: () => {
                    this.formCreateManger.reset();
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
        this.getManagers();
    }

    protected getManagers() {
        this.managerService.getManagers().subscribe({
            next: (value) => {
                this.managers = value.data;
            },
            error: (err) => {
                const message =
                    err.error?.message ||
                    'Không thể tải danh sách quản lý. Vui lòng thử lại sau.';
                this.showToast('error', 'Lỗi tải danh sách quản lý', message);
            },
        });
    }

    protected createManager() {
        this.errorMessage = '';

        const manager: CreateManager = {
            account: this.formCreateManger.get('account')?.value || '',
            password: this.formCreateManger.get('password')?.value || '',
            username: this.formCreateManger.get('username')?.value || '',
        };

        if (this.formCreateManger.invalid) {
            this.formCreateManger.markAllAsTouched();
            return;
        }

        this.managerService.createNewManager(manager).subscribe({
            next: (response) => {
                console.log(response);
                this.formCreateManger.reset();
                this.showToast(
                    'success',
                    'Manager created successfully',
                    response.message
                );
                this.getManagers();
            },
            error: (err) => {
                console.log(err);
                console.log(err.error.error.length);

                for (let index = 0; index < err.error.error.length; index++) {
                    this.formCreateManger
                        .get(err.error.error[index]['field'])
                        ?.setErrors({
                            backend: err.error.error[index]['message'],
                        });
                }
                this.showToast(
                    'error',
                    'Error creating manager',
                    err.error.message
                );
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
