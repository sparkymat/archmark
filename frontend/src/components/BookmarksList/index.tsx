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

  const loading = useSelector(selectLoading);
  const bookmarks = useSelector(selectBookmarks);

  return (
    <div>
      {bookmarks &&
        bookmarks.map(b => (
          <div>
            <a className="uk-link-muted" href={b.url} target="_blank">
              {b.title}
            </a>
          </div>
        ))}
    </div>
  );
};

export default BookmarksList;
