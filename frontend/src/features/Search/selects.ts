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

export const selectCategories = (state: RootState): string[] =>
  state.search.categories;

export const selectCategoryModalOpen = (state: RootState): boolean =>
  state.search.categoryModalOpen;

export const selectCategoryModalName = (state: RootState): string =>
  state.search.categoryModalName;

export const selectFilteredCategories = (state: RootState): string[] =>
  state.search.filteredCategories;

export const selectCategoryModalBookmarkID = (
  state: RootState,
): string | undefined => state.search.categoryModalBookmarkID;
