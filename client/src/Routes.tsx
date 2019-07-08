import React, { FC } from 'react';
import { HashRouter, Route, Redirect, Switch } from 'react-router-dom';

import Header from './components/Header';
import Main from './pages/Main';
import SearchPage from './pages/Search';
import Question from './pages/Question';
import QuestionEdit from './pages/QuestionEdit';
import AnswerEdit from './pages/AnswerEdit'

const Routes: FC<{}> = () => (
  <HashRouter>
    <Route path="/" component={Header} />
    <Switch>
      <Route path="/main" component={Main} />
      <Route path="/questions" component={SearchPage} />
      <Route path="/question/:id/answer/create" component={AnswerEdit} />
      <Route path="/question/:id/answer/:ansId" component={AnswerEdit} />
      <Route path="/question/create" component={QuestionEdit} />
      <Route path="/question/edit/:id" component={QuestionEdit} />
      <Route path="/question/:id" component={Question} />
      <Redirect from="/" to="/main" />
    </Switch>
  </HashRouter>
)

export default Routes;