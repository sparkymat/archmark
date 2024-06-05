import { PayloadAction, createSlice } from "@reduxjs/toolkit";

export interface State {
  loading: boolean;
  errorMessage: string;
  showError: boolean;
  path: string;
}

const initialState: State = {
  loading: false,
  errorMessage: "",
  showError: false,
  path: "/",
};

const slice = createSlice({
  name: "app",
  initialState,
  reducers: {
    updatePath: (state, action: PayloadAction<string>) => {
      state.path = action.payload;
    },
  },
});

export const { updatePath } = slice.actions;

export const { reducer } = slice;
