import {
  Column,
  PrimaryGeneratedColumn,
} from 'typeorm';

export class TaskEntity {
  @PrimaryGeneratedColumn()
  id: number;

  @Column()
  title: string;

  @Column()
  description: string;
}
