import React, { FC, useState, useEffect, useCallback } from 'react';
import { withRouter, RouteComponentProps } from 'react-router';
import './style.scss';

import { IGraphqlData, IList, IQuestion } from '../../model';
import Search from '../../components/Search';
import Questions from '../../components/Questions';
import Pagination from '../../components/Pagination';

const SearchPage: FC<RouteComponentProps> = ({ location, history }) => {
  const params = new URLSearchParams(location.search);
  const [keyword, setKeyword] = useState(params.get('keyword') || '');
  const [current, setCurrent] = useState(1);
  const [questions, setQuestions] = useState<IList<IQuestion>>({ totalCount: 0, nodes: [] });

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
            questions(page: ${current}, size: 20, keyword: "${keyword}")  {
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
    <main className="search-page container">
      <Search
        placeholder="回车搜索"
        value={keyword}
        onChange={setKeyword}
        onEnter={fetchQuestions}
      />
      <Questions data={questions} />
      <Pagination
        current={current}
        size={20}
        total={questions.totalCount}
        onChange={(current) => setCurrent(current)}
      />
    </main>
  )
}

export default withRouter(SearchPage);