import React from 'react';
import { Routes, Route } from 'react-router-dom';

import BookmarksList from '../BookmarksList';
import NewBookmark from '../NewBookmark';

const App = () => (
  <div>
    <div className="uk-padding-small uk-flex uk-flex-row uk-flex-center">
      {/* eslint-disable-next-line jsx-a11y/anchor-is-valid */}
      <a
        className="uk-link-muted uk-text-large"
        style={{ textDecoration: 'none', borderBottom: '1px dashed #999999' }}
        href="/#/"
      >
        archmark
      </a>
    </div>
    <Routes>
      <Route index element={<BookmarksList />} />
      <Route path="/page/:page" element={<BookmarksList />} />
      <Route path="/bookmark" element={<NewBookmark />} />
    </Routes>
  </div>
);
export default App;
