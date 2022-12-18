APP_NAME = test-sharing-vision
MIGRATIONS_FOLDER =  ${CURDIR}/migrations
DATABASE_URL = mysql://root:mysqlpw@tcp(localhost:3306)/article

docker.run: docker.sharingvision docker.visionmysql docker.redis migrate.up

migrate.up:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DATABASE_URL)" up

migrate.down:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DATABASE_URL)" down


docker.sharingvision.build:
	docker build -t sharingvision .

docker.sharingvision: docker.sharingvision.build
	docker run --rm -d \
		--name sharingvision \
		-p 8080:8080 \
		fiber

docker.visionmysql:
	docker run --rm -d \
		--name visionmysql \
		-e MYSQL_ROOT_PASSWORD=mysqlpw \
		-e MYSQL_DATABASE=article \
		-p 3306:3306 \
		mysql

docker.redis:
	docker run --rm -d \
		--name vision-redis \
		-p 6379:6379 \
		redis

docker.stop: docker.stop.sharingvision docker.stop.visionmysql docker.stop.redis

docker.stop.sharingvision:
	docker stop sharingvision

docker.stop.visionmysql:
	docker stop visionmysql
docker.stop.redis:
	docker stop vision-redis	

swag:
	swag init