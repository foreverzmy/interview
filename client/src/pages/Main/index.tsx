import React, { FC, useState, useCallback } from 'react';
import { withRouter, RouteComponentProps } from 'react-router';
import './style.scss';

import Search from '../../components/Search';
import Topics from './Topics';

const Main: FC<RouteComponentProps<any>> = ({ history }) => {
  const [keyword, setKeyword] = useState('');

  const handleSearch = useCallback(() => {
    history.push(`questions?keyword=${keyword}`)
  }, [keyword])

  return (
    <main className="main-page">
      <section className="banner">
        面试题
      </section>
      <Search
        placeholder="回车搜索"
        value={keyword}
        onChange={setKeyword}
        onEnter={handleSearch}
      />
      <Topics />
    </main>
  )
}

export default withRouter(Main);