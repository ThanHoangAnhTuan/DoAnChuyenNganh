import { Component, OnInit } from '@angular/core';
import { Toast } from 'primeng/toast';
import { NavbarComponent } from '../../../components/navbar/navbar.component';
import { MessageService } from 'primeng/api';
import { OrderService } from '../../../services/manager/order.service';
import { Order, OrderDetail } from '../../../models/manager/order.model';
import { TuiDataListWrapper, TuiDataListWrapperComponent } from '@taiga-ui/kit';
import { TuiTable } from '@taiga-ui/addon-table';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { TuiButton, TuiIcon } from '@taiga-ui/core';

@Component({
    selector: 'app-order',
    imports: [
        NavbarComponent,
        Toast,
        TuiDataListWrapperComponent,
        TuiDataListWrapper,
        TuiTable,
        ReactiveFormsModule,
        FormsModule,
        TuiButton,
        TuiIcon,
    ],
    templateUrl: './order.component.html',
    styleUrl: './order.component.scss',
    providers: [MessageService],
})
export class OrderComponent implements OnInit {
    protected orders: Order[] = [];
    protected orderDetails: OrderDetail[] = [];
    protected isChangeStatus: boolean = false;
    protected isChangeStatusConfirmOpen: boolean = false;
    protected isOrderDetailOpen: boolean = false;
    protected updateId: string = '';
    protected statusOptions = '';

    protected columns: string[] = [
        'ID',
        'Username',
        'Phone',
        'Email',
        'Accommodation Name',
        'Check In',
        'Check Out',
        'Final Total',
        'Order Status',
        'Order Detail',
        // 'Actions',
    ];

    protected detailColumns: string[] = [
        'Accommodation Name',
        'Total Price',
        'Guests',
    ];

    constructor(
        private messageService: MessageService,
        private orderService: OrderService
    ) { }

    ngOnInit() {
        this.orderService.getOrdersByManager().subscribe({
            next: (response) => {
                this.orders = response.data || [];

                console.log('Orders:', this.orders);
            },
            error: (error) => {
                const message =
                    error.error?.message ||
                    'Đã xảy ra lỗi. Vui lòng thử lại sau.';
                this.messageService.add({
                    severity: 'error',
                    summary: 'Error',
                    detail: message,
                });
            },
        });
    }

    protected changeOrderStatus() {
        this.orderService.changeOrderStatus(this.updateId, this.statusOptions).subscribe((response) => {
            this.messageService.add({
                severity: 'success',
                summary: 'Success',
                detail: 'Order status updated successfully.',
            });
            this.closeConfirmModal();
        })
    }

    protected changeStatusApply(id: string) {
        this.updateId = id;
        console.log('Update ID:', this.updateId);
        this.isChangeStatus = true;
    }

    protected onChangeStatus(event: Event) {
        const target = event.target as HTMLSelectElement;
        const value = target.value;
        this.changeStatusOption(value);
    }

    protected changeStatusOption(option: string) {
        this.statusOptions = option;
        console.log('Selected status option:', this.statusOptions);
        this.isChangeStatusConfirmOpen = true;
    }

    protected closeConfirmModal() {
        this.isChangeStatus = false;
        this.updateId = '';
        this.ngOnInit(); // Refresh the order list
    }

    protected openOrderDetail(orderDetails: OrderDetail[]) {
        this.orderDetails = orderDetails;
        this.isOrderDetailOpen = true;
        // console.log('Order Details:', this.orderDetails);
    }

    protected closeOrderDetail() {
        this.isOrderDetailOpen = false;
        this.orderDetails = [];
    }
}
