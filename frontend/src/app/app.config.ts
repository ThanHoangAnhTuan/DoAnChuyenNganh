import { provideEventPlugins } from '@taiga-ui/event-plugins';
import { provideAnimations } from '@angular/platform-browser/animations';
import { ApplicationConfig, provideZoneChangeDetection } from '@angular/core';
import { provideRouter } from '@angular/router';
import { providePrimeNG } from 'primeng/config';
import Aura from '@primeng/themes/aura';

import { routes } from './app.routes';
import {
    HTTP_INTERCEPTORS,
    provideHttpClient,
    withInterceptorsFromDi,
} from '@angular/common/http';
import { provideAnimationsAsync } from '@angular/platform-browser/animations/async';
import { AuthInterceptor } from './shared/interceptors/auth.interceptor';
import { provideCharts, withDefaultRegisterables } from 'ng2-charts';

export const appConfig: ApplicationConfig = {
    providers: [
        provideCharts(withDefaultRegisterables()),
        provideAnimations(),
        provideZoneChangeDetection({ eventCoalescing: true }),
        provideRouter(routes),
        provideEventPlugins(),
        provideAnimationsAsync(),
        provideHttpClient(),
        provideHttpClient(withInterceptorsFromDi()),
        {
            provide: HTTP_INTERCEPTORS,
            useClass: AuthInterceptor,
            multi: true,
        },
        providePrimeNG({
            theme: {
                preset: Aura,
                options: {
                    darkModeSelector: false || 'none',
                },
            },
        }),
    ],
};
