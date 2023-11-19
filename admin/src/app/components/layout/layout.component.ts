import { Component } from '@angular/core';
import { ActivatedRoute, ActivatedRouteSnapshot, NavigationEnd, Router } from '@angular/router';
import { filter, map } from 'rxjs';

@Component({
  selector: 'zine-layout',
  templateUrl: './layout.component.html',
  styleUrls: ['./layout.component.scss']
})
export class LayoutComponent {
  minimal = false;

  constructor(router:Router, route:ActivatedRoute) {
    router.events
    .pipe(
      filter(event => event instanceof NavigationEnd),
      map(() => route.snapshot),
      map(route => {
        while (route.firstChild) {
          route = route.firstChild;
        }
        return route;
      })
    )
    .subscribe((route: ActivatedRouteSnapshot) => {
      if(route.data['layout'] === 'empty') {
        this.minimal = true;
      }
    });
  }
}
