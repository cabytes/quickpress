import { NgModule, inject } from '@angular/core';
import { ActivatedRouteSnapshot, RouterModule, RouterStateSnapshot, Routes } from '@angular/router';
import { LayoutComponent } from './components/layout/layout.component';
import { AuthComponent } from './components/auth/auth.component';
import { DashboardComponent } from './components/dashboard/dashboard.component';
import { PostsComponent } from './components/posts/posts.component';
import { MediaComponent } from './components/media/media.component';
import { UsersComponent } from './components/users/users.component';
import { SettingsComponent } from './components/settings/settings.component';
import { PagesComponent } from './components/pages/pages.component';
import { CommentsComponent } from './components/comments/comments.component';
import { Auth } from './core/Auth';

const routes: Routes = [
  {
    path: '',
    component: LayoutComponent,
    canActivate: [function(route: ActivatedRouteSnapshot, state: RouterStateSnapshot) {
      return inject(Auth).check();
    }],
    canActivateChild: [
      function(childRoute: ActivatedRouteSnapshot, state: RouterStateSnapshot) {
        return inject(Auth).state;
      }
    ],
    children: [
      { path: '', component: DashboardComponent },
      { path: 'login', component: AuthComponent, data: {layout: 'empty'} },
      { path: 'posts', component: PostsComponent },
      { path: 'pages', component: PagesComponent },
      { path: 'comments', component: CommentsComponent },
      { path: 'media', component: MediaComponent },
      { path: 'users', component: UsersComponent },
      { path: 'settings', component: SettingsComponent },
    ]
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes, {
    initialNavigation: 'enabledBlocking',
    paramsInheritanceStrategy: 'always'
  })],
  exports: [RouterModule]
})
export class AppRoutingModule { }
