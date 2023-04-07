import React, { useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { useParams } from 'react-router-dom';
import {
  selectBookmarks,
  selectLoading,
  selectPageSize,
} from '../../features/BookmarksList/selects';
import { AppDispatch } from '../../store';
import fetchBookmarksList from '../../features/BookmarksList/fetchBookmarksList';

const BookmarksList = () => {
  const dispatch = useDispatch<AppDispatch>();

  const { page: pageNumberString } = useParams();

  let pageNumber = 1;

  if (pageNumberString) {
    try {
      pageNumber = parseInt(pageNumberString, 10);
    } catch (error) {
      console.log(error);
    }
  }

  const pageSize = useSelector(selectPageSize);

  useEffect(() => {
    dispatch(
      fetchBookmarksList({
        page_number: pageNumber,
        page_size: pageSize,
      }),
    );
  }, [dispatch, pageNumber, pageSize]);

  // TODO: Add loading spinner
  // const loading = useSelector(selectLoading);
  const bookmarks = useSelector(selectBookmarks);

  return (
    <div className="uk-container">
      {bookmarks &&
        bookmarks.map((b, i) => (
          <div className="uk-flex uk-flex-column uk-margin-bottom">
            <a
              className="uk-link-muted uk-text-large"
              style={{ fontWeight: 300, color: 'white' }}
              href={b.url}
              target="_blank"
              rel="noreferrer"
            >
              <span className="uk-text-default">
                {(pageNumber - 1) * pageSize + i + 1}.{' '}
              </span>
              <span>{b.title}</span>
            </a>
            <div className="uk-flex uk-flex-row">
              <a href="#" className="uk-link-muted">
                cached
              </a>
              <span className="uk-margin-small-left uk-margin-small-right">
                ⚬
              </span>
              <a href="#" className="uk-link-muted">
                add to reading list
              </a>
            </div>
          </div>
        ))}
    </div>
  );
};

export default BookmarksList;
