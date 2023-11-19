import { Injectable } from "@angular/core";
import { HttpClient } from '@angular/common/http';
import { BehaviorSubject } from 'rxjs'

@Injectable({ providedIn: 'root' })
export class Auth {

    state: BehaviorSubject<boolean> = new BehaviorSubject(true);

    constructor(private http: HttpClient) {
        
    }

    login(user: string, pass: string) {
        
        const req = this.http.post(`/api/auth`, {
            method: 'POST',
            body: {
                user,
                pass,
            }
        });

        req.subscribe(res => {
            console.log(res)
        })
    }

    check() {
        return this.http.get(`/api/auth/me`).subscribe(res => console.log('Res:', res))
    }
}