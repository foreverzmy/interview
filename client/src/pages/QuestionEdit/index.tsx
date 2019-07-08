import React, { FC, useCallback, useEffect, useState } from 'react';
import { withRouter, RouteComponentProps } from 'react-router';
import './style.scss';

import { IGraphqlData, IQuestion, IAnswer, IList } from '../../model';
import Markdown from '../../components/Markdown';

const QuestionPage: FC<RouteComponentProps> = ({ match }) => {
  const [question, setQuestion] = useState<IQuestion>({} as IQuestion);
  const [answers, setAnswers] = useState<IList<IAnswer>>({ totalCount: 0, nodes: [] });

  useEffect(() => {
    fetchQuestion();
  }, [])

  const fetchQuestion = useCallback(() => {
    const { id, ansId } = match.params as { id: number, ansId: number };

    fetch('/graphql', {
      method: 'POST',
      mode: 'cors',
      headers: {
        ContentType: 'application/json',
      },
      body: JSON.stringify({
        query: `
          {
            question(id: ${id})  {
              id
              title
              summary
              content
              difficulty
              createdAt
              updatedAt
            }
            answers(quId: ${id}, page: 1, size: 10) {
              totalCount
              nodes {
                id
                content
                createdAt
                updatedAt
              }
            }
          }
        `
      })
    })
      .then<IGraphqlData>((res) => res.json())
      .then(({ data }) => {
        const { question, answers } = data;
        setQuestion(question);
        setAnswers(answers);
      })
  }, [])

  return (
    <main className="question-page">
      <h1>{question.title}</h1>
      {question.content && <Markdown text={question.content} />}
      <article>
        {answers.nodes.map(answer => {
          return (
            <div key={answer.id}>
              <Markdown text={answer.content} />
            </div>
          )
        })}
        {answers.nodes.length === 0 && "暂无答案"}
        <p>共{answers.totalCount}条回答</p>
      </article>
    </main>
  )
}

export default withRouter(QuestionPage);