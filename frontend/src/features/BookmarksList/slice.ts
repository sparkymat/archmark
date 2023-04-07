import { createSlice } from '@reduxjs/toolkit';
import Bookmark from '../../models/Bookmark';
import fetchBookmarksList, {
  ErrorResponse,
  FetchBookmarksListResponse,
} from './fetchBookmarksList';

interface BookmarksListState {
  bookmarks: Bookmark[];
  errorMessage: string;
  showError?: boolean;
  loading?: boolean;
  pageSize: number;
  totalCount: number;
}

const initialState: BookmarksListState = {
  bookmarks: [],
  errorMessage: '',
  pageSize: 10,
  totalCount: 0,
};

const slice = createSlice({
  name: 'bookmarksList',
  initialState,
  reducers: {},
  extraReducers: builder => {
    // fetchBookmarksList
    builder.addCase(fetchBookmarksList.pending, state => {
      state.loading = true;
    });
    builder.addCase(fetchBookmarksList.fulfilled, (state, action) => {
      state.loading = false;
      state.bookmarks = (action.payload as FetchBookmarksListResponse).items;
      state.totalCount = (
        action.payload as FetchBookmarksListResponse
      ).total_count;
    });
    builder.addCase(fetchBookmarksList.rejected, (state, action) => {
      state.loading = false;
      state.errorMessage = (action.payload as ErrorResponse).error;
      state.showError = true;
    });
  },
});

export default slice.reducer;
