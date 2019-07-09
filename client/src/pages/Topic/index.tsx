import React, { FC, useCallback, useEffect, useState } from 'react';
import { withRouter, RouteComponentProps } from 'react-router';
import './style.scss';

import { IGraphqlData, IList, ITopic, IQuestion } from '../../model';
import Questions from '../../components/Questions';
import Pagination from '../../components/Pagination';

const TopicPage: FC<RouteComponentProps> = ({ match }) => {
  const { id } = match.params as { id: number };
  const [topic, setTopic] = useState<ITopic>({ slug: '' } as ITopic)
  const [questions, setQuestions] = useState<IList<IQuestion>>({ totalCount: 0, nodes: [] });
  const [current, setCurrnet] = useState(1);

  useEffect(() => {
    fetchData();
  }, [])

  const fetchData = useCallback(() => {
    fetch('/graphql', {
      method: 'POST',
      mode: 'cors',
      headers: {
        ContentType: 'application/json',
      },
      body: JSON.stringify({
        query: `
          {
            topic(id: ${id}) {
              id
              slug
            }
            questions(page: ${current}, size: 20, topicId: ${id}) {
              totalCount
              nodes{
                id
                title
                summary
                content
                difficulty
              }
            }
          }`
      })
    })
      .then<IGraphqlData>((res) => res.json())
      .then(({ data }) => {
        const { topic, questions } = data;
        setTopic(topic);
        setQuestions(questions);
      })
  }, [id]);

  return (
    <main className="topic-page">
      <h1>{topic.slug}</h1>
      <Questions data={questions} />
      <Pagination
        current={current}
        size={20}
        total={questions.totalCount}
        onChange={current => setCurrnet(current)}
      />
    </main>
  )
}

export default withRouter(TopicPage);
