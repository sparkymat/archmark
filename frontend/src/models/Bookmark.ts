interface Bookmark {
  id: string;
  title?: string;
  url: string;
  file_path?: string;
  status: 'pending' | 'fetched' | 'archived';
  created_at: string;
  updated_at: string;
}

export default Bookmark;
