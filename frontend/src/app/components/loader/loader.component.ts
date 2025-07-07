import { Component, Input } from '@angular/core';
import { TuiLoader } from '@taiga-ui/core';

@Component({
    selector: 'app-loader',
    imports: [TuiLoader],
    templateUrl: './loader.component.html',
    styleUrl: './loader.component.scss',
})
export class LoaderComponent {
    @Input() size: 's' | 'm' | 'l' | 'xl' = 'l';
    @Input() text: string = '';
    @Input() overlay: boolean = true;
    @Input() centered: boolean = true;
    @Input() transparent: boolean = false;
    @Input() inheritColor: boolean = false;
}
