import React, { ChangeEvent, useCallback, useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { useParams } from 'react-router-dom';
import {
  selectBookmarks,
  selectPageSize,
  selectQueryValue,
  selectTotalCount,
  selectCategoryModalBookmarkID,
  selectCategoryModalName,
  selectCategoryModalOpen,
  selectFilteredCategories,
  selectDeleteModalOpen,
  selectDeleteModalBookmarkID,
} from '../../features/Search/selects';
import searchBookmarks from '../../features/Search/searchBookmarks';
import { AppDispatch } from '../../store';
import BookmarksList from '../BookmarksList';
import fetchCategories from '../../features/Search/fetchCategories';
import updateBookmarkCategory from '../../features/Search/updateBookmarkCategory';

import {
  updateQueryValue,
  hideCategoryModal,
  showCategoryModal,
  updateCategoryModalName,
  updateCurrentQuery,
  updatePageNumber,
  showDeleteModal,
  hideDeleteModal,
} from '../../features/Search/slice';
import deleteBookmark from '../../features/Search/deleteBookmark';

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
    dispatch(updatePageNumber(pageNumber));
    dispatch(updateCurrentQuery(query || ''));
  }, [dispatch, pageNumber, query]);

  useEffect(() => {
    dispatch(
      searchBookmarks({
        query: query || '',
        page_number: pageNumber,
        page_size: pageSize,
      }),
    );
    dispatch(updateQueryValue(query || ''));
  }, [dispatch, pageNumber, pageSize, query]);

  // TODO: Add loading spinner
  // const loading = useSelector(selectLoading);
  const bookmarks = useSelector(selectBookmarks);
  const filteredCategories = useSelector(selectFilteredCategories);
  const categoryModalOpen = useSelector(selectCategoryModalOpen);
  const deleteModalOpen = useSelector(selectDeleteModalOpen);
  const categoryModalName = useSelector(selectCategoryModalName);
  const categoryModalBookmarkID = useSelector(selectCategoryModalBookmarkID);
  const deleteModalBookmarkID = useSelector(selectDeleteModalBookmarkID);

  const queryValue = useSelector(selectQueryValue);

  const queryValueChanged = useCallback(
    (evt: ChangeEvent<HTMLInputElement>) => {
      dispatch(updateQueryValue(evt.target.value));
    },
    [dispatch],
  );

  const searchSubmitted = useCallback(() => {
    if (queryValue) {
      window.location.href = `/#/search/${encodeURIComponent(queryValue)}`;
    } else {
      window.location.href = '/#/';
    }
  }, [queryValue]);

  const searchFormKeyUp = useCallback(
    (evt: React.KeyboardEvent<HTMLInputElement>) => {
      if (evt.keyCode === 13) {
        if (queryValue) {
          window.location.href = `/#/search/${encodeURIComponent(queryValue)}`;
        } else {
          window.location.href = '/#/';
        }
      }
    },
    [queryValue],
  );

  useEffect(() => {
    dispatch(fetchCategories());
  }, [dispatch, pageNumber, pageSize, bookmarks]);

  const changeCategoryClicked = useCallback(
    (bookmarkID: string) => {
      dispatch(showCategoryModal(bookmarkID));
    },
    [dispatch],
  );

  const deleteClicked = useCallback(
    (bookmarkID: string) => {
      dispatch(showDeleteModal(bookmarkID));
    },
    [dispatch],
  );

  const categoryModalNameChanged = useCallback(
    (val: string) => {
      dispatch(updateCategoryModalName(val));
    },
    [dispatch],
  );

  const dismissCategoryModalClicked = useCallback(() => {
    dispatch(hideCategoryModal());
  }, [dispatch]);

  const submitCategoryUpdate = useCallback(() => {
    if (categoryModalBookmarkID) {
      dispatch(
        updateBookmarkCategory({
          bookmarkID: categoryModalBookmarkID,
          category: categoryModalName,
        }),
      );
    }
  }, [categoryModalBookmarkID, categoryModalName, dispatch]);

  const dismissDeleteModalClicked = useCallback(() => {
    dispatch(hideDeleteModal());
  }, []);

  const submitDelete = useCallback(() => {
    if (deleteModalBookmarkID) {
      dispatch(deleteBookmark(deleteModalBookmarkID));
    }
  }, [deleteModalBookmarkID, dispatch]);

  const noop = useCallback(() => {}, []);

  return (
    <div className="uk-container">
      <div className="uk-margin-top uk-flex uk-flex-row uk-margin-bottom">
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
        <div className="uk-padding-small">
          <p>{`Showing results for '${query}'`}</p>
        </div>
      )}
      <BookmarksList
        bookmarks={bookmarks}
        pageNumber={pageNumber}
        pageSize={pageSize}
        query={query}
        totalCount={totalCount}
        showArchiveButton
        showUnarchiveButton={false}
        categories={filteredCategories}
        categoryModalOpen={categoryModalOpen}
        categoryModalName={categoryModalName}
        hideCategoryModal={dismissCategoryModalClicked}
        showCategoryModal={changeCategoryClicked}
        categoryModalNameChanged={categoryModalNameChanged}
        categoryModalSubmitted={submitCategoryUpdate}
        deleteModalOpen={deleteModalOpen}
        showDeleteModal={deleteClicked}
        hideDeleteModal={dismissDeleteModalClicked}
        deleteModalSubmitted={submitDelete}
        allowCategoryChange
        unarchiveClicked={noop}
      />
    </div>
  );
};

export default Search;
