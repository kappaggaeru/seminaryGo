# Go seminary
This is an API REST made using Go

## How to use it?
First you'll need to run the application
```
go run main.go
```
Then you can use either Postman or a browser or whatever you want to make http requests
### REST
|Method|Usage|
|-|-|
|GET|localhost:8080/movies/{id}|
|POST|localhost:8080/movies/add|
|PUT|localhost:8080/movies/update/{id}|
|DELETE|localhost:8080/movies/delete/{id}|
### JSON
```json
{
    "title":"pulp fiction",
    "year":1994,
    "director":"quentin tarantino"
}
```