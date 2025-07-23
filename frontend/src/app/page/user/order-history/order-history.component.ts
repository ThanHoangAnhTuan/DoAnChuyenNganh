import { Component } from '@angular/core';
import { Order } from '../../../models/manager/order.model';
import { OrderHistoryService } from '../../../services/user/order-history.service';
import { AccommodationService } from '../../../services/user/accommodation.service';
import { NavbarComponent } from "../../../components/navbar/navbar.component";

@Component({
  selector: 'app-order-history',
  imports: [NavbarComponent],
  templateUrl: './order-history.component.html',
  styleUrl: './order-history.component.scss'
})
export class OrderHistoryComponent {
  public orders: Order[] = [];
  imagesMap: { [key: string]: string[] } = {};

  constructor(
    private orderHistoryService: OrderHistoryService,
    private accommodationService: AccommodationService,
  ) { }

  public ngOnInit(): void {
    this.getOrdersByUser();
  }

  public getOrdersByUser(): void {
    this.orderHistoryService.getOrdersByUser().subscribe({
      next: (response) => {
        this.orders = response.data;
        this.orders.forEach(order => {
          this.loadImages(order.accommodation_id);
        });

        console.log('Order history fetched successfully:', this.orders);
      },
      error: (error) => {
        console.error('Error fetching order history:', error);
      }
    });
  }

  public loadImages(id: string): void {
    this.accommodationService.getAccommodationDetailById(id).subscribe({
      next: (response) => {
        this.imagesMap[id] = response.data.images || [];
      },
      error: (error) => {
        console.error('Lỗi khi tải ảnh:', error);
        this.imagesMap[id] = [];
      }
    });
  }
}
