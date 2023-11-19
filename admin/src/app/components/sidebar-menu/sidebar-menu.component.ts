import { Component, ContentChildren, QueryList } from '@angular/core';
import { SidebarMenuItemComponent } from '../sidebar-menu-item/sidebar-menu-item.component';

@Component({
  selector: 'zine-sidebar-menu',
  template: '<ng-content></ng-content>',
  styles: [':host { margin-top: 5px; display: block; }']
})
export class SidebarMenuComponent {

  @ContentChildren(SidebarMenuItemComponent) children: QueryList<SidebarMenuItemComponent> = new QueryList();

  ngAfterContentInit() {
    this.children.forEach(mi => {
      mi.activate.subscribe((current) => {
        this.children.forEach(m => {
          if (m !== current) {
            m.expanded = false;
            m.selected = false;
          } else {
            m.selected = true;
          }
        })
      })
    });
  }
}
