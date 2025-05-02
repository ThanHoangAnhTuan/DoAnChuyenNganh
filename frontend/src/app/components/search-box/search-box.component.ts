import {ChangeDetectionStrategy, Component} from '@angular/core';
import {TuiComboBoxModule, TuiInputDateRangeModule} from '@taiga-ui/legacy';
import {TuiButton, TuiTextfieldOptionsDirective} from '@taiga-ui/core';
import {FormControl, FormsModule, ReactiveFormsModule} from '@angular/forms';
import {TuiDataListWrapper, TuiDataListWrapperComponent, TuiFilterByInputPipe} from '@taiga-ui/kit';

@Component({
  selector: 'app-search-box',
  imports: [
    TuiComboBoxModule,
    TuiTextfieldOptionsDirective,
    FormsModule,
    TuiDataListWrapperComponent,
    TuiDataListWrapper,
    TuiFilterByInputPipe,
    TuiInputDateRangeModule,
    ReactiveFormsModule,
    TuiButton
  ],
  standalone: true,
  templateUrl: './search-box.component.html',
  styleUrl: './search-box.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export default class SearchBoxComponent {
  protected country = null;
  protected readonly countries = [
    'Ho Chi Minh',
    'Ha Noi',
    'Da Nang',
    'Nha Trang',
    'Nha Trang',
    'Hue',
    'Dong Nai',
    'Vung Tau',
    'Da Lat',
    'Can Tho',
    'Ninh Binh',
    'Bac Linh',
    'Binh Dinh',
    'Binh Thuan',
    'Cao Bang',
  ]
  protected readonly control = new FormControl();

}
