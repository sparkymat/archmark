import { PayloadAction, createSlice } from '@reduxjs/toolkit';
import createBookmark from './createBookmark';
import { ErrorResponse } from '../BookmarksList/fetchBookmarksList';

interface NewBookmarkState {
  url: string;
  errorMessage: string;
  showError?: boolean;
  loading?: boolean;
}

const initialState: NewBookmarkState = {
  url: '',
  errorMessage: '',
};

const slice = createSlice({
  name: 'newBookmark',
  initialState,
  reducers: {
    updateURL: (state, action: PayloadAction<string>) => {
      state.url = action.payload;
    },
  },
  extraReducers: builder => {
    // createBookmark
    builder.addCase(createBookmark.pending, state => {
      state.loading = true;
    });
    builder.addCase(createBookmark.fulfilled, state => {
      state.loading = false;
    });
    builder.addCase(createBookmark.rejected, (state, action) => {
      state.loading = false;
      state.errorMessage = (action.payload as ErrorResponse).error;
    });
  },
});

export const { updateURL } = slice.actions;

export default slice.reducer;
