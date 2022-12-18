# test-sharing-vison

Install [Docker](https://www.docker.com/get-started) and the following useful Go tools to your system:

[golang-migrate/migrate](https://github.com/golang-migrate/migrate#cli-usage) for apply migrations

Run project by this command:

make docker.run

To stop the project use this command :

make docker.run


if the migrate fail use :

migrate -path PATH_TO_THE_MIGRATIONS_FOLDER up -database "mysql://root:mysqlpw@tcp(localhost:3306)/article" up

