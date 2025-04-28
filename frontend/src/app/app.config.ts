import { ApplicationConfig, provideZoneChangeDetection } from '@angular/core';
import { provideRouter } from '@angular/router';
import { provideAnimationsAsync } from '@angular/platform-browser/animations/async';
import { providePrimeNG } from 'primeng/config';
import { routes } from './app.routes';

import Aura from '@primeng/themes/aura';
import {
    HTTP_INTERCEPTORS,
    provideHttpClient,
    withInterceptors,
    withInterceptorsFromDi,
} from '@angular/common/http';
import { CaseConverterInterceptor } from './shared/interceptors/camel-case.interceptor';

export const appConfig: ApplicationConfig = {
    providers: [
        provideZoneChangeDetection({ eventCoalescing: true }),
        provideRouter(routes),
        provideAnimationsAsync(),
        provideHttpClient(),
        provideHttpClient(withInterceptorsFromDi()),
        {provide: HTTP_INTERCEPTORS, useClass: CaseConverterInterceptor, multi: true},
        providePrimeNG({
            theme: {
                preset: Aura,
                options: {
                    darkModeSelector: '.dark-mode',
                },
            },
        }),
    ],
};
