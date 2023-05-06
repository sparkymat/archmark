import { createAsyncThunk } from '@reduxjs/toolkit';
import axios from 'axios';
import searchBookmarks, { ErrorResponse } from './searchBookmarks';

const deleteFunction = createAsyncThunk<string | ErrorResponse, string>(
  'features/deleteBookmark',
  async (id, thunkAPI) => {
    const csrf = (document.querySelector('meta[name="csrf-token"]') as any)
      .content;

    try {
      await axios.delete(`/api/bookmarks/${id}`, {
        headers: { 'X-CSRF-Token': csrf },
      });

      const rootState = thunkAPI.getState();
      thunkAPI.dispatch(
        searchBookmarks({
          page_size: (rootState as any).search.pageSize,
          page_number: (rootState as any).search.pageNumber,
          query: (rootState as any).search.currentQuery,
        }),
      );

      return id;
    } catch (error) {
      console.error(error);
      return thunkAPI.rejectWithValue(error.response.data.error);
    }
  },
);

export default deleteFunction;
