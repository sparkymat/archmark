import { createAsyncThunk } from '@reduxjs/toolkit';
import axios from 'axios';
import Bookmark from '../../models/Bookmark';
import { ErrorResponse } from '../BookmarksList/fetchBookmarksList';

export interface SearchRequest {
  page_size: number;
  page_number: number;
  query: string;
}

export interface SearchResponse {
  page_size: number;
  page_number: number;
  items: Bookmark[];
  total_count: number;
}

const searchBookmarks = createAsyncThunk<
  SearchResponse | ErrorResponse,
  SearchRequest
>('features/search/searchBookmarks', async (request, thunkAPI) => {
  try {
    const response = await axios.get(
      `/api/bookmarks/search?query=${encodeURIComponent(
        request.query,
      )}&page_size=${request.page_size}&page_number=${request.page_number}`,
    );
    return response.data;
  } catch (error) {
    console.error(error);
    return thunkAPI.rejectWithValue(error.response.data.error);
  }
});

export default searchBookmarks;
