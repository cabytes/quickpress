import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { LoaderComponent } from './app.component';
import { LayoutComponent } from './layout/layout.component';
import { AuthComponent } from './auth/auth.component';
import { HeaderComponent } from './header/header.component';
import { DashboardComponent } from './dashboard/dashboard.component';
import { SidebarComponent } from './sidebar/sidebar.component';
import { LogoComponent } from './logo/logo.component';
import { PostsComponent } from './posts/posts.component';
import { MediaComponent } from './media/media.component';
import { UsersComponent } from './users/users.component';
import { SettingsComponent } from './settings/settings.component';
import { PagesComponent } from './pages/pages.component';
import { CommentsComponent } from './comments/comments.component';

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
    CommentsComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule
  ],
  providers: [],
  bootstrap: [LoaderComponent]
})
export class AppModule { }
