# Introduction

Tried to learn some go. This repo houses excercies that I've done or
am doing.

Most of these are basic examples from the`A tour of Go` tutorial
series. Some that stand out to me are:


## mirror.go

### Time

Return current server time in different formats:

Sending a request to `http://localhost:8799/now/` will yield a response
similar to the following one:


```json
{
  "rfc3339": "2018-06-07T13:24:07+05:45",
  "ansi_c": "Thu Jun  7 13:24:07 2018",
  "unix_date": "Thu Jun  7 13:24:07 +0545 2018",
  "unix_seconds": 1528357147,
  "unix_nano_seconds": 1528357147091860422,
  "ip": "[::1]"
}
// POST http://localhost:8799/now/
// HTTP/1.1 200 OK
// Content-Type: application/json
```


### Status code
Mirror the status code requested through URL. Example:

A request on `http://localhost:8799/status/400/` will return
a response like this:

```json
{
  "message": "Bad Request",
  "status_code": 400,
  "ip": "[::1]",
  "method": "GET"
}
// GET http://localhost:8799/status/400/
// HTTP/1.1 400 Bad Request
// Content-Type: application/json
```

Another example: A `POST` request on `http://localhost:8799/status/200/` will
return a response like this:

```json
{
  "message": "OK",
  "status_code": 200,
  "ip": "[::1]",
  "method": "POST"
}
// POST http://localhost:8799/status/200/
// HTTP/1.1 200 OK
// Content-Type: application/json
```

Other HTTP methods such as `PATCH`, `PUT` etc. are supported.


## finder.go

Simple and dumb program that takes a string and an optional path and
recursively searches files with name containing the given string.
