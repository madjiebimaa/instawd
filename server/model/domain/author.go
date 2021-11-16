package domain

import "database/sql"

type Author struct {
	Id          string
	Name        string
	Link        sql.NullString
	Bio         sql.NullString
	Description sql.NullString
	QuoteCount  int
}
