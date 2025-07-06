import { NgIf } from '@angular/common';
import { Component } from '@angular/core';
import { FormsModule } from '@angular/forms';

@Component({
    selector: 'app-chat-box',
    imports: [NgIf, FormsModule],
    templateUrl: './chat-box.component.html',
    styleUrl: './chat-box.component.scss',
})
export class ChatBoxComponent {
    isChatBoxOpen = false;

    toggleChat() {
        this.isChatBoxOpen = !this.isChatBoxOpen;
    }

    closeChat() {
        this.isChatBoxOpen = false;
    }
}
