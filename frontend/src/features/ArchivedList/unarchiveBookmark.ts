import { createAsyncThunk } from '@reduxjs/toolkit';
import axios from 'axios';
import { ErrorResponse } from '../Search/searchBookmarks';
import fetchArchivedList from './fetchArchivedList';

const unarchiveBookmark = createAsyncThunk<string | ErrorResponse, string>(
  'features/archivedList/unarchiveBookmark',
  async (id, thunkAPI) => {
    const csrf = (document.querySelector('meta[name="csrf-token"]') as any)
      .content;

    try {
      await axios.post(
        `/api/bookmarks/${id}/unarchive`,
        {},
        {
          headers: { 'X-CSRF-Token': csrf },
        },
      );

      const rootState = thunkAPI.getState();
      thunkAPI.dispatch(
        fetchArchivedList({
          page_size: (rootState as any).archivedList.pageSize,
          page_number: (rootState as any).archivedList.pageNumber,
        }),
      );

      return id;
    } catch (error) {
      console.error(error);
      return thunkAPI.rejectWithValue(error.response.data.error);
    }
  },
);

export default unarchiveBookmark;
