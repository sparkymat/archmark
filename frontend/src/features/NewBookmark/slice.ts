import { PayloadAction, createSlice } from '@reduxjs/toolkit';

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
});

export const { updateURL } = slice.actions;

export default slice.reducer;
