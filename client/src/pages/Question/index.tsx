import React, { FC, useCallback, useEffect, useState } from 'react';
import { withRouter, RouteComponentProps } from 'react-router';
import { Link } from 'react-router-dom';
import './style.scss';

import { IGraphqlData, IQuestion, IAnswer, IList, ITopic } from '../../model';
import Markdown from '../../components/Markdown';
import Tag from '../../components/Tag';

const QuestionPage: FC<RouteComponentProps> = ({ match }) => {
  const { id } = match.params as { id: number };
  const [current] = useState(1);
  const [question, setQuestion] = useState<IQuestion>({} as IQuestion);
  const [answers, setAnswers] = useState<IList<IAnswer>>({ totalCount: 0, nodes: [] });
  const [topics, setTopics] = useState<IList<ITopic>>({ totalCount: 0, nodes: [] });

  useEffect(() => {
    fetchQuestion();
  }, [])

  const fetchQuestion = useCallback(() => {

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
              topics {
                totalCount
                nodes {
                  id
                  slug
                  createdAt
                  updatedAt
                }
              }
            }
            answers(quId: ${id}, page: ${current}, size: 10) {
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
        const { topics } = question;
        setQuestion(question);
        setAnswers(answers);
        setTopics(topics);
      })
  }, [id, current])

  return (
    <main className="answer-page container">
      <h1 className="title">题目</h1>
      <h2>{question.title}</h2>
      {question.summary && <p>{question.summary}</p>}
      {question.content && <Markdown text={question.content} />}
      {topics.totalCount > 0 && (
        topics.nodes.map(topic => (
          <Link to={`/topic/${topic.id}`}><Tag key={topic.id}>{topic.slug}</Tag></Link>
        ))
      )}
      <h1 className="title answer">答案</h1>
      <article>
        {answers.nodes.map(answer => {
          return (
            <div key={answer.id}>
              <Markdown text={answer.content} />
            </div>
          )
        })}
        {answers.nodes.length === 0 && "暂无答案"}
        {/* <p>共{answers.totalCount}条回答</p> */}
      </article>
    </main>
  )
}

export default withRouter(QuestionPage);