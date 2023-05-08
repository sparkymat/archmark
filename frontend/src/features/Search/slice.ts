import { PayloadAction, createSlice } from '@reduxjs/toolkit';
import Bookmark from '../../models/Bookmark';
import searchBookmarks, {
  SearchResponse,
  ErrorResponse,
} from './searchBookmarks';
import fetchCategories from './fetchCategories';
import updateBookmarkCategory from './updateBookmarkCategory';
import deleteBookmark from './deleteBookmark';

interface SearchState {
  bookmarks: Bookmark[];
  errorMessage: string;
  showError?: boolean;
  loading?: boolean;
  pageNumber: number;
  pageSize: number;
  totalCount: number;
  currentQuery: string;
  queryValue: string;
  categories: string[];
  categoryModalOpen: boolean;
  categoryModalBookmarkID?: string;
  categoryModalName: string;
  filteredCategories: string[];
  deleteModalOpen: boolean;
  deleteModalBookmarkID?: string;
}

const initialState: SearchState = {
  bookmarks: [],
  errorMessage: '',
  pageNumber: 1,
  pageSize: 10,
  totalCount: 0,
  currentQuery: '',
  queryValue: '',
  categories: [],
  categoryModalOpen: false,
  categoryModalName: '',
  filteredCategories: [],
  deleteModalOpen: false,
};

const slice = createSlice({
  name: 'search',
  initialState,
  reducers: {
    updatePageNumber: (state, action: PayloadAction<number>) => {
      state.pageNumber = action.payload;
    },
    updateCurrentQuery: (state, action: PayloadAction<string>) => {
      state.currentQuery = action.payload;
    },
    updateQueryValue: (state, action: PayloadAction<string>) => {
      state.queryValue = action.payload;
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

      state.filteredCategories = state.categories
        .sort()
        .filter(c =>
          c.toLocaleLowerCase().includes(action.payload.toLocaleLowerCase()),
        );
    },
    showDeleteModal: (state, action: PayloadAction<string>) => {
      state.deleteModalBookmarkID = action.payload;
      state.deleteModalOpen = true;
    },
    hideDeleteModal: state => {
      state.deleteModalBookmarkID = undefined;
      state.deleteModalOpen = false;
    },
  },
  extraReducers: builder => {
    // searchBookmarks
    builder.addCase(searchBookmarks.pending, state => {
      state.loading = true;
    });
    builder.addCase(searchBookmarks.fulfilled, (state, action) => {
      state.loading = false;
      state.bookmarks = (action.payload as SearchResponse).items;
      state.totalCount = (action.payload as SearchResponse).total_count;
    });
    builder.addCase(searchBookmarks.rejected, (state, action) => {
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

    // deleteBookmark
    builder.addCase(deleteBookmark.pending, state => {
      state.loading = true;
    });
    builder.addCase(deleteBookmark.fulfilled, state => {
      state.loading = false;
      state.deleteModalBookmarkID = undefined;
      state.deleteModalOpen = false;
    });
    builder.addCase(deleteBookmark.rejected, (state, action) => {
      state.loading = false;
      state.errorMessage = (action.payload as ErrorResponse).error;
      state.showError = true;
    });
  },
});

export const {
  updateQueryValue,
  showCategoryModal,
  hideCategoryModal,
  updateCategoryModalName,
  updateCurrentQuery,
  updatePageNumber,
  showDeleteModal,
  hideDeleteModal,
} = slice.actions;

export default slice.reducer;
