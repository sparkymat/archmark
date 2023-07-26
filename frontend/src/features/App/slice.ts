import { PayloadAction, createSlice } from '@reduxjs/toolkit';

export interface AppState {
  path: string;
}

const initialState: AppState = {
  path: '/',
};

const slice = createSlice({
  name: 'app',
  initialState,
  reducers: {
    updatePath: (state, action: PayloadAction<string>) => {
      state.path = action.payload;
    },
  },
});

export const { updatePath } = slice.actions;

export default slice.reducer;
