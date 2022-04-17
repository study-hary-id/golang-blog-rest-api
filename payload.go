package main

// constructArticle returns a new modified article structure for response json.
func constructArticle(article article) map[string]interface{} {
	return map[string]interface{}{
		"type": "articles",
		"id":   article.ID,
		"attributes": map[string]interface{}{
			"title":    article.Title,
			"date":     article.Date,
			"body":     article.Body,
			"featured": article.Featured,
		},
		"links": map[string]interface{}{
			"self":  article.getSelfLink(),
			"cover": article.Cover,
		},
	}
}

// constructItems checks the actual type of `data` and
// calls proper function to process the generic `data`,
// then returns the constructed items as a key value pair.
func constructItems(data interface{}) map[string]interface{} {
	items := map[string]interface{}{
		"data": nil,
	}
	switch data.(type) {
	case []article:
		articles := data.([]article)
		constructedItems := make([]interface{}, len(articles))
		for i, article := range articles {
			constructedItems[i] = constructArticle(article)
		}
		items["data"] = constructedItems
	default:
		items["data"] = data
	}
	return items
}
