package helper

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/madjiebimaa/go-random-quotes/app"
	"github.com/madjiebimaa/go-random-quotes/model/domain"
)

func fileToBytes(location string) []byte {
	file, err := os.Open(location)
	PanicIfError(err)
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	PanicIfError(err)

	return bytes
}

func ToSlugFromAuthorName(name string) string {
	res := strings.ReplaceAll(strings.ToLower(name), " ", "-")
	return res
}

type Tags struct {
	Tags []Tag `json:"tags"`
}

type Tag struct {
	Id   string `json:"_id"`
	Name string `json:"name"`
}

func TagsNoSQLToSQL(location string) {
	bytes := fileToBytes(location)
	var tags Tags

	json.Unmarshal(bytes, &tags)

	ctx := context.Background()

	tx, err := app.NewDB().Begin()
	PanicIfError(err)
	defer CommitOrRollBack(tx)

	SQL := "INSERT INTO quote_tag (id, name) VALUES (?, ?)"
	stmt, err := tx.PrepareContext(ctx, SQL)
	PanicIfError(err)
	defer stmt.Close()

	for i := 0; i < len(tags.Tags); i++ {
		stmt.ExecContext(ctx, tags.Tags[i].Id, tags.Tags[i].Name)
	}
}

type Authors struct {
	Authors []Author `json:"authors"`
}

type Author struct {
	Id          string `json:"_id"`
	Name        string `json:"name"`
	Link        string `json:"link"`
	Bio         string `json:"bio"`
	Description string `json:"description"`
	QuoteCount  int    `json:"quoteCount"`
}

func AuthorsNoSQLToSQL(location string) {
	bytes := fileToBytes(location)
	var authors Authors

	json.Unmarshal(bytes, &authors)

	ctx := context.Background()

	tx, err := app.NewDB().Begin()
	PanicIfError(err)
	defer CommitOrRollBack(tx)

	SQL := "INSERT INTO author (id, name, link, bio, description, quote_count) VALUES (?, ?, ?, ?, ?, ?)"
	stmt, err := tx.PrepareContext(ctx, SQL)
	PanicIfError(err)
	defer stmt.Close()

	for i := 0; i < len(authors.Authors); i++ {
		_, err := stmt.ExecContext(ctx, authors.Authors[i].Id, authors.Authors[i].Name, authors.Authors[i].Link, authors.Authors[i].Bio, authors.Authors[i].Description, authors.Authors[i].QuoteCount)
		PanicIfError(err)
	}
}

type Quotes struct {
	Quotes []Quote `json:"quotes"`
}

type Quote struct {
	Id       string `json:"_id"`
	Content  string `json:"content"`
	AuthorId string `json:"authorId"`
}

func QuotesNoSQLToSQL(location string) {
	bytes := fileToBytes(location)
	var quotes Quotes

	json.Unmarshal(bytes, &quotes)

	ctx := context.Background()

	tx, err := app.NewDB().Begin()
	PanicIfError(err)
	defer CommitOrRollBack(tx)

	SQL := "INSERT INTO quote (id, content, author_id) VALUES (?, ?, ?)"
	stmt, err := tx.PrepareContext(ctx, SQL)
	PanicIfError(err)
	defer stmt.Close()

	for i := 0; i < len(quotes.Quotes); i++ {
		stmt.ExecContext(ctx, quotes.Quotes[i].Id, quotes.Quotes[i].Content, quotes.Quotes[i].AuthorId)
	}
}

func AddSlugToAuthors() {
	ctx := context.Background()

	tx, err := app.NewDB().Begin()
	PanicIfError(err)

	SQL := "SELECT id, name, link, bio, description, quote_count FROM author"
	rows, err := tx.QueryContext(ctx, SQL)
	PanicIfError(err)

	var authors []domain.Author
	for rows.Next() {
		var author domain.Author
		rows.Scan(&author.Id, &author.Name, &author.Link, &author.Bio, &author.Description, &author.QuoteCount)
		authors = append(authors, author)
	}
	rows.Close()

	SQL = "UPDATE author SET slug = ? WHERE id = ?"
	stmt, err := tx.PrepareContext(ctx, SQL)
	PanicIfError(err)

	for i := 0; i < len(authors); i++ {
		slug := ToSlugFromAuthorName(authors[i].Name)
		stmt.ExecContext(ctx, slug, authors[i].Id)
	}
	stmt.Close()

	tx.Commit()
}

func QueryToStruct(ctx context.Context, keyCtx string, result interface{}) {
	query := ctx.Value(keyCtx)
	byteData, _ := json.Marshal(query)
	json.Unmarshal(byteData, &result)
}
