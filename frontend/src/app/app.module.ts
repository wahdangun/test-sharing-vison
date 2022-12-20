import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { ListArticleComponent } from './pages/list-article/list-article.component';
import { FormArticleComponent } from './pages/form-article/form-article.component';

@NgModule({
  declarations: [
    AppComponent,
    ListArticleComponent,
    FormArticleComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
