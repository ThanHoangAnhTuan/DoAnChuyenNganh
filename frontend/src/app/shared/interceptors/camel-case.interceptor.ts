import { Injectable } from '@angular/core';
import {
    HttpRequest,
    HttpHandler,
    HttpEvent,
    HttpInterceptor,
    HttpResponse,
} from '@angular/common/http';
import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';

@Injectable()
export class CaseConverterInterceptor implements HttpInterceptor {
    intercept(
        request: HttpRequest<unknown>,
        next: HttpHandler
    ): Observable<HttpEvent<unknown>> {
        // convert request from camelCase to snake_case
        console.log('request: ', request);
        const convertedRequest = this.convertRequest(request);
        console.log(convertedRequest);

        return next.handle(convertedRequest).pipe(
            map((event: HttpEvent<any>) => {
                if (event instanceof HttpResponse) {
                    // convert response from snake_case to camelCase
                    return event.clone({
                        body: this.toCamelCase(event.body),
                    });
                }
                return event;
            })
        );
    }

    private convertRequest(request: HttpRequest<any>): HttpRequest<any> {
        if (request.body instanceof FormData) {
            // Xử lý đặc biệt cho FormData
            const convertedFormData = new FormData();

            // Lấy tất cả các cặp key-value từ FormData gốc
            request.body.forEach((value, key) => {
                // Chuyển key từ camelCase sang snake_case
                const snakeKey = key.replace(/([A-Z])/g, '_$1').toLowerCase();

                // Nếu value là File hoặc Blob, giữ nguyên
                if (value instanceof File) {
                    convertedFormData.append(
                        snakeKey,
                        value,
                        (value as File).name
                    );
                } else if ((value as any) instanceof Blob) {
                    convertedFormData.append(snakeKey, value);
                } else if (typeof value === 'object' && value !== null) {
                    convertedFormData.append(
                        snakeKey,
                        JSON.stringify(this.toSnakeCase(value))
                    );
                }
                // Các trường hợp khác (string, number, boolean)
                else {
                    convertedFormData.append(snakeKey, value);
                }
            });

            return request.clone({ body: convertedFormData });
        } else if (request.body) {
            return request.clone({
                body: this.toSnakeCase(request.body),
            });
        }
        return request;
    }

    // conver object from camelCase to snake_case
    private toSnakeCase(obj: any): any {
        if (obj === null || obj === undefined || typeof obj !== 'object') {
            return obj;
        }

        if (Array.isArray(obj)) {
            return obj.map((item) => this.toSnakeCase(item));
        }

        return Object.keys(obj).reduce((acc, key) => {
            const snakeKey = key.replace(/([A-Z])/g, '_$1').toLowerCase();
            acc[snakeKey] = this.toSnakeCase(obj[key]);
            return acc;
        }, {} as any);
    }

    // convert object from snake_case to camelCase
    private toCamelCase(obj: any): any {
        if (obj === null || obj === undefined || typeof obj !== 'object') {
            return obj;
        }

        if (Array.isArray(obj)) {
            return obj.map((item) => this.toCamelCase(item));
        }

        return Object.keys(obj).reduce((acc, key) => {
            const camelKey = key.replace(/_([a-z])/g, (_, letter) =>
                letter.toUpperCase()
            );
            acc[camelKey] = this.toCamelCase(obj[key]);
            return acc;
        }, {} as any);
    }
}
