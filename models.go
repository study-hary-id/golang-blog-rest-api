package main

import (
	"database/sql"
	"fmt"
	"net/url"
	"time"
)

// Our in-memory datastore,
// datastore used to map the list of users.
// Remember to guard map access with a mutex for concurrent access.
//type datastore struct {
//	m map[string]user
//	*sync.RWMutex
//}

type article struct {
	ID       int       `json:"id"`
	Title    string    `json:"title"`
	Body     string    `json:"body"`
	Date     time.Time `json:"date"`
	Featured bool      `json:"featured"`
	Cover    string    `json:"cover"`
}

// getSelfLink returns the link to current article.
func (a *article) getSelfLink() string {
	return fmt.Sprintf("http://localhost:8000/articles/%d", a.ID)
}

/*
API to communicate between article with datastore.
*/

// bindRowsArticles binds the database query rows then returns article slice.
func bindRowsArticles(rows *sql.Rows) (interface{}, error) {
	var articles []article
	for rows.Next() {
		article := article{}
		err := rows.Scan(
			&article.ID,
			&article.Title,
			&article.Body,
			&article.Featured,
			&article.Cover,
			&article.Date,
		)
		if err != nil {
			return nil, err
		}
		articles = append(articles, article)
	}
	return articles, nil
}

// buildQueryArticles builds database query from URL queries,
// returns the query and it's required values.
func buildQueryArticles(queries url.Values) (string, []interface{}) {
	var (
		query    = `SELECT * FROM articles `
		featured = queries.Get("featured")
		limit    = queries.Get("limit")
		sort     = queries.Get("sort")
		values   []interface{}
	)

	// Checks featured query.
	if featured != "" {
		query += `WHERE featured=? `
		values = append(values, URLConv(featured))
	}

	// TODO: Known bug maybe, cannot assign DESC/ASC to `?`.
	// That's why they use hardcode SQL syntax.
	// TODO: Replace ORDER BY `id` to `created_at`.
	switch sort {
	case "desc":
		query += `ORDER BY id DESC `
	case "asc":
		query += `ORDER BY id ASC `
	}

	// Check limit query.
	if limit != "" {
		query += `LIMIT ? `
		values = append(values, URLConv(limit))
	}
	return query, values
}

// getArticles returns the list of articles, from URL queries.
func getArticles(queries url.Values) ([]article, error) {
	query, values := buildQueryArticles(queries)
	// Query to database with values.
	rows, err := db.Query(query, values...)
	if err != nil {
		return nil, err
	}
	// Bind rows to articles.
	articles, err := bindRowsArticles(rows)
	if err != nil {
		return nil, err
	}
	return articles.([]article), nil
}
