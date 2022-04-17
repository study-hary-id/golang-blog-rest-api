package main

import (
	"fmt"
	"time"
)

// Our in-memory datastore,
// datastore used to map the list of users.
// Remember to guard map access with a mutex for concurrent access.
//type datastore struct {
//	m map[string]user
//	*sync.RWMutex
//}

// article is defined article type.
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

// getArticles returns the list of all articles.
func getArticles() ([]article, error) {
	rows, err := db.Query(`SELECT * FROM articles`)
	if err != nil {
		return nil, err
	}

	// Bind rows to articles.
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
