import React, { FC, useCallback, useEffect, useState } from 'react';
import { withRouter, RouteComponentProps } from 'react-router';
import './style.scss';

import { IGraphqlData, IQuestion } from '../../model';
import Markdown from '../../components/Markdown';
import Editor from '../../components/Editor';

const AnswerEditPage: FC<RouteComponentProps> = ({ match, history }) => {
  const [question, setQuestion] = useState<IQuestion>({} as IQuestion);
  const [content, setContent] = useState('');
  const { id, ansId } = match.params as { id: number, ansId: number };

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
            }
            answer(id: ${ansId}) {
              id
              content
              createdAt
              updatedAt
            }
          }
        `
      })
    })
      .then<IGraphqlData>((res) => res.json())
      .then(({ data }) => {
        const { question, answer } = data;
        setQuestion(question);
        setContent(answer.content);
      })
  }, [])

  const handleSubmit = useCallback(() => {
    fetch('/graphql', {
      method: 'POST',
      mode: 'cors',
      headers: {
        ContentType: 'application/json',
      },
      body: JSON.stringify({
        query: `
          mutation($id: Int! ,$content: String!) {
            updateAnswer(id:$id, content: $content) {
              success
            }
          }`,
        variables: { id: ansId, content }
      })
    })
      .then(() => {
        history.push(`/question/${id}`)
      }).catch(() => { })
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