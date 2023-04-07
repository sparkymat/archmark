// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package dbx

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Bookmark struct {
	ID        int64
	UserID    int64
	Url       string
	Title     string
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}

type SchemaMigration struct {
	Version int64
	Dirty   bool
}

type User struct {
	ID                int64
	Username          pgtype.Text
	Name              string
	EncryptedPassword string
	CreatedAt         pgtype.Timestamp
	UpdatedAt         pgtype.Timestamp
}