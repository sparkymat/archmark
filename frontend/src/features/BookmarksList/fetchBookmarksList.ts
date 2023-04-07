import { createAsyncThunk } from '@reduxjs/toolkit';
import axios from 'axios';
import Bookmark from '../../models/Bookmark';

export interface FetchBookmarksListRequest {
  page_size: number;
  page_number: number;
}

export interface FetchBookmarksListResponse {
  page_size: number;
  page_number: number;
  items: Bookmark[];
  total_count: number;
}

export interface ErrorResponse {
  error: string;
}

const fetchBookmarksList = createAsyncThunk<
  FetchBookmarksListResponse | ErrorResponse,
  FetchBookmarksListRequest
>('features/bookmarksList/fetchBookmarksList', async (request, thunkAPI) => {
  try {
    const response = await axios.get(
      `/api/bookmarks?page_size=${request.page_size}&page_number=${request.page_number}`,
    );
    return response.data;
  } catch (error) {
    console.error(error);
    return thunkAPI.rejectWithValue(error.response.data.error);
  }
});

export default fetchBookmarksList;
