export interface ITopic {
  id: number;
  slug: string;
  createdAt: number;
  updatedAt: number;
}

export interface ITopics {
  totalCount: number;
  nodes: ITopic[];
}