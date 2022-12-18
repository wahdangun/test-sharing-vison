APP_NAME = test-sharing-vision

docker.run: docker.sharingvision docker.sharingvision docker.redis migrate.up

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
		--name cgapp-redis \
		--network dev-network \
		-p 6379:6379 \
		redis

docker.stop: docker.stop.sharingvision docker.stop.visionmysql docker.stop.redis

docker.stop.sharingvision:
	docker stop sharing

docker.stop.mysql:
	docker stop visionmysql

swag:
	swag init