# MK Device Manager
MK Device Manager is a service that meant to handle device messages and manage firmware versions

## Prerequisites

To develop mk-device-manager, you will need:

1. [Go](https://go.dev/doc/install) (version 1.21)

## Installation
Ensure all  pre-requisites are satisfied before carrying out installation and clone the repo

## Database Creation
Create a PostgreSQL database and name it appropriately. Check this [reference](https://github.com/Mike-Kimani/whitepointinventory/blob/master/.env#L2) .The database is named `whitepointinventory` in this case

## Apply all Database Migrations
``
cd internal/database/sqlc/schema
``

``
goose postgres postgres://{userName}:{password}@localhost:5432/{databaseName} up
``

#### Build
``
make build
``

#### Build and Start the Server
``
make run
``

## Documentation
The documentation is found [here](documentation)

