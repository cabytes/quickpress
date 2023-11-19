import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { LoaderComponent } from './app.component';
import { LayoutComponent } from './components/layout/layout.component';
import { AuthComponent } from './components/auth/auth.component';
import { HeaderComponent } from './components/header/header.component';
import { DashboardComponent } from './components/dashboard/dashboard.component';
import { SidebarComponent } from './components/sidebar/sidebar.component';
import { LogoComponent } from './components/logo/logo.component';
import { PostsComponent } from './components/posts/posts.component';
import { MediaComponent } from './components/media/media.component';
import { UsersComponent } from './components/users/users.component';
import { SettingsComponent } from './components/settings/settings.component';
import { PagesComponent } from './components/pages/pages.component';
import { CommentsComponent } from './components/comments/comments.component';
import { SidebarMenuItemComponent } from './components/sidebar-menu-item/sidebar-menu-item.component';
import { SidebarMenuComponent } from './components/sidebar-menu/sidebar-menu.component';
import { AvatarComponent } from './components/avatar/avatar.component';
import { HTTP_INTERCEPTORS, HttpClientModule } from '@angular/common/http';
import { ApiInterceptor } from './core/ApiInterceptor';

@NgModule({
  declarations: [
    LoaderComponent,
    LayoutComponent,
    AuthComponent,
    HeaderComponent,
    DashboardComponent,
    SidebarComponent,
    LogoComponent,
    PostsComponent,
    MediaComponent,
    UsersComponent,
    SettingsComponent,
    PagesComponent,
    CommentsComponent,
    SidebarMenuItemComponent,
    SidebarMenuComponent,
    AvatarComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
  ],
  providers: [
    { provide: HTTP_INTERCEPTORS, useClass: ApiInterceptor, multi: false },
  ],
  bootstrap: [LoaderComponent]
})
export class AppModule { }
