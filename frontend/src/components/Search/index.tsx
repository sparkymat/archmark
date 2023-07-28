import React, { ChangeEvent, useCallback, useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { useParams } from 'react-router-dom';
import { Button, Flex, TextInput } from '@mantine/core';
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
  selectDeleteModalBookmark,
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
import Bookmark from '../../models/Bookmark';

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
  const deleteModalBookmark = useSelector(selectDeleteModalBookmark);

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
    (bookmark: Bookmark) => {
      dispatch(showDeleteModal(bookmark));
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
  }, [dispatch]);

  const submitDelete = useCallback(() => {
    if (deleteModalBookmark) {
      dispatch(deleteBookmark(deleteModalBookmark.id));
    }
  }, [deleteModalBookmark, dispatch]);

  const noop = useCallback(() => {}, []);

  return (
    <Flex direction="column">
      <Flex my="md">
        <TextInput
          sx={{ flex: 1 }}
          value={queryValue}
          onChange={queryValueChanged}
          onKeyUp={searchFormKeyUp}
          placeholder="Search here"
          // eslint-disable-next-line jsx-a11y/no-autofocus
          autoFocus
        />
        <Button variant="outline" onClick={searchSubmitted} ml="sm">
          Search
        </Button>
      </Flex>
      {query && (
        <div className="uk-padding-small">
          <p>{`Showing results for '${query}'`}</p>
        </div>
      )}
      <BookmarksList
        bookmarks={bookmarks}
        pageNumber={pageNumber}
        pageSize={pageSize}
        basePath={query ? `/#/search/${encodeURIComponent(query)}/` : '/#/'}
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
    </Flex>
  );
};

export default Search;
