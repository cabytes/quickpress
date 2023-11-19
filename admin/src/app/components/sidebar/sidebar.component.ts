import { Component } from '@angular/core';
import { ThemeService } from '../../core/theme.service';

@Component({
  selector: 'zine-sidebar',
  templateUrl: './sidebar.component.html',
  styleUrls: ['./sidebar.component.scss']
})
export class SidebarComponent {
  
  themeChangerChecked: boolean = true;

  constructor(private theme: ThemeService){
    if (theme.theme === 'light') {
      this.themeChangerChecked = true;
    }
  }

  onThemeChange(checked: boolean) {
    this.theme.setTheme(checked ? 'light' : 'dark')
  }
}
