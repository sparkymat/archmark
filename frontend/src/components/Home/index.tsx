import React, { useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { useParams } from 'react-router-dom';
import {
  selectBookmarks,
  selectPageSize,
  selectTotalCount,
} from '../../features/BookmarksList/selects';
import { AppDispatch } from '../../store';
import fetchBookmarksList from '../../features/BookmarksList/fetchBookmarksList';
import BookmarkList from '../BookmarksList';

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
    <BookmarkList
      bookmarks={bookmarks}
      pageNumber={pageNumber}
      pageSize={pageSize}
      totalCount={totalCount}
    />
  );
};

export default Home;
