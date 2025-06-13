# gator

boot.dev "Build a Blog Aggregator in Go" project  
https://www.boot.dev/lessons/14b7179b-ced3-4141-9fa5-e67dbc3e5242

# Notes

## TODO

-   user name is case sensitive
-   there was a note in the lesson about handling date format from feeds, all the tests I ran were consistant, so I didn't do any extra yet

## doing an install does not make sense

-   I don't see why someone would install the app, binary is ulsess by itself
-   they need the source `sql/schema` files to create the tables
-   so eiter have the repo localy like a sane person; or dig around in the install files, `go/pkg/mod/github.com/...`

# dependencies

## Go

https://webinstall.dev/go/

## goose

I had to run this twice, it has a bunch of dependencies

```shell
go install github.com/pressly/goose/v3/cmd/goose@latest
```

## Postgres

```shell
# install postgres
sudo apt update
sudo apt install postgresql postgresql-contrib
psql --version

# set system user password
sudo passwd postgres
boot

# start server
sudo service postgresql start
```

# configuraion

-   create `~/.gatorconfig.json`

```json
{
    "db_url": "postgres://postgres:postgres@localhost:5432/gator?sslmode=disable"
}
```

-   create `gator` DB

```shell
sudo -u postgres psql
CREATE DATABASE gator;
# connect
\c gator
# set DB user password
ALTER USER postgres PASSWORD 'postgres';
# test
SELECT version();
# exit psql
exit
```

# install app

```shell
go install github.com/scGetStuff/gator@latest
```

## create tables

```shell
cd sql/schema
goose postgres "postgres://postgres:postgres@localhost:5432/gator" up
```

## test

```shell
gator
```

# commands to make gator do stuff

-   `register <NAME>`
-   `login <NAME>`
-   `reset` clears the data
-   `users` lists all registered
-   `addfeed "<NAME>" "<URL>"` add feed to the DB
-   `feeds` lists feed records

-   `follow`
-   `unfollow`
-   `following`
-   `browse`

-   `agg <time>` infinite loop storing posts to the DB from any feeds added  
    https://pkg.go.dev/time#ParseDuration
