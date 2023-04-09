import Bookmark from '../../models/Bookmark';
import { RootState } from '../../store';

export const selectQueryValue = (state: RootState): string =>
  state.search.queryValue;

export const selectLoading = (state: RootState): boolean =>
  state.search.loading || false;

export const selectBookmarks = (state: RootState): Bookmark[] =>
  state.search.bookmarks;

export const selectTotalCount = (state: RootState): number =>
  state.search.totalCount;

export const selectPageSize = (state: RootState): number =>
  state.search.pageSize;
