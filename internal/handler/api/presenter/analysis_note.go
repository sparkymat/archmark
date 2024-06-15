package presenter

import (
	"time"

	"github.com/sparkymat/archmark/internal/dbx"
)

type AnalysisNote struct {
	ID         string `json:"id"`
	Content    string `json:"content"`
	Source     string `json:"source"`
	BookmarkId string `json:"bookmarkId"`
	CreatedAt  string `json:"createdAt"`
	UpdatedAt  string `json:"updatedAt"`
}

func AnalysisNoteFromModel(m dbx.AnalysisNote) AnalysisNote {
	item := AnalysisNote{
		ID:        m.ID.String(),
		CreatedAt: m.CreatedAt.Time.Format(time.RFC3339),
		UpdatedAt: m.UpdatedAt.Time.Format(time.RFC3339),
	}

	item.Content = m.Content

	item.Source = string(m.Source)

	bookmarkId := m.BookmarkID.String()
	item.BookmarkId = bookmarkId

	return item
}
