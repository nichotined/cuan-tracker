# Cuan Tracker Backend Service

Implementation in Go and PostgreSQL

## Migrate Commands

[Important] Install go-migrate with brew
> $ brew install golang-migrate

To create new migration script
> $ migrate create -ext sql -seq -dir script_name

To up
> $ ./binary migrate_up

To down
> $ ./binary migrate_down

## Server Commands

To start server
> $ ./binary server