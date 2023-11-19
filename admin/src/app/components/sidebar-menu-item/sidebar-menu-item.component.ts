import { Component, ContentChildren, EventEmitter, HostBinding, HostListener, Input, Output, QueryList } from '@angular/core';

@Component({
  selector: 'zine-sidebar-menu-item',
  templateUrl: './sidebar-menu-item.component.html',
  styleUrls: ['./sidebar-menu-item.component.scss']
})
export class SidebarMenuItemComponent {
  @HostBinding('class.child')
  @Input() child: boolean = false;
  @Input() href: string = '';
  @Input() icon: string = '';
  @Input() text: string = '';
  @Input() expanded = false;

  @HostBinding('class.selected')
  @Input() selected = false;
  
  @Output() activate: EventEmitter<SidebarMenuItemComponent> = new EventEmitter(true);

  @ContentChildren(SidebarMenuItemComponent) children:QueryList<SidebarMenuItemComponent> = new QueryList();

  @HostListener('click', ['$event'])
  onMenuItemClick(event: MouseEvent) {
    if (!this.children.length) {
      this.activate.emit(this);
    }
    event.stopPropagation();
    if (this.children.length) {
      this.expanded = !this.expanded;
    }
  }

  ngAfterContentInit() {
    this.children.forEach(mi => mi.child = true);
  }
}
