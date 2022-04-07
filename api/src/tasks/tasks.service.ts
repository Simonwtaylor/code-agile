import { Injectable } from '@nestjs/common';

@Injectable()
export class TasksService {
  public getTasks() {
    return [{
      id: 123,
      title: 'Task 1',
    }];
  }
}
