import { Component, inject, OnInit } from '@angular/core';
import { AbstractControl, FormControl, FormGroup, FormsModule, ReactiveFormsModule, ValidationErrors, Validators } from '@angular/forms';
import { TuiTable } from '@taiga-ui/addon-table';
import { TuiButton, TuiDialogContext, TuiDialogService, TuiTextfield } from '@taiga-ui/core';
import { TuiInputModule, TuiSelectModule } from '@taiga-ui/legacy';
import { PolymorpheusContent } from '@taiga-ui/polymorpheus';
import { CreateManager, Manager } from '../../../models/admin/manager.model';
import { ManagerService } from '../../../services/admin/manager.service';
import { NavbarComponent } from "../../../components/navbar/navbar.component";

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
    NavbarComponent
],
  templateUrl: './manager.component.html',
  styleUrl: './manager.component.scss'
})
export class ManagerComponent implements OnInit {
  protected managers!: Manager[];
  protected errorMessage: string = '';
  protected columns: string[] = [
    'Account',
    'User Name',
    'Login Time',
    'Logout Time',
    'Is Deleted',
    'Created At',
    'Updated At',
    'Action',
    'Show Accommodation',
  ];

  protected formCreateManger = new FormGroup({
    account: new FormControl('', Validators.required),
    password: new FormControl('', Validators.required),
    confirm: new FormControl('', Validators.required)
  }, { validators: this.passwordsMatchValidator });

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
  protected openDialogCreate(content: PolymorpheusContent<TuiDialogContext<string, void>>): void {
    this.formCreateManger.reset();

    this.dialogs.open<string>(content, {
      label: 'Create Manager',
    }).subscribe({
      next: result => {
        console.log("Dialog result:", result);
      },
      complete: () => {
        console.log('Dialog closed');
      },
      error: err => {
        console.error('Dialog error:', err);
      }
    });
  }

  constructor(
    private managerService: ManagerService,
  ) { }

  ngOnInit(): void {

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

    this.managerService
      .createNewManager(manager).subscribe({
        next: (response) => {
          // this.managers.push(response.data);

          // console.log("Message: ", response.message);
          // console.log("Status code: ", response.code);

          this.formCreateManger.reset();
        },
        error: (err) => {
          console.log("Message:", err.error.message);
          this.errorMessage = err.error.message;
        }
      });
  }

  protected passwordsMatchValidator(group: AbstractControl): ValidationErrors | null {
    const password = group.get('password')?.value;
    const confirm = group.get('confirm')?.value;
    return password === confirm ? null : { passwordMismatch: true };
  }
}
