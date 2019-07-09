import { IList } from './data';
import { ITopic } from './topic';

export interface IQuestion {
  id: number,
  title: string,
  summary: string,
  content: string;
  difficulty: 1 | 2 | 3,
  createdAt: number,
  updatedAt: number;
  topics: IList<ITopic>;
}