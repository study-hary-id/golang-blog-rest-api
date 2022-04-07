package main

import (
	"sync"
	"time"
)

// links contain two resource, one is detail article,
// and the second is the cover of the article.
type links struct {
	Self  string `json:"self"`
	Cover string `json:"cover"`
}

// attributes contain the content of an article and
// metadata for the article itself.
type attributes struct {
	Title    string    `json:"title"`
	Body     string    `json:"body"`
	Date     time.Time `json:"date"`
	Featured bool      `json:"featured"`
}

// article contains structure for article resource.
type article struct {
	Type       string     `json:"type"`
	ID         string     `json:"id"`
	Attributes attributes `json:"attributes"`
	Links      links      `json:"links"`
}

// Our in-memory datastore,
// datastore used to map the list of users.
// Remember to guard map access with a mutex for concurrent access.
type datastore struct {
	m map[string]article
	*sync.RWMutex
}

// payload is a structure for json response data.
type payload struct {
	Data []article `json:"data"`
}
