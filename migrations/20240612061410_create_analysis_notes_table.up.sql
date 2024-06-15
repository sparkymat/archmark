CREATE EXTENSION IF NOT EXISTS moddatetime;

CREATE TYPE analysis_note_source AS ENUM (
'user',
'service'
);

CREATE TYPE analysis_note_state AS ENUM (
'ready',
'pending',
'failed'
);

CREATE TABLE analysis_notes (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  content text NOT NULL,
  source analysis_note_source NOT NULL,
  state analysis_note_state NOT NULL DEFAULT 'ready',
  bookmark_id uuid REFERENCES bookmarks(id) NOT NULL,
  created_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE TRIGGER analysis_notes_updated_at
  BEFORE UPDATE
  ON analysis_notes
  FOR EACH ROW
    EXECUTE FUNCTION moddatetime(updated_at);
