# Gin Skeleton

This is a simple skeleton project for Gin HTTP web framework.

## Prepare

1. Create database: `$ cat resources/sql/db.sql | sqlite3 storage/db/db.sqlite`

### Run it:

1. `$ go run server.go`
2. Browse to http://localhost:9000

### Run test

1. `$ go test`

### Notice

Set `storage/` folder permission to writable when deploy to production environment

## Key directories

* `models`: Model files
* `public`: Static file root
* `resources/sql`: sql dump file for sample database
* `routes`: Controller files
* `storage/log`: Log files
* `storage/db`: SQLite DB files
* `utils`: Helper and Utility files
* `views`: Template files

## Key files

* `server.go`: Entry point to application
* `server_test.go`: API layer unit test
* `config.toml`: Configuration
* `views/index.tpl.html`: Template file for the home page