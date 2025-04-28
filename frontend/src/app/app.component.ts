import { Component, OnInit, ViewChild } from '@angular/core';
import { AutoCompleteModule } from 'primeng/autocomplete';
import { AccommodationService } from './services/accommodation.service';
import {
    Accommodation,
    CreateAccommodation,
    UpdateAccommodation,
} from './models/accommodation.model';
import { ButtonModule } from 'primeng/button';
import { TableModule } from 'primeng/table';
import { Dialog } from 'primeng/dialog';
import { InputTextModule } from 'primeng/inputtext';
import { FormControl, FormGroup, ReactiveFormsModule } from '@angular/forms';
import { CheckboxModule, Checkbox } from 'primeng/checkbox';
import { FileUpload } from 'primeng/fileupload';
import { CommonModule } from '@angular/common';
import { BadgeModule } from 'primeng/badge';
import { ToastModule } from 'primeng/toast';
import { MessageService } from 'primeng/api';

@Component({
    selector: 'app-root',
    imports: [
        AutoCompleteModule,
        TableModule,
        ButtonModule,
        Dialog,
        InputTextModule,
        Checkbox,
        ReactiveFormsModule,
        CheckboxModule,
        FileUpload,
        CommonModule,
        BadgeModule,
        ToastModule,
        CommonModule,
    ],
    providers: [MessageService],
    templateUrl: './app.component.html',
    styleUrl: './app.component.css',
})
export class AppComponent implements OnInit {
    accommodations!: Accommodation[];
    dialogCreate: boolean = false;
    dialogUpdate: boolean = false;

    customUpload = {
        progressbar: {
            height: '0px',
        },
    };

    files: File[] = [];
    isOldImage: boolean = false;
    idAccommodation: string = '';
    oldImageName: string = '';
    oldImageType: string = '';
    baseUrl: string = 'http://localhost:8080/uploads/';
    edittedImage: boolean = false;

    @ViewChild('fileUploader') fileUploader!: FileUpload;

    acommodationForm = new FormGroup({
        name: new FormControl(''),
        country: new FormControl(''),
        city: new FormControl(''),
        district: new FormControl(''),
        description: new FormControl(''),
        facilities: new FormGroup({
            wifi: new FormControl(false),
            airCondition: new FormControl(false),
            tv: new FormControl(false),
        }),
        googleMap: new FormControl(''),
        propertySurrounds: new FormGroup({
            restaurant: new FormControl(false),
            bar: new FormControl(false),
        }),
        rules: new FormControl(''),
    });

    constructor(private accommodationService: AccommodationService) {}

    ngOnInit() {
        this.accommodationService.getAccommodations().subscribe((response) => {
            this.accommodations = response.data;
        });
    }

    getMimeType(fileName: string): string {
        const extension = fileName.split('.').pop()?.toLowerCase();
        switch (extension) {
            case 'jpg':
            case 'jpeg':
                return 'image/jpeg';
            case 'png':
                return 'image/png';
            case 'gif':
                return 'image/gif';
            case 'bmp':
                return 'image/bmp';
            case 'webp':
                return 'image/webp';
            case 'svg':
                return 'image/svg+xml';
            default:
                return 'application/octet-stream';
        }
    }

    choose(event: any, callback: any) {
        callback();
    }

    onRemoveTemplatingFile(
        event: any,
        file: any,
        removeFileCallback: any,
        index: any
    ) {
        removeFileCallback(event, index);
        console.log('remove: ', this.files);
    }

    onClearTemplatingUpload(clear: any) {
        clear();
    }

    onRemoveOldImage() {
        this.oldImageName = '';
        this.oldImageType = '';
        this.isOldImage = false;
        this.edittedImage = true;
        console.log(this.edittedImage);
    }

    onSelectedFiles(event: any) {
        this.files = event.currentFiles;
        this.edittedImage = true;
        console.log(this.edittedImage);
    }

    async addDefaultImage(url: string, filename: string) {
        const response = await fetch(url);
        const blob = await response.blob();
        const file = new File([blob], filename, { type: blob.type });
        const objectURL = URL.createObjectURL(file);

        this.files.push(file);
    }

