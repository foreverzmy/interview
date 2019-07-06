import React, { FC } from 'react';
import { HashRouter, Route, Redirect, Switch } from 'react-router-dom';

import Main from './pages/Main';
import SearchPage from './pages/Search';

const Routes: FC<{}> = () => (
  <HashRouter>
    <Switch>
      <Route path="/main" component={Main} />
      <Route path="/questions" component={SearchPage} />
      <Redirect from="/" to="/main" />
    </Switch>
  </HashRouter>
)

export default Routes;