import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react';
import { AnalysisNote } from '../models/AnalysisNote';

export interface ListResponse {
  items: AnalysisNote[];
  totalCount: number;
  pageNumber: number;
  pageSize: number;
}

export interface CreateRequest {
  content: string;
  bookmarkID: string;
}

export interface UpdateContentRequest {
  id: string;
  content: string;
}

export const api = createApi({
  reducerPath: 'analysisNotes',
  baseQuery: fetchBaseQuery({ baseUrl: '/api' }),
  tagTypes: ['AnalysisNote'],
  endpoints: builder => ({
    recent: builder.query<ListResponse, string>({
      query: parentID => `/bookmarks/${parentID}/analysis_notes/recent`,
      providesTags: [{ type: 'AnalysisNote', id: 'LIST' }],
    }),
    show: builder.query<AnalysisNote, string>({
      query: id => `analysis_notes/${id}`,
      providesTags: (_result, _error, arg) => [
        { type: 'AnalysisNote', id: arg },
      ],
    }),
    create: builder.mutation<AnalysisNote, CreateRequest>({
      query: body => ({
        url: 'analysis_notes',
        method: 'POST',
        body,
        headers: {
          'X-CSRF-Token': (
            document.querySelector('meta[name="csrf-token"]') as any
          ).content,
        },
      }),
      invalidatesTags: [
        { type: 'AnalysisNote', id: 'LIST' },
        { type: 'AnalysisNote', id: 'SEARCH' },
      ],
    }),
    destroy: builder.mutation<void, string>({
      query: id => ({
        url: `analysis_notes/${id}`,
        method: 'DELETE',
        headers: {
          'X-CSRF-Token': (
            document.querySelector('meta[name="csrf-token"]') as any
          ).content,
        },
      }),
      invalidatesTags: (_result, _error, arg) => [
        { type: 'AnalysisNote', id: arg },
        { type: 'AnalysisNote', id: 'LIST' },
        { type: 'AnalysisNote', id: 'SEARCH' },
      ],
    }),
    updateContent: builder.mutation<AnalysisNote, UpdateContentRequest>({
      query: ({ id, content }) => ({
        url: `analysis_notes/${id}/update_content`,
        method: 'PATCH',
        body: { content },
        headers: {
          'X-CSRF-Token': (
            document.querySelector('meta[name="csrf-token"]') as any
          ).content,
        },
      }),
      invalidatesTags: (_result, _error, arg) => [
        { type: 'AnalysisNote', id: 'LIST' },
        { type: 'AnalysisNote', id: 'SEARCH' },
        { type: 'AnalysisNote', id: arg.id },
      ],
    }),
  }),
});

export const {
  useRecentQuery,
  useCreateMutation,
  useShowQuery,
  useUpdateContentMutation,
  useDestroyMutation,
} = api;
