import { Component } from '@angular/core';
import { ThemeService } from '../theme.service';

@Component({
  selector: 'zine-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.scss']
})
export class HeaderComponent {
  logoColor = '#000'
  constructor(theme: ThemeService) {
    theme.change.subscribe(v => {
      if (v === 'dark') {
        this.logoColor = '#FFF';
      } else {
        this.logoColor = '#000'
      }
    })
  }
}
