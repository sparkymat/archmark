import Bookmark from '../../models/Bookmark';
import { RootState } from '../../store';

export const selectLoading = (state: RootState): boolean =>
  state.bookmarksList.loading || false;

export const selectBookmarks = (state: RootState): Bookmark[] =>
  state.bookmarksList.bookmarks;

export const selectTotalCount = (state: RootState): number =>
  state.bookmarksList.totalCount;

export const selectPageSize = (state: RootState): number =>
  state.bookmarksList.pageSize;
