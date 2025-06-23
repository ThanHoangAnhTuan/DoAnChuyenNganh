import { Component, ElementRef, ViewChild } from '@angular/core';
import {
    AbstractControl,
    FormControl,
    FormGroup,
    ValidatorFn,
    ReactiveFormsModule,
} from '@angular/forms';

import { TuiValidationError } from '@taiga-ui/cdk';
import { TuiFiles } from '@taiga-ui/kit';
import { TuiCardLarge } from '@taiga-ui/layout';
import { TuiButton, TuiIcon } from '@taiga-ui/core';
import { ActivatedRoute, Router } from '@angular/router';
import { ImageService } from '../../../services/manager/image.service';
import { NavbarComponent } from '../../../components/navbar/navbar.component';

@Component({
    selector: 'app-media-library',
    imports: [ReactiveFormsModule, TuiCardLarge, TuiFiles, TuiButton, NavbarComponent],
    templateUrl: './media-library.component.html',
    styleUrl: './media-library.component.scss',
})
export class MediaLibraryComponent {
    protected imagesPreview: string[] = [];
    protected oldImages: string[] | null = null;

    protected readonly apiURl = 'http://localhost:8080/uploads/';

    @ViewChild('fileInput')
    fileInput!: ElementRef<HTMLInputElement>;
    protected id = '';

    protected isDetailMode = false;

    constructor(
        private activatedRoute: ActivatedRoute,
        private router: Router,
        private imageService: ImageService
    ) {}

    ngOnInit() {
        this.activatedRoute.params.subscribe((params) => {
            this.isDetailMode = this.router.url.includes('/detail/');
            this.id = params['id'];
            this.imageService
                .getImages(this.id, this.isDetailMode)
                .subscribe((response) => {
                    this.oldImages = response.data;
                });
        });
    }

    protected formImages = new FormGroup({
        images: new FormControl<File[] | []>([], [this.maxFilesLength(20)]),
    });

    protected onSelected(event: Event): void {
        const input = event.target as HTMLInputElement;
        console.log('upload');

        if (input.files && input.files.length > 0) {
            const filesArray = Array.from(input.files); // convert FileList -> File[]

            filesArray.forEach((file) => {
                const reader = new FileReader();
                reader.onload = () => {
                    this.imagesPreview.push(reader.result as string);
                };
                reader.readAsDataURL(file);
            });
        }
    }

    protected onRemoveOldImage(imageName: string): void {
        if (this.oldImages) {
            this.oldImages = this.oldImages.filter(
                (item) => item !== imageName
            );
        }
    }

    protected onRemove(index: number): void {
        if (
            this.imagesPreview &&
            index >= 0 &&
            index < this.imagesPreview.length
        ) {
            this.imagesPreview.splice(index, 1);
        }

        const currentFiles = this.formImages.get('images')?.value as File[];
        if (currentFiles && currentFiles.length > index) {
            const updatedFiles = currentFiles.filter((_, i) => i !== index);

            this.formImages
                .get('images')
                ?.setValue(updatedFiles.length ? updatedFiles : null);

            // reset form if no file
            if (updatedFiles.length === 0) {
                this.formImages.get('images')?.markAsPristine();
                this.formImages.get('images')?.markAsUntouched();
            }
        }
    }

    protected maxFilesLength(maxLength: number): ValidatorFn {
        const stringErr = `Error: maximum limit - ${maxLength} files for upload`;

        return ({ value }: AbstractControl) => {
            if (!value || !Array.isArray(value)) {
                return null;
            }

            return value.length > maxLength
                ? { maxLength: new TuiValidationError(stringErr) }
                : null;
        };
    }

    protected uploadFiles() {
        const formImages = this.formImages.get('images')?.value;

        this.imageService
            .uploadImages(
                this.oldImages ?? [],
                formImages ?? [],
                this.id,
                this.isDetailMode
            )
            .subscribe((response) => {
                this.oldImages = [];
                if (response.data) {
                    this.oldImages.push(...response.data);
                }
                this.imagesPreview = [];
            });
    }
}
