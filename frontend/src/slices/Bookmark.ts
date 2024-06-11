import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react';
import { Bookmark } from '../models/Bookmark';

export interface ListResponse {
  items: Bookmark[];
  totalCount: number;
  pageNumber: number;
  pageSize: number;
}

export interface FetchRecentRequest {
  pageSize: number;
  pageNumber: number;
}

export interface SearchRequest {
  query: string;
  pageSize: number;
  pageNumber: number;
}

export interface CreateRequest {
  title: string;
  url: string;
  description: string;
}

export interface UpdateTitleRequest {
  id: string;
  title: string;
}

export interface UpdateUrlRequest {
  id: string;
  url: string;
}

export interface UpdateDescriptionRequest {
  id: string;
  description: string;
}

export interface UpdateTranscriptRequest {
  id: string;
  transcript: string;
}

export interface UploadImageRequest {
  id: string;
  formData: FormData;
}

export const api = createApi({
  reducerPath: 'bookmarks',
  baseQuery: fetchBaseQuery({ baseUrl: '/api' }),
  tagTypes: ['Bookmark'],
  endpoints: builder => ({
    recent: builder.query<ListResponse, FetchRecentRequest>({
      query: ({ pageSize, pageNumber }) =>
        `bookmarks/recent?pageSize=${pageSize}&pageNumber=${pageNumber}`,
      providesTags: [{ type: 'Bookmark', id: 'LIST' }],
    }),
    search: builder.query<ListResponse, SearchRequest>({
      query: ({ query, pageSize, pageNumber }) =>
        `bookmarks/search?query=${encodeURIComponent(query)}&pageSize=${pageSize}&pageNumber=${pageNumber}`,
      providesTags: [{ type: 'Bookmark', id: 'SEARCH' }],
    }),
    show: builder.query<Bookmark, string>({
      query: id => `bookmarks/${id}`,
      providesTags: (_result, _error, arg) => [{ type: 'Bookmark', id: arg }],
    }),
    create: builder.mutation<Bookmark, CreateRequest>({
      query: body => ({
        url: 'bookmarks',
        method: 'POST',
        body,
        headers: {
          'X-CSRF-Token': (
            document.querySelector('meta[name="csrf-token"]') as any
          ).content,
        },
      }),
      invalidatesTags: [
        { type: 'Bookmark', id: 'LIST' },
        { type: 'Bookmark', id: 'SEARCH' },
      ],
    }),
    destroy: builder.mutation<void, string>({
      query: id => ({
        url: `bookmarks/${id}`,
        method: 'DELETE',
        headers: {
          'X-CSRF-Token': (
            document.querySelector('meta[name="csrf-token"]') as any
          ).content,
        },
      }),
      invalidatesTags: (_result, _error, arg) => [
        { type: 'Bookmark', id: arg },
        { type: 'Bookmark', id: 'LIST' },
        { type: 'Bookmark', id: 'SEARCH' },
      ],
    }),
    updateTitle: builder.mutation<Bookmark, UpdateTitleRequest>({
      query: ({ id, title }) => ({
        url: `bookmarks/${id}/update_title`,
        method: 'PATCH',
        body: { title },
        headers: {
          'X-CSRF-Token': (
            document.querySelector('meta[name="csrf-token"]') as any
          ).content,
        },
      }),
      invalidatesTags: (_result, _error, arg) => [
        { type: 'Bookmark', id: 'LIST' },
        { type: 'Bookmark', id: 'SEARCH' },
        { type: 'Bookmark', id: arg.id },
      ],
    }),
    updateUrl: builder.mutation<Bookmark, UpdateUrlRequest>({
      query: ({ id, url }) => ({
        url: `bookmarks/${id}/update_url`,
        method: 'PATCH',
        body: { url },
        headers: {
          'X-CSRF-Token': (
            document.querySelector('meta[name="csrf-token"]') as any
          ).content,
        },
      }),
      invalidatesTags: (_result, _error, arg) => [
        { type: 'Bookmark', id: 'LIST' },
        { type: 'Bookmark', id: 'SEARCH' },
        { type: 'Bookmark', id: arg.id },
      ],
    }),
    updateDescription: builder.mutation<Bookmark, UpdateDescriptionRequest>({
      query: ({ id, description }) => ({
        url: `bookmarks/${id}/update_description`,
        method: 'PATCH',
        body: { description },
        headers: {
          'X-CSRF-Token': (
            document.querySelector('meta[name="csrf-token"]') as any
          ).content,
        },
      }),
      invalidatesTags: (_result, _error, arg) => [
        { type: 'Bookmark', id: 'LIST' },
        { type: 'Bookmark', id: 'SEARCH' },
        { type: 'Bookmark', id: arg.id },
      ],
    }),
    updateTranscript: builder.mutation<Bookmark, UpdateTranscriptRequest>({
      query: ({ id, transcript }) => ({
        url: `bookmarks/${id}/update_transcript`,
        method: 'PATCH',
        body: { transcript },
        headers: {
          'X-CSRF-Token': (
            document.querySelector('meta[name="csrf-token"]') as any
          ).content,
        },
      }),
      invalidatesTags: (_result, _error, arg) => [
        { type: 'Bookmark', id: 'LIST' },
        { type: 'Bookmark', id: 'SEARCH' },
        { type: 'Bookmark', id: arg.id },
      ],
    }),
    uploadImage: builder.mutation<Bookmark, UploadImageRequest>({
      query: ({ id, formData }) => ({
        url: `bookmarks/${id}/upload_image`,
        method: 'PATCH',
        body: formData,
        headers: {
          'X-CSRF-Token': (
            document.querySelector('meta[name="csrf-token"]') as any
          ).content,
        },
      }),
      invalidatesTags: (_result, _error, arg) => [
        { type: 'Bookmark', id: 'LIST' },
        { type: 'Bookmark', id: 'SEARCH' },
        { type: 'Bookmark', id: arg.id },
      ],
    }),
  }),
});

export const {
  useRecentQuery,
  useSearchQuery,
  useCreateMutation,
  useShowQuery,
  useUpdateTitleMutation,
  useUpdateUrlMutation,
  useUpdateDescriptionMutation,
  useUpdateTranscriptMutation,
  useUploadImageMutation,
  useDestroyMutation,
} = api;
