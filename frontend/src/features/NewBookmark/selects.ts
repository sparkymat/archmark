import { RootState } from '../../store';

export const selectURL = (state: RootState): string => state.newBookmark.url;

export const selectLoading = (state: RootState): boolean =>
  state.newBookmark.loading || false;
