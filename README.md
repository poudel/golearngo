# Introduction

Tried to learn some go. This repo houses excercies that I've done or
am doing.

Most of these are basic examples from the`A tour of Go` tutorial
series. Some that stand out to me are:


## mirror.go

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
