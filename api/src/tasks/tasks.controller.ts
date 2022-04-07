import { Controller, Get } from '@nestjs/common';
import { TasksService } from './tasks.service';

@Controller('tasks')
export class TasksController {
  constructor(
    private readonly tasksService: TasksService,
  ) { }

  @Get()
  public getTasks() {
    return this.tasksService.getTasks();
  }
}
