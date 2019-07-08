export interface IQuestion {
  id: number,
  title: string,
  summary: string,
  content: string;
  difficulty: 1 | 2 | 3,
  createdAt: number,
  updatedAt: number;
}