import { RootState } from '.';
import { State as App } from '../slices/App';

export const selectApp = (state: RootState): App => state.app;
