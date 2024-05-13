import { Injectable } from '@nestjs/common';
import { HttpService } from '@nestjs/axios';
import { AxiosResponse } from 'axios';
import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';

@Injectable()
export class AppService {
  constructor(private httpService: HttpService) {}

  getHello(): Observable<any[]> {
    return this.httpService.get('https://jsonplaceholder.typicode.com/posts').pipe(
      map((response: AxiosResponse<any[]>) => response.data),
    );
  }
}
