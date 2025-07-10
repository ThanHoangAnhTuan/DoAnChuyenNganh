import { Component, OnInit } from '@angular/core';
import { Toast } from 'primeng/toast';
import { NavbarComponent } from '../../../components/navbar/navbar.component';
import { MessageService } from 'primeng/api';
import { OrderService } from '../../../services/manager/order.service';
import { Order } from '../../../models/manager/order.model';
import { TuiDataListWrapper, TuiDataListWrapperComponent } from '@taiga-ui/kit';
import { TuiTable } from '@taiga-ui/addon-table';

@Component({
    selector: 'app-order',
    imports: [
        NavbarComponent,
        Toast,
        TuiDataListWrapperComponent,
        TuiDataListWrapper,
        TuiTable
    ],
    templateUrl: './order.component.html',
    styleUrl: './order.component.scss',
    providers: [MessageService],
})
export class OrderComponent implements OnInit {
    orders: Order[] = [];

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

    constructor(
        private messageService: MessageService,
        private orderService: OrderService
    ) {}

    ngOnInit() {
        this.orderService.getOrdersByManager().subscribe({
            next: (response) => {
                this.orders = response.data || [];
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
}
