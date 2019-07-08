import React, { FC, useCallback, useEffect, useState } from 'react';
import { withRouter, RouteComponentProps } from 'react-router';
import './style.scss';

import { IGraphqlData, IQuestion } from '../../model';
import Markdown from '../../components/Markdown';
import Editor from '../../components/Editor';

const AnswerEditPage: FC<RouteComponentProps> = ({ match, history, location }) => {
  const [question, setQuestion] = useState<IQuestion>({} as IQuestion);
  const [content, setContent] = useState('');
  const { id, ansId } = match.params as { id: number, ansId: number };
  const isCreate = /create/.test(location.pathname);

  useEffect(() => {
    fetchQuestion();
  }, [])

  const fetchQuestion = useCallback(() => {

    let query = `
        question(id: ${id})  {
          id
          title
          summary
          content
          difficulty
          createdAt
          updatedAt
        }`

    if (!isCreate) {
      query += `
        answer(id: ${ansId}) {
          id
          content
          createdAt
          updatedAt
        }`
    }

    fetch('/graphql', {
      method: 'POST',
      mode: 'cors',
      headers: {
        ContentType: 'application/json',
      },
      body: JSON.stringify({
        query: `{${query}}`,
      })
    })
      .then<IGraphqlData>((res) => res.json())
      .then(({ data }) => {
        const { question, answer } = data;
        setQuestion(question);
        setContent(answer.content);
      })
  }, [])

  const update = useCallback(() => {
    fetch('/graphql', {
      method: 'POST',
      mode: 'cors',
      headers: {
        ContentType: 'application/json',
      },
      body: JSON.stringify({
        query: `
          mutation($id: Int!, $content: String!) {
            updateAnswer(id: $id, content: $content) {
              success
            }
        }`,
        variables: { id, content }
      })
    })
      .then(() => {
        history.push(`/question/${id}`)
      }).catch(() => { })
  }, [content]);

  const create = useCallback(() => {
    fetch('/graphql', {
      method: 'POST',
      mode: 'cors',
      headers: {
        ContentType: 'application/json',
      },
      body: JSON.stringify({
        query: `
          mutation($quId: Int! ,$content: String!) {
            createAnswer(quId: $quId, content: $content) {
              id
            }
          }`,
        variables: { quId: id, content }
      })
    })
      .then(() => {
        history.push(`/question/${id}`)
      }).catch(() => { })
  }, [content]);

  const handleSubmit = useCallback(() => {
    if (isCreate) {
      create();
    } else {
      update();
    }
  }, [content])

  return (
    <main className="answer-page">
      <h1>{question.title}</h1>
      {question.content && <Markdown text={question.content} />}
      <button onClick={handleSubmit} >保存</button>
      <article className="editor-wrap">
        <Editor value={content} onChange={setContent} />
      </article>
    </main>
  )
}

export default withRouter(AnswerEditPage);