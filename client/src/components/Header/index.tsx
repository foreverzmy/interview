import React, { FC } from 'react';
import { Link } from 'react-router-dom';
import { withRouter, RouteComponentProps } from 'react-router';
import './style.scss';

const Header: FC<RouteComponentProps> = ({ }) => {

  return (
    <header id="header">
      <nav>
        <Link to="/">首页</Link>
        <Link to="/questions">面试题</Link>
      </nav>
    </header>
  )
}

export default withRouter(Header);
