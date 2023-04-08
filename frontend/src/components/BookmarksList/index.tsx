import React, { useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { useParams } from 'react-router-dom';
import moment from 'moment';
import {
  selectBookmarks,
  selectPageSize,
  selectTotalCount,
} from '../../features/BookmarksList/selects';
import { AppDispatch } from '../../store';
import fetchBookmarksList from '../../features/BookmarksList/fetchBookmarksList';
import URLDisplay from '../URLDisplay';
import Paginator from '../Paginatior';

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
  const totalCount = useSelector(selectTotalCount);

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
      <div className="uk-flex uk-flex-row-reverse uk-margin-top uk-margin-bottom">
        {/* eslint-disable-next-line jsx-a11y/anchor-is-valid */}
        <a href="/#/bookmark" className="uk-link-muted">
          New bookmark
        </a>
      </div>
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
              {b.title ? <span>{b.title}</span> : <URLDisplay value={b.url} />}
            </a>
            <div className="uk-flex uk-flex-row">
              <span>added {moment(b.created_at).toNow(true)} ago</span>
              {/* eslint-disable-next-line jsx-a11y/anchor-is-valid */}
              {b.file_path && (
                <>
                  <span className="uk-margin-small-left uk-margin-small-right">
                    ⚬
                  </span>
                  <a
                    href={`/uploads/${b.file_path}`}
                    target="_blank"
                    rel="noreferrer"
                    className="uk-link"
                  >
                    cached
                  </a>
                </>
              )}
            </div>
          </div>
        ))}
      {bookmarks && (
        <div className="uk-flex uk-flex-row uk-flex-center uk-margin-bottom">
          <Paginator
            pageNumber={pageNumber}
            pageSize={pageSize}
            totalCount={totalCount}
          />
        </div>
      )}
    </div>
  );
};

export default BookmarksList;
