package data

import "database/sql"

type Models struct {
	blog    BlogModel
	comment CommentModel
}

func (m *Models) New(db *sql.DB) Models {
	return Models{
		blog: BlogModel{DB: db},
		comment: CommentModel{DB: db},
	}
}