    showDialogCreate() {
        this.acommodationForm.reset();
        this.dialogCreate = true;
        this.isOldImage = false;
        this.oldImageName = '';
        this.oldImageType = '';
        this.files = [];
        this.fileUploader.clear()
    }

    showDialogUpdate(accommodation: Accommodation) {
        this.acommodationForm.reset();
        this.files = [];
        this.fileUploader.clear()
        this.acommodationForm.setValue({
            name: accommodation.name,
            country: accommodation.country,
            city: accommodation.city,
            district: accommodation.district,
            description: accommodation.description,
            facilities: {
                wifi: accommodation.facilities.wifi,
                airCondition: accommodation.facilities.airCondition,
                tv: accommodation.facilities.tv,
            },
            googleMap: accommodation.googleMap,
            propertySurrounds: {
                restaurant: accommodation.propertySurrounds.restaurant,
                bar: accommodation.propertySurrounds.bar,
            },
            rules: accommodation.rules,
        });
        this.dialogUpdate = true;
        this.addDefaultImage(this.baseUrl + accommodation.image, 'default.jpg');
        this.oldImageName = accommodation.image;
        this.oldImageType = this.getMimeType(accommodation.image);
        this.isOldImage = true;
        this.idAccommodation = accommodation.id;
    }

    handleCreate() {
        const accommodation: CreateAccommodation = {
            name: this.acommodationForm.value.name || '',
            city: this.acommodationForm.value.city || '',
            country: this.acommodationForm.value.country || '',
            district: this.acommodationForm.value.district || '',
            description: this.acommodationForm.value.description || '',
            image: this.files,
            facilities: {
                wifi: this.acommodationForm.value.facilities?.wifi || false,
                airCondition:
                    this.acommodationForm.value.facilities?.airCondition ||
                    false,
                tv: this.acommodationForm.value.facilities?.tv || false,
            },
            googleMap: this.acommodationForm.value.googleMap || '',
            propertySurrounds: {
                restaurant:
                    this.acommodationForm.value.propertySurrounds?.restaurant ||
                    false,
                bar:
                    this.acommodationForm.value.propertySurrounds?.bar || false,
            },
            rules: this.acommodationForm.value.rules || '',
        };

        this.accommodationService
            .createAccommodation(accommodation)
            .subscribe((response) => {
                this.accommodations.push(response.data);
                this.acommodationForm.reset();
            });
        this.dialogCreate = false;
    }

    handleUpdate() {
        const accommodation: UpdateAccommodation = {
            name: this.acommodationForm.value.name || '',
            city: this.acommodationForm.value.city || '',
            country: this.acommodationForm.value.country || '',
            district: this.acommodationForm.value.district || '',
            description: this.acommodationForm.value.description || '',
            image: [],
            facilities: {
                wifi: this.acommodationForm.value.facilities?.wifi || false,
                airCondition:
                    this.acommodationForm.value.facilities?.airCondition ||
                    false,
                tv: this.acommodationForm.value.facilities?.tv || false,
            },
            googleMap: this.acommodationForm.value.googleMap || '',
            propertySurrounds: {
                restaurant:
                    this.acommodationForm.value.propertySurrounds?.restaurant ||
                    false,
                bar:
                    this.acommodationForm.value.propertySurrounds?.bar || false,
            },
            rules: this.acommodationForm.value.rules || '',
            id: this.idAccommodation,
        };

        if (this.files.length > 0) {
            accommodation.image = this.files;
        }

        this.accommodationService
            .updateAccommodation(accommodation)
            .subscribe((response) => {
                console.log('Accommodation updated:', response);
                this.accommodations = this.accommodations.map(
                    (accommodation) => {
                        if (accommodation.id === response.data.id) {
                            return response.data;
                        } else {
                            return accommodation;
                        }
                    }
                );
                this.acommodationForm.reset();
                this.dialogUpdate = false;
            });
    }

    handleDeleteAccommodation(id: string) {
        this.accommodationService.deleteAccommodation(id).subscribe((response) => {
            this.accommodations = this.accommodations.filter(
                (accommodation) => accommodation.id !== id
            );
        });
    }
}
