import React, { ChangeEvent, useCallback, useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { useParams } from 'react-router-dom';
import {
  selectBookmarks,
  selectPageSize,
  selectQueryValue,
  selectTotalCount,
} from '../../features/Search/selects';
import searchBookmarks from '../../features/Search/searchBookmarks';
import { AppDispatch } from '../../store';
import BookmarksList from '../BookmarksList';
import { updateQueryValue } from '../../features/Search/slice';

const Search = () => {
  const dispatch = useDispatch<AppDispatch>();
  const { page: pageNumberString, query } = useParams();

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
    if (query) {
      dispatch(
        searchBookmarks({
          query,
          page_number: pageNumber,
          page_size: pageSize,
        }),
      );
      dispatch(updateQueryValue(query));
    }
  }, [dispatch, pageNumber, pageSize, query]);

  // TODO: Add loading spinner
  // const loading = useSelector(selectLoading);
  const bookmarks = useSelector(selectBookmarks);

  const queryValue = useSelector(selectQueryValue);

  const queryValueChanged = useCallback(
    (evt: ChangeEvent<HTMLInputElement>) => {
      dispatch(updateQueryValue(evt.target.value));
    },
    [dispatch],
  );

  const searchSubmitted = useCallback(() => {
    window.location.href = `/#/search/${encodeURIComponent(queryValue)}`;
  }, [queryValue]);

  const searchFormKeyUp = useCallback(
    (evt: React.KeyboardEvent<HTMLInputElement>) => {
      if (evt.keyCode === 13) {
        window.location.href = `/#/search/${encodeURIComponent(queryValue)}`;
      }
    },
    [queryValue],
  );

  return (
    <div className="uk-container">
      <div className="uk-margin-top uk-flex uk-flex-row">
        <input
          className="uk-input uk-form uk-border-rounded"
          type="text"
          value={queryValue}
          onChange={queryValueChanged}
          onKeyUp={searchFormKeyUp}
          placeholder="Search here"
          // eslint-disable-next-line jsx-a11y/no-autofocus
          autoFocus
        />
        <button
          type="button"
          onClick={searchSubmitted}
          className="uk-button uk-margin-small-left uk-border-rounded"
        >
          Search
        </button>
      </div>
      {query && (
        <>
          <div className="uk-padding-small uk-margin-small-top">
            <p>{`Showing results for '${query}'`}</p>
          </div>
          <BookmarksList
            bookmarks={bookmarks}
            pageNumber={pageNumber}
            pageSize={pageSize}
            totalCount={totalCount}
          />
        </>
      )}
    </div>
  );
};

export default Search;
