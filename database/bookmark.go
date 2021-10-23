package database

import "github.com/sparkymat/archmark/model"

func (s *service) ListBookmarks(query string, page uint32, pageSize uint32) ([]model.Bookmark, error) {
	panic("unimplemented")
	/*
		var bookmarks []model.Bookmark

		offset := int((page - 1) * pageSize)
		stmnt := s.conn

		if query != "" {
			stmnt = stmnt.Where("to_tsvector(content) @@ to_tsquery(?)", query)
		} else {
			stmnt = stmnt.Order("created_at desc")
		}

		if result := stmnt.Offset(offset).Limit(int(pageSize)).Find(&bookmarks); result.Error != nil {
			return nil, result.Error
		}

		return bookmarks, nil
	*/
}

func (s *service) FindBookmark(id uint) (*model.Bookmark, error) {
	panic("unimplemented")
	/*
		bookmark := &model.Bookmark{}

		if result := s.conn.Find(bookmark, id); result.Error != nil {
			return nil, result.Error
		}

		return bookmark, nil
	*/
}

func (s *service) CreateBookmark(bookmark *model.Bookmark) error {
	panic("unimplemented")
	/*
		result := s.conn.Create(bookmark)

		return result.Error
	*/
}

func (s *service) MarkBookmarkCompleted(id uint) error {
	panic("unimplemented")
	/*
		result := s.conn.Model(&model.Bookmark{}).Where("id = ?", id).Update("status", "completed")

		return result.Error
	*/
}
