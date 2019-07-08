import { ITopic } from "./topic";
import { IQuestion } from './question';
import { IAnswer } from './answer';

export interface IList<T> {
  totalCount: number;
  nodes: T[];
}

export interface IGraphqlData {
  data: {
    question: IQuestion,
    questions: IList<IQuestion>,
    topic: ITopic,
    topics: IList<ITopic>,
    answer: IAnswer,
    answers: IList<IAnswer>,
  }
}