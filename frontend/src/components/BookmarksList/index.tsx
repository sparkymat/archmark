import React, { ChangeEvent, useCallback } from 'react';
import Modal from 'react-modal';
import moment from 'moment';
import Bookmark from '../../models/Bookmark';
import URLDisplay from '../URLDisplay';
import Paginator from '../Paginatior';

interface BookmarksListProps {
  basePath: string;
  bookmarks: Bookmark[];
  categories: string[];
  pageNumber: number;
  pageSize: number;
  // eslint-disable-next-line react/require-default-props
  totalCount: number;
  allowCategoryChange: boolean;
  categoryModalOpen: boolean;
  categoryModalName: string;
  showArchiveButton: boolean;
  showUnarchiveButton: boolean;
  deleteModalOpen: boolean;
  // eslint-disable-next-line react/no-unused-prop-types, react/require-default-props
  categoryModalBookmarkID?: string;
  hideCategoryModal(): void;
  showCategoryModal(_bookmarkID: string): void;
  categoryModalNameChanged(_val: string): void;
  categoryModalSubmitted(): void;
  hideDeleteModal(): void;
  showDeleteModal(_bookmarkID: string): void;
  deleteModalSubmitted(): void;
  unarchiveClicked(_bookmarkID: string): void;
}

const BookmarksList = ({
  basePath,
  bookmarks,
  categories,
  pageNumber,
  pageSize,
  totalCount,
  allowCategoryChange,
  categoryModalOpen,
  deleteModalOpen,
  hideCategoryModal,
  showCategoryModal,
  hideDeleteModal,
  showDeleteModal,
  categoryModalName,
  categoryModalNameChanged,
  categoryModalSubmitted,
  showArchiveButton,
  showUnarchiveButton,
  deleteModalSubmitted,
  unarchiveClicked,
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
              {allowCategoryChange && (
                // eslint-disable-next-line jsx-a11y/anchor-is-valid, jsx-a11y/click-events-have-key-events, jsx-a11y/no-static-element-interactions
                <a
                  className="uk-link-muted"
                  onClick={() => showCategoryModal(b.id)}
                >
                  {b.category ? b.category : 'Uncategorized'}
                </a>
              )}
              {!allowCategoryChange && (
                <span>{b.category ? b.category : 'Uncategorized'}</span>
              )}
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
              {showArchiveButton && (
                <>
                  <span className="uk-margin-small-left uk-margin-small-right">
                    ⚬
                  </span>
                  {/* eslint-disable-next-line jsx-a11y/anchor-is-valid, jsx-a11y/click-events-have-key-events, jsx-a11y/no-static-element-interactions */}
                  <a className="uk-link" onClick={() => showDeleteModal(b.id)}>
                    delete
                  </a>
                </>
              )}
              {showUnarchiveButton && (
                <>
                  <span className="uk-margin-small-left uk-margin-small-right">
                    ⚬
                  </span>
                  {/* eslint-disable-next-line jsx-a11y/anchor-is-valid, jsx-a11y/click-events-have-key-events, jsx-a11y/no-static-element-interactions */}
                  <a className="uk-link" onClick={() => unarchiveClicked(b.id)}>
                    restore
                  </a>
                </>
              )}
            </div>
          </div>
        ))}
      {bookmarks && bookmarks.length > 0 && (
        <div className="uk-flex uk-flex-row uk-flex-center uk-margin-bottom">
          <Paginator
            basePath={basePath}
            pageNumber={pageNumber}
            pageSize={pageSize}
            totalCount={totalCount}
          />
        </div>
      )}
      <Modal
        isOpen={categoryModalOpen}
        onRequestClose={hideCategoryModal}
        contentLabel="Category modal"
        style={{ overlay: { backgroundColor: 'rgba(60,60,60,0.6)' } }}
        className="uk-container-small uk-background-secondary uk-margin-auto uk-padding-small uk-margin-large-top"
      >
        <div className="uk-container-small uk-flex uk-flex-row uk-flex-between uk-modal-title">
          <h2 className="uk-modal-title">Choose category</h2>
          {/* eslint-disable-next-line max-len */}
          {/* eslint-disable-next-line jsx-a11y/anchor-has-content, jsx-a11y/anchor-is-valid, jsx-a11y/click-events-have-key-events, jsx-a11y/no-static-element-interactions */}
        </div>
        <div className="uk-flex uk-flex-row">
          <input
            type="text"
            className="uk-input"
            // eslint-disable-next-line jsx-a11y/no-autofocus
            autoFocus={categoryModalOpen}
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
      </Modal>
      <Modal
        isOpen={deleteModalOpen}
        onRequestClose={hideDeleteModal}
        contentLabel="Category modal"
        style={{ overlay: { backgroundColor: 'rgba(60,60,60,0.6)' } }}
        className="uk-container-small uk-background-secondary uk-margin-auto uk-padding uk-margin-large-top uk-width-1-1 uk-width-1-3@m uk-width-1-4@l"
      >
        <div className="uk-container-small uk-flex uk-flex-row uk-flex-between uk-modal-title">
          <h2 className="uk-modal-title">Are you sure?</h2>
        </div>
        <div className="uk-flex uk-flex-row uk-flex-between uk-margin-top">
          <button
            type="button"
            className="uk-button uk-border-rounded"
            onClick={() => hideDeleteModal()}
          >
            No
          </button>

          <button
            type="button"
            className="uk-button uk-button-danger uk-border-rounded"
            onClick={() => deleteModalSubmitted()}
          >
            Yes
          </button>
        </div>
      </Modal>
    </div>
  );
};

export default BookmarksList;
