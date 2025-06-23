# japa-tracker
Golang based rest app for tracking japa.

# endpoints

| endpoint| verb | commemts |
|-|-|-|
|/japa/{date}|GET| count for a date | 
|/japa|POST|adding/updating count for a date|
|/japa-count/{date}|GET| getting complete count of japa till the given date|

# Sample curl commands

```
curl --header "Content-Type: application/json" --request POST --data '{"date":"2025-10-10","count":32}' http://localhost:8080/japa

curl --request GET  http://localhost:8080/japa-count/2025-01-23

curl --request GET  http://localhost:8080/japa/2025-01-23
```
