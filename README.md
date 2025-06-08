# gator

boot.dev "Build a Blog Aggregator in Go" project  
https://www.boot.dev/lessons/14b7179b-ced3-4141-9fa5-e67dbc3e5242

# TODO: CH5 L3 does not make sense

it is suposd to be develper documentation, but

-   there is no mention of the `goose` depencency for migration
-   source files are required to run the migration to create the tables
-   they would never run `go install`, becaues the binary is usless without a complete DB

# dependencies

-   Go https://webinstall.dev/go/
-   Postgres

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
    "db_url": "postgres://example"
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
# test, hit q to get back
SELECT version();
# exit psql
exit
```

-   create tables  
    `goose` migration program

```shell

```

# install app

```shell
go install github.com/scGetStuff/gator@latest

# test
gator
```

# commands to make gator do stuff

-   `register <NAME>`
-   `login <NAME>`
-   `reset` clears the data from
-   users
-   feeds
-   agg

-   addfeed
-   follow
-   unfollow
-   following
-   browse
    register
    login
    reset
    users
    feeds
    agg

addfeed
follow
unfollow
following
browse
