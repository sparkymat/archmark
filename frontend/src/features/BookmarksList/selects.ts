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

export const selectCategories = (state: RootState): string[] =>
  state.bookmarksList.categories;

export const selectCategoryModalOpen = (state: RootState): boolean =>
  state.bookmarksList.categoryModalOpen;

export const selectCategoryModalName = (state: RootState): string =>
  state.bookmarksList.categoryModalName;

export const selectFilteredCategories = (state: RootState): string[] =>
  state.bookmarksList.filteredCategories;

export const selectCategoryModalBookmarkID = (
  state: RootState,
): string | undefined => state.bookmarksList.categoryModalBookmarkID;
