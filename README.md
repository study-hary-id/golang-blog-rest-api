# RESTful API for Blog App

Simple RESTful API to serve blog data to a blogging website, this is an example of RESTful API using Go without any
framework. The main purpose of this project is to be used on my own simple blog website.

## Endpoints

| HTTP Verb | Path        | Action            | Resource |
| --------- | ----------- | ----------------- | -------- |
| GET       | `/articles` | Show all articles | Articles |

### Show all articles

```http request
GET /articles
```

#### Successful response

```json
{
  "links": {
    "self": "https://example.com/articles",
    "next": "https://example.com/articles?page[offset]=2",
    "last": "https://example.com/articles?page[offset]=10"
  },
  "data": [{
    "type": "articles",
    "id": "1",
    "attributes": {
      "title": "JSON:API paints my bikeshed!",
      "date": "July 23, 2019",
      "comments": "0",
      "body": ""
    },
    "links": {
      "self": "https://example.com/articles/1",
      "cover": "https://example.com/static/img/1.jpg"
    }
  }]
}
```
