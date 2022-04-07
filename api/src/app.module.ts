import { Module } from '@nestjs/common';
import { TypeOrmModule } from '@nestjs/typeorm'; 
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { TaskEntity } from './lib';
import { TasksModule } from './tasks/tasks.module';

@Module({
  imports: [
    TasksModule,
    TypeOrmModule.forRoot({
      type: 'postgres',
      host: process.env.POSTGRES_HOST ?? 'localhost',
      port: +process.env.POSTGRES_PORT ?? 5432,
      username: process.env.POSTGRES_USERNAME ?? 'postgres',
      password: process.env.POSTGRES_PASSWORD ?? 'p@ssw0rd',
      database: process.env.POSTGRES_DATABASE ?? 'peercode',
      entities: [
        TaskEntity,
      ],
      synchronize: true,
      ssl: false,
    }),
  ],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule {}
