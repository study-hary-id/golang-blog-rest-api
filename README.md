# RESTful API for Blog App

Simple RESTful API to serve blog data to a blogging website, this is an example of RESTful API using Go without any
framework. The main purpose of this project is to be used on my own simple blog website.

## Endpoints

| HTTP Verb | Path         | Action                   | Resource |
| --------- | ------------ | ------------------------ | -------- |
| GET       | `/articles`  | Redirect to `/articles/` | Articles |
| GET       | `/articles/` | Get all articles         | Articles |

### Get all articles

```http request
GET /articles 301 Moved Permanently
```
```http request
GET /articles/ 200 OK
```

#### Successful response

```json
{
  "data": [
    {
      "type": "articles",
      "id": "1",
      "attributes": {
        "title": "JSON:API paints my bikeshed!",
        "date": "July 23, 2019",
        "body": "",
        "featured": true
      },
      "links": {
        "self": "https://example.com/articles/1",
        "cover": "https://example.com/static/img/1.jpg"
      }
    }
  ]
}
```
