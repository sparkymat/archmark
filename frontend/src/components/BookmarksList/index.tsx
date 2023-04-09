import React from 'react';
import moment from 'moment';
import Bookmark from '../../models/Bookmark';
import URLDisplay from '../URLDisplay';
import Paginator from '../Paginatior';

interface BookmarksListProps {
  bookmarks: Bookmark[];
  pageNumber: number;
  pageSize: number;
  totalCount: number;
}

const BookmarksList = ({
  bookmarks,
  pageNumber,
  pageSize,
  totalCount,
}: BookmarksListProps) => (
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
    {bookmarks && bookmarks.length > 0 && (
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

export default BookmarksList;
