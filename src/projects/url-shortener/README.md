# A simple URL shortener 

## To run the API
```bash
$ cd ~/src/projects/url-shortener
$ go run main.go
```
## Requirements 
- Go installed 
- MySQL 

## Testing 

To shorten the long URL: -
`POST` http://localhost:5050/shorten

```bash
curl --location --request POST 'http://localhost:5050/shorten' \
--header 'Content-Type: application/json' \
--data-raw '{
    "long_url":"https://www.whitesourcesoftware.com/free-developer-tools/blog/golang-dependency-management"
}'
```

Response :- 
```json
{
    "id": 6,
    "slug": "JYOLkQP",
    "short_url": "http://localhost:5050/JYOLkQP",
    "long_url": "https://www.whitesourcesoftware.com/free-developer-tools/blog/golang-dependency-management",
    "created_at": "2021-10-11T15:37:00+01:00"
}
```

Copy the shortened URL and paste in your browser, it redirects to the long URL. 