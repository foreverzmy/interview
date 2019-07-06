import { ITopics } from "./topic";

export interface IGraphqlData {
  data: {
    topics: ITopics
  }
}