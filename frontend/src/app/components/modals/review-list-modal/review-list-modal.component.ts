import { DatePipe, NgClass, NgFor, NgIf } from '@angular/common';
import {
    Component,
    ElementRef,
    inject,
    Input,
    OnInit,
    ViewChild,
} from '@angular/core';
import { TuiIcon } from '@taiga-ui/core';
import { CreateNewReview } from '../../../models/user/review.model';
import { ReviewService } from '../../../services/user/review.service';
import {
    FormControl,
    FormGroup,
    FormsModule,
    ReactiveFormsModule,
    Validators,
} from '@angular/forms';
import { TuiRating } from '@taiga-ui/kit';
import { TuiInputModule } from '@taiga-ui/legacy';
import { MessageService } from 'primeng/api';
import { Toast } from 'primeng/toast';
import { ButtonModule } from 'primeng/button';
import { GetToken } from '../../../shared/token/token';
@Component({
    selector: 'app-review-list-modal',
    imports: [
        NgIf,
        NgFor,
        NgClass,
        TuiIcon,
        FormsModule,
        TuiRating,
        TuiInputModule,
        ReactiveFormsModule,
        Toast,
        ButtonModule,
    ],
    providers: [DatePipe, MessageService],
    templateUrl: './review-list-modal.component.html',
    styleUrl: './review-list-modal.component.scss',
})
export class ReviewListModalComponent implements OnInit {
    @Input() reviews: any[] = [];
    @Input() avarageRating: number = 0;
    @Input() accommodationId: string = '';
    @ViewChild('listTop') listTop!: ElementRef;
    currentPage: number = 1;
    isInputOrderIdModalOpen: boolean = false;
    isAddRivewModalOpen: boolean = false;
    newTitle: string = '';
    newComment: string = '';
    newRating: number = 0;
    totalPages: number = 0;
    inputForm = new FormGroup({
        orderId: new FormControl('', Validators.minLength(8)),
    });
    orderIdValue: string = '';

    constructor(
        private datePipe: DatePipe,
        private reviewService: ReviewService,
        private messageService: MessageService
    ) {}

    ngOnInit(): void {
        this.totalPages = Math.ceil(this.reviews.length / 10);
    }

    addReview() {
        const token = GetToken();
        console.log(token);


        if (!this.newTitle || this.newTitle.trim() === '') {
            this.showToast(
                'warn',
                'Đánh giá bắt buộc',
                'Vui lòng nhập tiêu đề đánh giá.'
            );
            return;
        } else if (!this.newComment || this.newComment.trim() === '') {
            this.showToast(
                'warn',
                'Nội dung đánh giá bắt buộc',
                'Vui lòng nhập nội dung đánh giá.'
            );
            return;
        } else if (this.newRating <= 0 || this.newRating > 5) {
            this.showToast(
                'warn',
                'Đánh giá không hợp lệ',
                'Vui lòng chọn đánh giá từ 1 đến 5.'
            );
            return;
        } else if (!token) {
            console.log(token);
            this.showToast(
                'error',
                'Yêu cầu xác thực',
                'Vui lòng đăng nhập trước'
            );
            return;
        }

        const newReview: CreateNewReview = {
            accommodation_id: this.accommodationId,
            title: this.newTitle,
            comment: this.newComment,
            rating: this.newRating,
            order_id: this.orderIdValue,
        };

        this.reviewService.addReview(newReview).subscribe({
            next: (response) => {
                this.showToast('success', 'Đánh giá của bạn đã được thêm', '');
                // console.log('Review has been added successfull:', response);
                console.log('Đánh giá đã được thêm thành công:', response.data);
                
                // Add the new review to the top of the list
                setTimeout(() => {
                    this.reviews.unshift(response.data);
                }, 1000);
                // Update total pages
                this.totalPages = Math.ceil(this.reviews.length / 10);
                // Reset form fields
                this.newTitle = '';
                this.newComment = '';
                this.newRating = 0;
                this.isAddRivewModalOpen = false;
            },
            error: (error) => {
                console.error('Lỗi khi thêm đánh giá:', error);
                this.showToast(
                    'error',
                    'Lỗi khi thêm đánh giá',
                    error.error?.message || 'Đã xảy ra lỗi khi thêm đánh giá.'
                );
                // Reset form fields
                this.newTitle = '';
                this.newComment = '';
                this.newRating = 0;
                this.isAddRivewModalOpen = false;
            },
        });
    }
    showToast(
        severity: 'success' | 'info' | 'warn' | 'error',
        summary: string,
        detail: string
    ): void {
        this.messageService.add({ severity, summary, detail });
    }

    formatDate(dateStr: string): string {
        const date = new Date(dateStr);
        const hoursMinutes = this.datePipe.transform(date, 'HH:mm');
        const day = date.getUTCDate();
        const month = date.getUTCMonth() + 1;
        const year = date.getUTCFullYear();
        return `${hoursMinutes} ngày ${day} tháng ${month} năm ${year}`;
    }

    onChangePage(page: number) {
        this.currentPage = page;
        // Cuộn đến phần tử
        this.listTop.nativeElement.scrollIntoView({ behavior: 'smooth' });
    }

    onPreviousPage(): void {
        if (this.currentPage > 1) {
            this.currentPage--;
            this.listTop.nativeElement.scrollIntoView({ behavior: 'smooth' });
        } else {
            document.getElementById('previous-page')?.blur();
        }
    }

    onNextPage(): void {
        if (this.currentPage < this.totalPages) {
            this.currentPage++;
            this.listTop.nativeElement.scrollIntoView({ behavior: 'smooth' });
        } else {
            document.getElementById('next-page')?.blur();
        }
    }

    onOpenInputOrderIdModal(): void {
        this.isInputOrderIdModalOpen = true;
    }

    onCloseInputOrderIdModal(): void {
        this.isInputOrderIdModalOpen = false;
    }

    submitOrderId(): void {
        const orderId = this.inputForm.value.orderId ?? '';
        const token = GetToken();

        if (!token) {
            this.showToast(
                'error',
                'Yêu cầu xác thực',
                'Vui lòng đăng nhập trước'
            );
            return;
        } else if (orderId == '' || orderId.length < 8) {
            this.showToast(
                'warn',
                'Thông tin bắt buộc',
                'Vui lòng nhập một ID đơn hàng hợp lệ'
            );
            return;
        } else {
            this.onCloseInputOrderIdModal();
            this.onOpenAddReviewModal();
            this.orderIdValue = orderId;
            return;
        }
    }

    onOpenAddReviewModal(): void {
        this.isAddRivewModalOpen = true;
    }

    onCloseAddReviewModal(): void {
        this.isAddRivewModalOpen = false;
    }
}
