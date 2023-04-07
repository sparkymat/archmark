import { createAsyncThunk } from '@reduxjs/toolkit';
import axios from 'axios';
import Bookmark from '../../models/Bookmark';
import { ErrorResponse } from '../BookmarksList/fetchBookmarksList';

const createFunction = createAsyncThunk<Bookmark | ErrorResponse, string>(
  'features/fetchFunctionDetails',
  async (url, thunkAPI) => {
    const csrf = (document.querySelector('meta[name="csrf-token"]') as any)
      .content;

    try {
      const response = await axios.post(
        '/api/bookmarks',
        { url },
        { headers: { 'X-CSRF-Token': csrf } },
      );

      window.location.href = '/';

      return response.data as Bookmark;
    } catch (error) {
      console.error(error);
      return thunkAPI.rejectWithValue(error.response.data.error);
    }
  },
);

export default createFunction;
