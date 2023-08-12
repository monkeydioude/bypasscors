all: image-build run

run:
	docker run -it -p 8080:8080 --rm --name bypasscors drannoc/bypasscors

image-build:
	docker build --network="host" -t drannoc/bypasscors .

docker-push: check-env
	docker login -u $(DOCKER_USER) --password $(DOCKER_PASSWD)
	docker push drannoc/bypasscors

update-remote: image-build docker-push

build:
	docker run --rm --user "$(id -u)":"$(id -g)" --network="host" -v "$(PWD)":/usr/src/bypasscors -w /usr/src/bypasscors golang:alpine go build -v -buildvcs=false

build-prod: build
	rm -rf /home/mkd/go/bin/bypasscors
	mv bypasscors /home/mkd/go/bin

check-env:
ifndef DOCKER_USER
	$(error DOCKER_USER is undefined)
endif
ifndef DOCKER_PASSWD
	$(error DOCKER_PASSWD is undefined)
endif