# Roblox Username Checker

Written in Go for speed
edit: it ended up not being fast. at all.

## Prerequisites

- [Go](https://go.dev) installed.

## Usage

- Insert the usernames you'd like to check into `usernames.txt`
- List of words [here](https://github.com/dk-e/word-list) to try
- Run the below command:

```go
go run main.go
```

- The usernames that are available (or banned) will appear in a newly created `available.txt` file.
