import { createAsyncThunk } from '@reduxjs/toolkit';
import axios from 'axios';
import Bookmark from '../../models/Bookmark';
import searchBookmarks from './searchBookmarks';
import { ErrorResponse } from '../BookmarksList/fetchBookmarksList';

interface UpdateBookmarkCategoryRequest {
  bookmarkID: string;
  category: string;
}

const updateBookmarkCategory = createAsyncThunk<
  Bookmark | ErrorResponse,
  UpdateBookmarkCategoryRequest
>(
  'features/search/updateBookmarkCategory',
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
        searchBookmarks({
          page_size: (rootState as any).search.pageSize,
          page_number: (rootState as any).search.pageNumber,
          query: (rootState as any).search.currentQuery,
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
