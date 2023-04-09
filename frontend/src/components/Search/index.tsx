import React, { useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { useParams } from 'react-router-dom';
import {
  selectBookmarks,
  selectPageSize,
  selectTotalCount,
} from '../../features/Search/selects';
import searchBookmarks from '../../features/Search/searchBookmarks';
import { AppDispatch } from '../../store';
import BookmarksList from '../BookmarksList';

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
    }
  }, [dispatch, pageNumber, pageSize, query]);

  // TODO: Add loading spinner
  // const loading = useSelector(selectLoading);
  const bookmarks = useSelector(selectBookmarks);

  return (
    <BookmarksList
      bookmarks={bookmarks}
      pageNumber={pageNumber}
      pageSize={pageSize}
      totalCount={totalCount}
    />
  );
};

export default Search;
