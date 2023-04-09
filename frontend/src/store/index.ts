import { configureStore } from '@reduxjs/toolkit';
import bookmarksListReducer from '../features/BookmarksList/slice';
import newBookmarkReducer from '../features/NewBookmark/slice';
import searchReducer from '../features/Search/slice';

export const store = configureStore({
  reducer: {
    bookmarksList: bookmarksListReducer,
    newBookmark: newBookmarkReducer,
    search: searchReducer,
  },
});

// Infer the `RootState` and `AppDispatch` types from the store itself
export type RootState = ReturnType<typeof store.getState>;

// Inferred type: {posts: PostsState, comments: CommentsState, users: UsersState}
export type AppDispatch = typeof store.dispatch;
