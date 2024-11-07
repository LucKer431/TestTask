.PHONY: build up down clean

build:
    docker-compose build

up:
    docker-compose up -d

down:
    docker-compose down

clean:
    docker-compose down --volumes

push:
    docker tag testtask_app:latest luck000000/testtask:latest
    docker push luck000000/testtask:latest