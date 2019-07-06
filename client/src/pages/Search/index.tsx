import React, { FC, useState } from 'react';
import { withRouter, RouteComponentProps } from 'react-router';

import Search from '../../components/Search'

const SearchPage: FC<RouteComponentProps> = ({ location }) => {
  const params = new URLSearchParams(location.search);
  const [keyword, setKeyword] = useState(params.get('keyword') || '');

  return (
    <main>
      搜索页
      <Search
        placeholder="回车搜索"
        value={keyword}
        onChange={setKeyword}
      />
    </main>
  )
}

export default withRouter(SearchPage);