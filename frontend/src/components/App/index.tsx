import React from 'react';
import { Routes, Route } from 'react-router-dom';
import NewBookmark from '../NewBookmark';
import Search from '../Search';

const App = () => (
  <div>
    <div className="uk-padding-small uk-flex uk-flex-row uk-flex-center">
      <p className="uk-text-muted uk-text-lead uk-margin-remove">archmark</p>
    </div>
    <div className="uk-flex uk-flex-row uk-flex-center">
      <a
        className="uk-link-muted"
        style={{ borderBottom: '1px dashed #999999' }}
        href="/"
      >
        Home
      </a>
      <a
        className="uk-link-muted uk-margin-left"
        style={{ borderBottom: '1px dashed #999999' }}
        href="/#/bookmark"
      >
        New bookmark
      </a>
    </div>
    <Routes>
      <Route index element={<Search />} />
      <Route path="/page/:page" element={<Search />} />
      <Route path="/bookmark" element={<NewBookmark />} />
      <Route path="/search" element={<Search />} />
      <Route path="/search/:query" element={<Search />} />
      <Route path="/search/:query/page/:page" element={<Search />} />
    </Routes>
  </div>
);
export default App;
