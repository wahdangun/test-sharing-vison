# test-sharing-vison

Install [Docker](https://www.docker.com/get-started) and the following useful Go tools to your system:

[golang-migrate/migrate](https://github.com/golang-migrate/migrate#cli-usage) for apply migrations

Run project by this command:

make docker.run

To stop the project use this command :

make docker.run


if the migrate fail use :

migrate -path PATH_TO_THE_MIGRATIONS_FOLDER up -database "mysql://root:mysqlpw@tcp(localhost:3306)/article" up

url postman :
https://www.postman.com/bold-spaceship-42896/workspace/test-sharing-vision/request/13477179-ca5471a9-69c9-4bee-8f09-fbadbc5428f3
