import { DatePipe, NgClass, NgFor, NgIf } from '@angular/common';
import {
    Component,
    ElementRef,
    inject,
    Input,
    OnInit,
    ViewChild,
} from '@angular/core';
import { TuiAlertService, TuiIcon } from '@taiga-ui/core';
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

    private readonly alerts = inject(TuiAlertService);

    protected getAlert(label: string, content: string): void {
        this.alerts
            .open(content, {
                label: label,
                appearance: 'negative',
                autoClose: 5000,
            })
            .subscribe();
    }

    constructor(
        private datePipe: DatePipe,
        private reviewService: ReviewService,
        private messageService: MessageService
    ) {}

    ngOnInit(): void {
        this.totalPages = Math.ceil(this.reviews.length / 10);
    }

    addReview() {
        const token = sessionStorage.getItem('token');

        if (!this.newTitle || this.newTitle.trim() === '') {
            this.showToast(
                'warn',
                'Review Required',
                'Please enter review title.'
            );
            return;
        } else if (!this.newComment || this.newComment.trim() === '') {
            this.showToast(
                'warn',
                'Review Required',
                'Please enter review content.'
            );
            return;
        } else if (this.newRating <= 0 || this.newRating > 5) {
            this.showToast(
                'warn',
                'Rating Required',
                'Please select a rating between 1 and 5.'
            );
            return;
        } else if (!token) {
            this.showToast(
                'error',
                'Authentication Required',
                'Please log in first'
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
                this.showToast(
                    'success',
                    'Review Submitted',
                    'Your review has been added successfully!'
                );
                // Add the new review to the top of the list
                setTimeout(() => {
                    this.reviews.unshift(response);
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
                    'Review Failed',
                    'Have error when adding review. Please try again later.'
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
        const token = sessionStorage.getItem('token');

        if (!token) {
            this.showToast(
                'error',
                'Authentication Required',
                'Please log in first'
            );
            return;
        } else if (orderId == '' || orderId.length < 8) {
            this.showToast(
                'warn',
                'Input Required',
                'Please input a valid order ID'
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
