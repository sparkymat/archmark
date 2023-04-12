import React, { useCallback, useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { useParams } from 'react-router-dom';
import {
  selectBookmarks,
  selectCategoryModalBookmarkID,
  selectCategoryModalName,
  selectCategoryModalOpen,
  selectFilteredCategories,
  selectPageSize,
  selectTotalCount,
} from '../../features/BookmarksList/selects';
import { AppDispatch } from '../../store';
import fetchBookmarksList from '../../features/BookmarksList/fetchBookmarksList';
import BookmarkList from '../BookmarksList';
import fetchCategories from '../../features/BookmarksList/fetchCategories';
import {
  hideCategoryModal,
  showCategoryModal,
  updateCategoryModalName,
  updatePageNumber,
} from '../../features/BookmarksList/slice';
import updateBookmarkCategory from '../../features/BookmarksList/updateBookmarkCategory';

const Home = () => {
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

  useEffect(() => {
    dispatch(updatePageNumber(pageNumber));
  }, [dispatch, pageNumber]);

  const pageSize = useSelector(selectPageSize);
  const totalCount = useSelector(selectTotalCount);

  // TODO: Add loading spinner
  // const loading = useSelector(selectLoading);
  const bookmarks = useSelector(selectBookmarks);
  const filteredCategories = useSelector(selectFilteredCategories);
  const categoryModalOpen = useSelector(selectCategoryModalOpen);
  const categoryModalName = useSelector(selectCategoryModalName);
  const categoryModalBookmarkID = useSelector(selectCategoryModalBookmarkID);

  useEffect(() => {
    dispatch(
      fetchBookmarksList({
        page_number: pageNumber,
        page_size: pageSize,
      }),
    );
  }, [dispatch, pageNumber, pageSize]);

  useEffect(() => {
    dispatch(fetchCategories());
  }, [dispatch, pageNumber, pageSize, bookmarks]);

  const changeCategoryClicked = useCallback(
    (bookmarkID: string) => {
      dispatch(showCategoryModal(bookmarkID));
    },
    [dispatch],
  );

  const categoryModalNameChanged = useCallback((val: string) => {
    dispatch(updateCategoryModalName(val));
  }, []);

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

  return (
    <BookmarkList
      bookmarks={bookmarks}
      categories={filteredCategories}
      pageNumber={pageNumber}
      pageSize={pageSize}
      totalCount={totalCount}
      categoryModalOpen={categoryModalOpen}
      hideCategoryModal={dismissCategoryModalClicked}
      showCategoryModal={changeCategoryClicked}
      categoryModalName={categoryModalName}
      categoryModalNameChanged={categoryModalNameChanged}
      categoryModalSubmitted={submitCategoryUpdate}
    />
  );
};

export default Home;
