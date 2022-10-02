# Calibre Database

## Connection

There are a lot of sqlite3 drivers in Golang. The most common used is https://github.com/mattn/go-sqlite3.
But it requires cgo enabled.

We choose the https://github.com/zombiezen/go-sqlite as the default sqlite client which has the best performance in pure Go.
It doesn't provide any Go sql driver implementation.
You should use SQL and prepare statements for getting the best performance.
