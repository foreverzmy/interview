import React, { FC, useState, useEffect, useCallback } from 'react';
import { withRouter, RouteComponentProps } from 'react-router';
import './style.scss';

import { IGraphqlData, IList, IQuestion } from '../../model';
import Search from '../../components/Search';
import Questions from '../../components/Questions';

const SearchPage: FC<RouteComponentProps> = ({ location, history }) => {
  const params = new URLSearchParams(location.search);
  const [keyword, setKeyword] = useState(params.get('keyword') || '');
  const [questions, setQuestions] = useState<IList<IQuestion>>({ totalCount: 0, nodes: [] })

  useEffect(() => {
    fetchQuestions();
  }, [])

  const fetchQuestions = useCallback(() => {
    history.push({
      pathname: '/questions',
      search: `?keyword=${keyword}`
    })

    fetch('/graphql', {
      method: 'POST',
      mode: 'cors',
      headers: {
        ContentType: 'application/json',
      },
      body: JSON.stringify({
        query: `
          {
            questions(page: 1, size: 10, keyword: "${keyword}")  {
              totalCount
              nodes{
                id
                title
                summary
                content
                difficulty
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
        const { questions } = data;
        setQuestions(questions);
      })
  }, [keyword])

  return (
    <main className="search-page">
      <Search
        placeholder="回车搜索"
        value={keyword}
        onChange={setKeyword}
        onEnter={fetchQuestions}
      />
      <Questions data={questions} />
    </main>
  )
}

export default withRouter(SearchPage);