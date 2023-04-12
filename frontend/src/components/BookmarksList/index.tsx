import React, { ChangeEvent, useCallback } from 'react';
import moment from 'moment';
import Bookmark from '../../models/Bookmark';
import URLDisplay from '../URLDisplay';
import Paginator from '../Paginatior';

interface BookmarksListProps {
  bookmarks: Bookmark[];
  categories: string[];
  pageNumber: number;
  pageSize: number;
  totalCount: number;
  categoryModalOpen: boolean;
  categoryModalName: string;
  // eslint-disable-next-line react/no-unused-prop-types, react/require-default-props
  categoryModalBookmarkID?: string;
  hideCategoryModal(): void;
  showCategoryModal(_bookmarkID: string): void;
  categoryModalNameChanged(_val: string): void;
  categoryModalSubmitted(): void;
}

const BookmarksList = ({
  bookmarks,
  categories,
  pageNumber,
  pageSize,
  totalCount,
  categoryModalOpen,
  hideCategoryModal,
  showCategoryModal,
  categoryModalName,
  categoryModalNameChanged,
  categoryModalSubmitted,
}: BookmarksListProps) => {
  const nameChange = useCallback(
    (evt: ChangeEvent<HTMLInputElement>) => {
      categoryModalNameChanged(evt.target.value);
    },
    [categoryModalNameChanged],
  );

  return (
    <div className="uk-container">
      {(!bookmarks || bookmarks.length === 0) && (
        <div
          className="uk-padding uk-flex uk-flex-center"
          style={{ border: '1px dashed #999999' }}
        >
          <p className="uk-margin-remove">No bookmarks found</p>
        </div>
      )}
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
              <span className="uk-margin-small-left uk-margin-small-right">
                ⚬
              </span>
              {/* eslint-disable-next-line jsx-a11y/anchor-is-valid, jsx-a11y/click-events-have-key-events, jsx-a11y/no-static-element-interactions */}
              <a
                className="uk-link-muted"
                onClick={() => showCategoryModal(b.id)}
              >
                {b.category ? b.category : 'Uncategorized'}
              </a>
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
      <div
        id="update-category-modal"
        // eslint-disable-next-line react/no-unknown-property
        uk-modal=""
        className={categoryModalOpen ? 'uk-modal uk-open' : ''}
        style={
          categoryModalOpen
            ? { overscrollBehavior: 'contain', display: 'block' }
            : { display: 'none' }
        }
      >
        <div className="uk-modal-dialog uk-modal-body uk-margin-top uk-background-secondary uk-border-rounded">
          <div className="uk-flex uk-flex-row uk-flex-between uk-modal-title">
            <h2 className="uk-modal-title">Choose category</h2>
            {/* eslint-disable-next-line max-len */}
            {/* eslint-disable-next-line jsx-a11y/anchor-has-content, jsx-a11y/anchor-is-valid, jsx-a11y/click-events-have-key-events, jsx-a11y/no-static-element-interactions */}
            <a
              className="uk-icon-button"
              onClick={hideCategoryModal}
              // eslint-disable-next-line react/no-unknown-property
              uk-icon="close"
            />
          </div>
          <div className="uk-flex uk-flex-row">
            <input
              type="text"
              className="uk-input"
              value={categoryModalName}
              onChange={nameChange}
              placeholder="Type new category or choose below"
            />
            <button
              type="button"
              className="uk-button uk-margin-small-left"
              onClick={categoryModalSubmitted}
            >
              Update
            </button>
          </div>
          <table className="uk-table">
            <tbody>
              {categories &&
                categories.slice(0, 5).map(c => (
                  <tr className="uk-padding-bottom">
                    <td className="uk-text-large uk-padding-remove">
                      {/* eslint-disable-next-line jsx-a11y/anchor-is-valid, jsx-a11y/click-events-have-key-events, jsx-a11y/no-static-element-interactions */}
                      <a
                        className="uk-link-muted"
                        onClick={() => categoryModalNameChanged(c)}
                      >
                        {c}
                      </a>
                    </td>
                  </tr>
                ))}
            </tbody>
          </table>
        </div>
      </div>
    </div>
  );
};

export default BookmarksList;
