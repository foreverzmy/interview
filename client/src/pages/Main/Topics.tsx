import React, { FC, useEffect, useState } from 'react';

import { Link } from 'react-router-dom';

import Tag from '../../components/Tag';

import { IGraphqlData, ITopic } from '../../model';

const Topics: FC<{}> = () => {
  const [topics, setTopics] = useState<ITopic[]>([]);

  useEffect(() => {
    fetch('/graphql', {
      method: 'POST',
      mode: 'cors',
      headers: {
        ContentType: 'application/json',
      },
      body: JSON.stringify({
        query: `
          {
            topics {
              totalCount
              nodes{
                id
                slug
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
        const { topics } = data;
        const { nodes } = topics;

        setTopics(nodes);
      })
  }, [])

  return (
    <article className="topics">
      {topics.map(topic => {
        return (<Link key={topic.id} to={`topic/${topic.id}`}><Tag >{topic.slug}</Tag></Link>)
      })}
    </article>
  )
}

export default Topics;