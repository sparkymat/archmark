import { PayloadAction, createSlice } from '@reduxjs/toolkit';
import Bookmark from '../../models/Bookmark';
import fetchArchivedList, {
  FetchArchivedListResponse,
} from './fetchArchivedList';
import { ErrorResponse } from '../Search/searchBookmarks';

interface ArchivedListState {
  bookmarks: Bookmark[];
  errorMessage: string;
  showError?: boolean;
  loading?: boolean;
  pageNumber: number;
  pageSize: number;
  totalCount: number;
}

const initialState: ArchivedListState = {
  bookmarks: [],
  errorMessage: '',
  pageNumber: 1,
  pageSize: 10,
  totalCount: 0,
};

const slice = createSlice({
  name: 'archivedList',
  initialState,
  reducers: {
    updatePageNumber: (state, action: PayloadAction<number>) => {
      state.pageNumber = action.payload;
    },
  },
  extraReducers: builder => {
    // searchBookmarks
    builder.addCase(fetchArchivedList.pending, state => {
      state.loading = true;
    });
    builder.addCase(fetchArchivedList.fulfilled, (state, action) => {
      state.loading = false;
      state.bookmarks = (action.payload as FetchArchivedListResponse).items;
      state.totalCount = (
        action.payload as FetchArchivedListResponse
      ).total_count;
    });
    builder.addCase(fetchArchivedList.rejected, (state, action) => {
      state.loading = false;
      state.errorMessage = (action.payload as ErrorResponse).error;
      state.showError = true;
    });
  },
});

export const { updatePageNumber } = slice.actions;

export default slice.reducer;
