import { createAsyncThunk } from '@reduxjs/toolkit';
import axios from 'axios';
import Bookmark from '../../models/Bookmark';
import { ErrorResponse } from '../Search/searchBookmarks';

export interface FetchArchivedListRequest {
  page_size: number;
  page_number: number;
}

export interface FetchArchivedListResponse {
  page_size: number;
  page_number: number;
  items: Bookmark[];
  total_count: number;
}

const fetchArchivedList = createAsyncThunk<
  FetchArchivedListResponse | ErrorResponse,
  FetchArchivedListRequest
>('features/archivedList/fetchArchivedList', async (request, thunkAPI) => {
  try {
    const response = await axios.get(
      `/api/bookmarks/archived?page_size=${request.page_size}&page_number=${request.page_number}`,
    );
    return response.data;
  } catch (error) {
    console.error(error);
    return thunkAPI.rejectWithValue(error.response.data.error);
  }
});

export default fetchArchivedList;
