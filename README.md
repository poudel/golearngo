# Introduction

Tried to learn some go. This repo houses excercies that I've done or
am doing.

Most of these are basic examples from the`A tour of Go` tutorial
series. Some that stand out to me are:


## mirror.go

Mirror the status code requested through URL. Example:

A request on `http://localhost:8799/status/400/` will return
a response like this:

```
Bad Request
GET http://localhost:8799/status/400/
HTTP/1.1 400 Bad Request
```

Another example: A request on `http://localhost:8799/status/200/` will
return a response like this:

```
OK
GET http://localhost:8799/status/200/
HTTP/1.1 200 OK
```


## finder.go

Simple and dumb program that takes a string and an optional path and
recursively searches files with name containing the given string.
