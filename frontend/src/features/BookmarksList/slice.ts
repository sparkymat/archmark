import { PayloadAction, createSlice } from '@reduxjs/toolkit';
import Bookmark from '../../models/Bookmark';
import fetchBookmarksList, {
  ErrorResponse,
  FetchBookmarksListResponse,
} from './fetchBookmarksList';
import fetchCategories from './fetchCategories';
import updateBookmarkCategory from './updateBookmarkCategory';

interface BookmarksListState {
  bookmarks: Bookmark[];
  categories: string[];
  errorMessage: string;
  showError?: boolean;
  loading?: boolean;
  pageNumber: number;
  pageSize: number;
  totalCount: number;
  categoryModalOpen: boolean;
  categoryModalBookmarkID?: string;
  categoryModalName: string;
  filteredCategories: string[];
}

const initialState: BookmarksListState = {
  bookmarks: [],
  categories: [],
  errorMessage: '',
  pageNumber: 1,
  pageSize: 10,
  totalCount: 0,
  categoryModalOpen: false,
  categoryModalName: '',
  filteredCategories: [],
};

const slice = createSlice({
  name: 'bookmarksList',
  initialState,
  reducers: {
    updatePageNumber: (state, action: PayloadAction<number>) => {
      state.pageNumber = action.payload;
    },
    showCategoryModal: (state, action: PayloadAction<string>) => {
      state.categoryModalBookmarkID = action.payload;
      state.categoryModalOpen = true;
    },
    hideCategoryModal: state => {
      state.categoryModalOpen = false;
      state.categoryModalBookmarkID = undefined;
    },
    updateCategoryModalName: (state, action: PayloadAction<string>) => {
      state.categoryModalName = action.payload;

      state.filteredCategories = state.categories.filter(c =>
        c.toLocaleLowerCase().includes(action.payload.toLocaleLowerCase()),
      );
    },
  },
  extraReducers: builder => {
    // fetchBookmarksList
    builder.addCase(fetchBookmarksList.pending, state => {
      state.loading = true;
    });
    builder.addCase(fetchBookmarksList.fulfilled, (state, action) => {
      state.loading = false;
      state.bookmarks = (action.payload as FetchBookmarksListResponse).items;
      state.totalCount = (
        action.payload as FetchBookmarksListResponse
      ).total_count;
    });
    builder.addCase(fetchBookmarksList.rejected, (state, action) => {
      state.loading = false;
      state.errorMessage = (action.payload as ErrorResponse).error;
      state.showError = true;
    });

    // fetchCategories
    builder.addCase(fetchCategories.pending, state => {
      state.loading = true;
    });
    builder.addCase(fetchCategories.fulfilled, (state, action) => {
      state.loading = false;
      state.categories = action.payload as string[];
      state.filteredCategories = action.payload as string[];
    });
    builder.addCase(fetchCategories.rejected, (state, action) => {
      state.loading = false;
      state.errorMessage = (action.payload as ErrorResponse).error;
      state.showError = true;
    });

    // updateBookmarkCategory
    builder.addCase(updateBookmarkCategory.pending, state => {
      state.loading = true;
    });
    builder.addCase(updateBookmarkCategory.fulfilled, state => {
      state.loading = false;
      state.categoryModalBookmarkID = undefined;
      state.categoryModalName = '';
      state.categoryModalOpen = false;
    });
    builder.addCase(updateBookmarkCategory.rejected, (state, action) => {
      state.loading = false;
      state.errorMessage = (action.payload as ErrorResponse).error;
      state.showError = true;
    });
  },
});

export const {
  showCategoryModal,
  hideCategoryModal,
  updateCategoryModalName,
  updatePageNumber,
} = slice.actions;

export default slice.reducer;
