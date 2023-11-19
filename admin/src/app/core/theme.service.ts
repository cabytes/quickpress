import { EventEmitter, Injectable, Output } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class ThemeService {
  theme: string = 'light'
  @Output() change = new EventEmitter<string>();
  constructor() {
    this.change.emit(this.theme);
  }
  setTheme(theme: string) {
    this.theme = theme;
    this.change.emit(theme);
  }
}
