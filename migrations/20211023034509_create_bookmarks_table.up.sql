CREATE TABLE public.bookmarks (
    id bigint PRIMARY KEY,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp with time zone,
    url text NOT NULL,
    title text NOT NULL,
    status text NOT NULL,
    content text,
    file_name text
);
