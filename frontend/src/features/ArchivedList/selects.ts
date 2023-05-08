import Bookmark from '../../models/Bookmark';
import { RootState } from '../../store';

export const selectLoading = (state: RootState): boolean =>
  state.archivedList.loading || false;

export const selectBookmarks = (state: RootState): Bookmark[] =>
  state.archivedList.bookmarks;

export const selectTotalCount = (state: RootState): number =>
  state.archivedList.totalCount;

export const selectPageSize = (state: RootState): number =>
  state.archivedList.pageSize;
