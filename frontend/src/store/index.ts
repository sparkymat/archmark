import { configureStore } from '@reduxjs/toolkit';
import newBookmarkReducer from '../features/NewBookmark/slice';
import searchReducer from '../features/Search/slice';
import archivedListReducer from '../features/ArchivedList/slice';
import appReducer from '../features/App/slice';

export const store = configureStore({
  reducer: {
    newBookmark: newBookmarkReducer,
    search: searchReducer,
    archivedList: archivedListReducer,
    app: appReducer,
  },
});

// Infer the `RootState` and `AppDispatch` types from the store itself
export type RootState = ReturnType<typeof store.getState>;

// Inferred type: {posts: PostsState, comments: CommentsState, users: UsersState}
export type AppDispatch = typeof store.dispatch;
