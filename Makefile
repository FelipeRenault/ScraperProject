build-app:
	docker build -t pelando -f etc/docker/dockerfiles/app/Dockerfile .

build-mysql:
	docker build -t pelando-mysql -f etc/docker/dockerfiles/mysql/Dockerfile .

build: build-mysql build-app

run:
	docker-compose up --attach app
