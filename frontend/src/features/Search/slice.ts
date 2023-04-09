import { PayloadAction, createSlice } from '@reduxjs/toolkit';
import Bookmark from '../../models/Bookmark';
import searchBookmarks, {
  ErrorResponse,
  SearchResponse,
} from './searchBookmarks';

interface SearchState {
  bookmarks: Bookmark[];
  errorMessage: string;
  showError?: boolean;
  loading?: boolean;
  pageSize: number;
  totalCount: number;
  queryValue: string;
}

const initialState: SearchState = {
  bookmarks: [],
  errorMessage: '',
  pageSize: 10,
  totalCount: 0,
  queryValue: '',
};

const slice = createSlice({
  name: 'search',
  initialState,
  reducers: {
    updateQueryValue: (state, action: PayloadAction<string>) => {
      state.queryValue = action.payload;
    },
  },
  extraReducers: builder => {
    // searchBookmarks
    builder.addCase(searchBookmarks.pending, state => {
      state.loading = true;
    });
    builder.addCase(searchBookmarks.fulfilled, (state, action) => {
      state.loading = false;
      state.bookmarks = (action.payload as SearchResponse).items;
      state.totalCount = (action.payload as SearchResponse).total_count;
    });
    builder.addCase(searchBookmarks.rejected, (state, action) => {
      state.loading = false;
      state.errorMessage = (action.payload as ErrorResponse).error;
      state.showError = true;
    });
  },
});

export const { updateQueryValue } = slice.actions;

export default slice.reducer;
