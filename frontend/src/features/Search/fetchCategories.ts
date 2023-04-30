import { createAsyncThunk } from '@reduxjs/toolkit';
import axios from 'axios';
import { ErrorResponse } from '../BookmarksList/fetchBookmarksList';

const fetchCategories = createAsyncThunk<string[] | ErrorResponse>(
  'features/search/fetchCategories',
  async (request, thunkAPI) => {
    try {
      const response = await axios.get('/api/categories');
      return response.data;
    } catch (error) {
      console.error(error);
      return thunkAPI.rejectWithValue(error.response.data.error);
    }
  },
);

export default fetchCategories;
