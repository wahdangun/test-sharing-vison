import { Inject, Injectable } from '@angular/core';
import { BehaviorSubject, Observable } from 'rxjs';
import { tap } from 'rxjs/operators';
import { API_URL } from '@core/api.token';
import { HttpClient, HttpErrorResponse } from '@angular/common/http';
import { AllInvoicesRes, Invoice } from '@core/models/Invoice';
import { createParamsFromObject } from '@core/utils/create-params-from-object';
import { InvoiceFilter } from '@core/models/InvoiceFilter';
import { AuthService } from '@core/services/auth.service';
import { MatSnackBar } from '@angular/material/snack-bar';
export interface InvoiceState {
  filters: InvoiceFilter | null;
}

export const initialState: InvoiceState = {
  filters: null,
};

@Injectable({ providedIn: 'root' })
export class InvoiceService {
  private invoiceState = new BehaviorSubject<InvoiceState>(initialState);

  get state(): InvoiceState {
    return this.invoiceState.getValue();
  }

  constructor(
    @Inject(API_URL) private api: string,
    private http: HttpClient,
     private snackBar: MatSnackBar,
    private auth: AuthService
  ) {}

  all(
    page: number,
    pageSize: number,
    filters?: Partial<InvoiceFilter>
  ): Observable<AllInvoicesRes> {
    let finalFilters: Partial<InvoiceFilter> = this.state.filters ?? {};

    if (filters) {
      finalFilters = filters;
    }

    if (this.auth.role === 'user') {
      finalFilters = { ...finalFilters, user: this.auth.state.userId! };
    }

    const params = createParamsFromObject(finalFilters)
      .append('page', page)
      .append('pageSize', pageSize);

    return this.http
      .get<AllInvoicesRes>(`${this.api}/article`, { params })
      .pipe(tap(x => console.log(x)));
  } 

  getById(id: number): Observable<Invoice> {
    var invoices = this.http.get<Invoice>(`${this.api}/article/${id}`);
    return invoices;
  }
  delete(id: number): void {
     this.http.delete(`${this.api}/article/${id}`)
     .subscribe(
      value => {
        this.openSnackBar('Invoice berhasil di hapus', 'success-snackbar');
      },
      error => {
        if (error instanceof HttpErrorResponse) {
          if (error.status === 401) {
            this.openSnackBar('Anda tidak memiliki akses', 'error-snackbar');
          }
          if (typeof error.error.message === 'string') {
            this.openSnackBar(error.error.message, 'error-snackbar');
          }
        }
      }
    );
  }
  openSnackBar(message: string, panelClass: string): void {
    this.snackBar.open(message, '', {
      duration: 3000,
      horizontalPosition: 'center',
      verticalPosition: 'top',
      panelClass,
    });
  }
  
}
function subscribe(arg0: (value: any) => void, arg1: (error: any) => void) {
  throw new Error('Function not implemented.');
}

