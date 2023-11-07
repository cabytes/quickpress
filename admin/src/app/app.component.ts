  import { Component, HostBinding } from '@angular/core';
import { ThemeService } from './theme.service';

@Component({
  selector: 'zine-admin-loader',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class LoaderComponent {
  
  @HostBinding("attr.data-theme") theme: string = 'light';

  constructor(theme: ThemeService) {
    theme.change.subscribe(v => this.theme = v);
  }
}
