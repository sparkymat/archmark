import React, { useCallback, useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { useParams } from 'react-router-dom';
import {
  selectBookmarks,
  selectPageSize,
  selectTotalCount,
} from '../../features/ArchivedList/selects';
import { AppDispatch } from '../../store';
import BookmarksList from '../BookmarksList';
import { updatePageNumber } from '../../features/ArchivedList/slice';
import fetchArchivedList from '../../features/ArchivedList/fetchArchivedList';

const ArchivedList = () => {
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
    dispatch(updatePageNumber(pageNumber));
  }, [dispatch, pageNumber]);

  useEffect(() => {
    dispatch(
      fetchArchivedList({
        page_number: pageNumber,
        page_size: pageSize,
      }),
    );
  }, [dispatch, pageNumber, pageSize]);

  // TODO: Add loading spinner
  // const loading = useSelector(selectLoading);
  const bookmarks = useSelector(selectBookmarks);

  const noop = useCallback(() => {}, []);

  return (
    <div className="uk-container uk-margin-top">
      <BookmarksList
        bookmarks={bookmarks}
        pageNumber={pageNumber}
        pageSize={pageSize}
        totalCount={totalCount}
        showArchiveButton={false}
        allowCategoryChange={false}
        categories={[]}
        categoryModalOpen={false}
        categoryModalName=""
        hideCategoryModal={noop}
        showCategoryModal={noop}
        categoryModalNameChanged={noop}
        categoryModalSubmitted={noop}
        deleteModalOpen={false}
        showDeleteModal={noop}
        hideDeleteModal={noop}
        deleteModalSubmitted={noop}
      />
    </div>
  );
};

export default ArchivedList;
