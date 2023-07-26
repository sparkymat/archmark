import { RootState } from '../../store';

// eslint-disable-next-line import/prefer-default-export
export const selectPath = (state: RootState): string => state.app.path;
