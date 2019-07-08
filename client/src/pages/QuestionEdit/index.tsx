import React, { FC, useCallback, useEffect, useState, ChangeEvent } from 'react';
import { withRouter, RouteComponentProps } from 'react-router';
import './style.scss';

import { IGraphqlData } from '../../model';
import { difficultyOptions } from '../../constant';
import Editor from '../../components/Editor';
import Search from '../../components/Search';

const QuestionEditPage: FC<RouteComponentProps> = ({ match, location, history }) => {
  const { id } = match.params as { id: number };

  const [title, setTitle] = useState('');
  const [difficulty, setDifficulty] = useState<1 | 2 | 3>(1);
  const [summary, setSummary] = useState('');
  const [content, setContent] = useState('');
  const isCreate = /create/.test(location.pathname);

  useEffect(() => {
    if (!isCreate) {
      fetchQuestion();
    }
  }, [])

  const fetchQuestion = useCallback(() => {
    const { id } = match.params as { id: number, ansId: number };

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
        const { question } = data;
        const { title, difficulty, summary, content } = question;
        setTitle(title);
        setDifficulty(difficulty);
        setSummary(summary);
        setContent(content);
      })
  }, []);

  const handleDifficultyChange = useCallback((evt: ChangeEvent<HTMLInputElement>) => {
    setDifficulty(parseInt(evt.target.value, 10) as 1 | 2 | 3);
  }, [])

  const create = useCallback(() => {
    fetch('/graphql', {
      method: 'POST',
      mode: 'cors',
      headers: {
        ContentType: 'application/json',
      },
      body: JSON.stringify({
        query: `
          mutation($title: String!,$summary:String!, $content: String!, $difficulty: Int!) {
            createQuestion(title: $title,summary: $summary, content: $content, difficulty: $difficulty) {
              id
            }
          }`,
        variables: { title, difficulty, summary, content }
      })
    })
      .then<IGraphqlData>((res) => res.json())
      .then(({ data }) => {
        const { createQuestion } = data;
        history.push(`/question/${createQuestion.id}`)
      }).catch(() => { })
  }, [title, difficulty, summary, content]);

  const update = useCallback(() => {
    fetch('/graphql', {
      method: 'POST',
      mode: 'cors',
      headers: {
        ContentType: 'application/json',
      },
      body: JSON.stringify({
        query: `
          mutation($id:Int!,$title: String!,$summary:String!, $content: String!, $difficulty: Int!) {
            updateQuestion(
              id: $id,
              title: $title,
              summary: $summary,
              content: $content,
              difficulty: $difficulty
            ) {
              success
            }
          }`,
        variables: { id, title, difficulty, summary, content }
      })
    })
      .then<IGraphqlData>((res) => res.json())
      .then(({ data }) => {
        const { updateQuestion } = data;
        if (updateQuestion.success) {
          history.push(`/question/${id}`)
        }
      }).catch(() => { })
  }, [title, difficulty, summary, content]);

  const handleSave = useCallback(() => {
    if (title === '') {
      alert('标题不能为空')
      return;
    }
    if (isCreate) {
      create();
    } else {
      update();
    }
  }, [title, isCreate, create, update]);

  return (
    <main className="question-edit-page container">
      <section>
        <label htmlFor="title">标题：</label>
        <Search placeholder="请输入标题" value={title} onChange={setTitle} />
      </section>
      <section>
        <label htmlFor="title">难度：</label>
        {difficultyOptions.map(option => (
          <label className="radio" key={option.level}>
            <input
              type="radio"
              name="difficulty"
              value={option.level}
              checked={difficulty === option.level}
              onChange={handleDifficultyChange}
            />
            {option.text}
          </label>
        ))}
      </section>
      <section>
        <label htmlFor="title">简介：</label>
        <textarea value={summary} onChange={evt => setSummary(evt.target.value)} />
      </section>
      <section>
        <label htmlFor="title">内容：</label>
        <Editor value={content} onChange={setContent} />
      </section>
      <section className="actions">
        <button className="submit-button" onClick={handleSave}>保存</button>
      </section>
    </main>
  )
}

export default withRouter(QuestionEditPage);