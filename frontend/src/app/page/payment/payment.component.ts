import { CommonModule, NgIf } from '@angular/common';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, RouterModule } from '@angular/router';
import { format } from 'date-fns';
import { AccommodationDetailService } from '../../services/user/accommodation-detail.service';

@Component({
  selector: 'app-payment',
  imports: [
    NgIf,
    CommonModule,
    RouterModule,
  ],
  templateUrl: './payment.component.html',
  styleUrl: './payment.component.scss'
})
export class PaymentComponent implements OnInit {
  amount: number = 0;
  bank: string = '';
  order_id: string = '';
  pay_date: string = '';
  response_code: string = '';
  transaction_no: string = '';

  constructor(
    private route: ActivatedRoute,
    // private accommodationDetailService = AccommodationDetailService,
  ) { }

  ngOnInit(): void {
    this.route.queryParams.subscribe(params => {
      this.amount = params['amount'];
      this.bank = params['bank_code'];
      this.order_id = params['order_id'];
      // this.pay_date = params['pay_date'];
      const rawDate = params['pay_date'];
      this.response_code = params['response_code'];
      this.transaction_no = params['transaction_no'];

      const formattedDate = new Date(
        Number(rawDate.slice(0, 4)),        // year
        Number(rawDate.slice(4, 6)) - 1,    // month (0-based)
        Number(rawDate.slice(6, 8)),        // day
        Number(rawDate.slice(8, 10)),       // hour
        Number(rawDate.slice(10, 12)),      // minute
        Number(rawDate.slice(12, 14))       // second
      );
      this.pay_date = format(formattedDate, 'dd/MM/yyyy HH:mm:ss');

      // console.log("amount: ", this.amount);
      // console.log("bank: ", this.bank);
      // console.log("order id: ", this.order_id);
      // console.log("pay date: ", this.pay_date);
      // console.log("respone code: ", this.response_code);
      // console.log("transaction no: ", this.transaction_no);
    });
  }


}
