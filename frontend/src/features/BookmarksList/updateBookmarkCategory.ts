import { createAsyncThunk } from '@reduxjs/toolkit';
import axios from 'axios';
import Bookmark from '../../models/Bookmark';
import fetchBookmarksList, { ErrorResponse } from './fetchBookmarksList';

interface UpdateBookmarkCategoryRequest {
  bookmarkID: string;
  category: string;
}

const updateBookmarkCategory = createAsyncThunk<
  Bookmark | ErrorResponse,
  UpdateBookmarkCategoryRequest
>(
  'features/fetchFunctionDetails',
  async ({ bookmarkID, category }, thunkAPI) => {
    const csrf = (document.querySelector('meta[name="csrf-token"]') as any)
      .content;

    try {
      const response = await axios.post(
        `/api/bookmarks/${bookmarkID}/update_category`,
        { category },
        { headers: { 'X-CSRF-Token': csrf } },
      );

      const rootState = thunkAPI.getState();
      thunkAPI.dispatch(
        fetchBookmarksList({
          page_size: rootState.bookmarksList.pageSize,
          page_number: rootState.bookmarksList.pageNumber,
        }),
      );

      return response.data as Bookmark;
    } catch (error) {
      console.error(error);
      return thunkAPI.rejectWithValue(error.response.data.error);
    }
  },
);

export default updateBookmarkCategory;
