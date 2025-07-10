import { Component, inject, Injector, OnInit } from '@angular/core';
import {
    FormControl,
    FormGroup,
    FormsModule,
    ReactiveFormsModule,
    Validators,
} from '@angular/forms';
import { TuiTable } from '@taiga-ui/addon-table';
import {
    TuiButton,
    TuiDialogService,
    TuiTextfield,
    TuiAppearance,
} from '@taiga-ui/core';
import type { PolymorpheusContent } from '@taiga-ui/polymorpheus';
import { TuiInputModule, TuiSelectModule } from '@taiga-ui/legacy';
import {
    TuiConfirmService,
    TuiFiles,
    TuiSelect,
    TuiDataListWrapperComponent,
    TuiDataListWrapper,
} from '@taiga-ui/kit';
import { TuiResponsiveDialogService } from '@taiga-ui/addon-mobile';
import { TuiCardLarge } from '@taiga-ui/layout';
import { ActivatedRoute } from '@angular/router';
import { TuiInputTimeModule } from '@taiga-ui/legacy';
import { NavbarComponent } from '../../../components/navbar/navbar.component';
import { MessageService } from 'primeng/api';
import { Toast } from 'primeng/toast';
import { ButtonModule } from 'primeng/button';
import { RoomService } from '../../../services/manager/room.service';
import { CreateRoom, Room } from '../../../models/manager/room.model';
import { LoaderComponent } from '../../../components/loader/loader.component';
import { finalize } from 'rxjs';

@Component({
    selector: 'app-room',
    imports: [
        TuiTable,
        TuiButton,
        TuiInputModule,
        FormsModule,
        ReactiveFormsModule,
        TuiTextfield,
        TuiAppearance,
        TuiCardLarge,
        TuiFiles,
        TuiInputTimeModule,
        TuiSelect,
        TuiSelectModule,
        TuiDataListWrapperComponent,
        TuiDataListWrapper,
        NavbarComponent,
        Toast,
        ButtonModule,
        LoaderComponent,
    ],
    templateUrl: './room.component.html',
    styleUrl: './room.component.scss',
    providers: [
        MessageService,
        TuiConfirmService,
        {
            provide: TuiDialogService,
            useExisting: TuiResponsiveDialogService,
        },
    ],
})
export class RoomComponent implements OnInit {
    protected columns: string[] = ['ID', 'Name', 'Status', 'Action'];

    protected status: string[] = ['available', 'unavailable', 'occupied'];
    isLoading: boolean = false;
    private readonly dialogs = inject(TuiDialogService);

    protected formCreateRoom = new FormGroup({
        prefix: new FormControl('', Validators.required),
        quantity: new FormControl<number | null>(null, Validators.required),
    });

    protected formUpdateRoom = new FormGroup({
        name: new FormControl('', Validators.required),
        status: new FormControl('', Validators.required),
    });

    protected accommodationDetailId: string = '';
    protected accommodationRoomId: string = '';
    protected rooms: Room[] = [];

    constructor(
        private roomService: RoomService,
        private messageService: MessageService,
        private route: ActivatedRoute
    ) {}

    showToast(
        severity: 'success' | 'info' | 'warn' | 'error',
        summary: string,
        detail: string
    ): void {
        this.messageService.add({ severity, summary, detail });
    }

    ngOnInit() {
        this.isLoading = true;
        this.route.params.subscribe((params) => {
            this.accommodationDetailId = params['id'];
            this.roomService
                .getAccommodationRooms(params['id'])
                .pipe(finalize(() => (this.isLoading = false)))
                .subscribe((response) => {
                    this.rooms = response.data;
                });
        });
    }

    protected openDialogCreate(content: PolymorpheusContent): void {
        this.formCreateRoom.reset();
        this.dialogs
            .open(content, {
                label: 'Create Room',
            })
            .subscribe({
                complete: () => {
                    this.formCreateRoom.reset();
                },
            });
    }

    protected openDialogUpdate(content: PolymorpheusContent, room: Room) {
        this.formUpdateRoom.reset();
        this.accommodationRoomId = room.id;
        this.formUpdateRoom.patchValue({
            name: room.name,
            status: room.status,
        });
        this.dialogs
            .open(content, {
                label: 'Update Room',
            })
            .subscribe({
                complete: () => {
                    this.formUpdateRoom.reset();
                },
            });
    }

    protected createRoom() {
        if (this.formCreateRoom.invalid) {
            this.formCreateRoom.markAllAsTouched();
            return;
        }
        const room: CreateRoom = {
            prefix: this.formCreateRoom.get('prefix')?.value || '',
            quantity: Number(this.formCreateRoom.get('quantity')?.value) || 0,
            accommodation_type_id: this.accommodationDetailId,
        };
        this.isLoading = true;
        this.roomService
            .createAccommodationRoom(room)
            .pipe(finalize(() => (this.isLoading = false)))
            .subscribe({
                next: (response) => {
                    this.rooms.push(...response.data);
                    this.formCreateRoom.reset();
                    this.showToast(
                        'success',
                        'Phòng đã được tạo thành công',
                        'Bạn có thể xem chi tiết phòng trong danh sách'
                    );
                },
                error: (error) => {
                    console.error('Error creating room:', error);
                    this.showToast(
                        'error',
                        'Tạo phòng thất bại',
                        'Vui lòng thử lại sau'
                    );
                },
            });
    }

    protected updateRoom() {
        if (this.formUpdateRoom.invalid) {
            this.formUpdateRoom.markAllAsTouched();
            return;
        }
        const room: Room = {
            id: this.accommodationRoomId,
            name: this.formUpdateRoom.get('name')?.value || '',
            status: this.formUpdateRoom.get('status')?.value || '',
        };
        this.isLoading = true;
        this.roomService
            .updateAccommodationRoom(room)
            .pipe(finalize(() => (this.isLoading = false)))
            .subscribe({
                next: (response) => {
                    this.rooms = this.rooms.map((room) => {
                        if (room.id === response.data.id) {
                            return response.data;
                        } else {
                            return room;
                        }
                    });
                    this.showToast(
                        'success',
                        'Cập nhật phòng thành công',
                        'Bạn có thể xem chi tiết phòng trong danh sách'
                    );
                },
                error: (error) => {
                    console.error('Lỗi khi thêm đánh giá:', error);
                    this.showToast(
                        'error',
                        'Cập nhật phòng thất bại',
                        'Cập nhật phòng thất bại, vui lòng thử lại sau'
                    );
                },
            });
    }

    protected deleteRoom(id: string) {
        this.isLoading = true;
        this.roomService
            .deleteAccommodationRoom(id)
            .pipe(finalize(() => (this.isLoading = false)))
            .subscribe({
                next: (value) => {
                    this.rooms = this.rooms.filter((room) => room.id !== id);
                    this.showToast(
                        'success',
                        'Xoá phòng thành công',
                        'Phòng đã được xoá khỏi danh sách'
                    );
                },
                error: (err) => {
                    this.showToast(
                        'error',
                        'Xoá phòng thất bại',
                        err.error.message || 'Vui lòng thử lại sau'
                    );
                },
                complete() {},
            });
    }
}
