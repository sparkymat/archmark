import React from 'react';
import { Routes, Route } from 'react-router-dom';
import NewBookmark from '../NewBookmark';
import Search from '../Search';
import ArchivedList from '../ArchivedList';

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
        home
      </a>
      <a
        className="uk-link-muted uk-margin-left"
        style={{ borderBottom: '1px dashed #999999' }}
        href="/#/bookmark"
      >
        new
      </a>
      <a
        className="uk-link-muted uk-margin-left"
        style={{ borderBottom: '1px dashed #999999' }}
        href="/#/deleted"
      >
        deleted
      </a>
    </div>
    <Routes>
      <Route index element={<Search />} />
      <Route path="/page/:page" element={<Search />} />
      <Route path="/bookmark" element={<NewBookmark />} />
      <Route path="/search" element={<Search />} />
      <Route path="/search/:query" element={<Search />} />
      <Route path="/search/:query/page/:page" element={<Search />} />
      <Route path="/deleted" element={<ArchivedList />} />
      <Route path="/deleted/page/:page" element={<ArchivedList />} />
    </Routes>
  </div>
);
export default App;
