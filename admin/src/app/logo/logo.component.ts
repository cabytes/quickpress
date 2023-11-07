import { Component, Input } from '@angular/core';

@Component({
  selector: 'zine-logo',
  templateUrl: './logo.component.html',
  styleUrls: ['./logo.component.scss']
})
export class LogoComponent {
  @Input() color: string = '#3F68D5';
}
