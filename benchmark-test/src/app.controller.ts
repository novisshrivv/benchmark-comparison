import { Controller, Get, Res } from '@nestjs/common';
import { AppService } from './app.service';
import { AxiosResponse } from 'axios';
import { Observable } from 'rxjs';
import {Response} from 'express'


function cpuIntensiveTask() {
  // Simulate a CPU-intensive task
  let sum = 0;
  for (let i = 0; i < 100000000; i++) {
    sum += i;
  }
}

@Controller()
export class AppController {
  constructor(private readonly appService: AppService) {}

  @Get()
  getHello(): Observable<any[]> {
    return this.appService.getHello();
  }

  @Get('/cpu-intensive')
  async handleRequest(@Res() res: Response) {
    const start = new Date();
    cpuIntensiveTask();
    const elapsed = new Date().getTime() - start.getTime();
    res.send(`CPU-intensive task completed in ${elapsed} milliseconds\n`);
  }

}
