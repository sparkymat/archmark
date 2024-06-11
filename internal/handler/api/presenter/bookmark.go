package presenter

import (
	"time"

	"github.com/sparkymat/archmark/internal/dbx"
)

type Bookmark struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Url         string  `json:"url"`
	Description string  `json:"description"`
	Image       *string `json:"image"`
	Transcript  string  `json:"transcript"`
	UserId      string  `json:"userId"`
	CreatedAt   string  `json:"createdAt"`
	UpdatedAt   string  `json:"updatedAt"`
}

func BookmarkFromModel(m dbx.Bookmark) Bookmark {
	item := Bookmark{
		ID:        m.ID.String(),
		CreatedAt: m.CreatedAt.Time.Format(time.RFC3339),
		UpdatedAt: m.UpdatedAt.Time.Format(time.RFC3339),
	}

	item.Title = m.Title

	item.Url = m.Url

	item.Description = m.Description

	item.Transcript = m.Transcript

	if m.Image.Valid {
		item.Image = &m.Image.String
	}

	userId := m.UserID.String()
	item.UserId = userId

	return item
}
