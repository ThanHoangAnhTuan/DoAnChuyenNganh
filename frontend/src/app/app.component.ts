import { TuiRoot } from '@taiga-ui/core';
import { Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { ChatBoxComponent } from "./components/chat-box/chat-box.component";
@Component({
    selector: 'app-root',
    imports: [RouterOutlet, TuiRoot, ChatBoxComponent],
    templateUrl: './app.component.html',
    styleUrl: './app.component.scss',
})
export class AppComponent {}
