.PHONY: go swagger js rust python java

all:
	docker-compose up

go:
	docker-compose run --rm chirpstack-v3-api-go

swagger:
	docker-compose run --rm chirpstack-v3-api-swagger

js:
	docker-compose run --rm chirpstack-v3-api-js

rust:
	docker-compose run --rm chirpstack-v3-api-rust

python:
	docker-compose run --rm chirpstack-v3-api-python

java:
	docker-compose run --rm chirpstack-v3-api-java

java-current-user:
	CURRENT_UID="$(shell id -u):$(shell id -g)" CURRENT_HOME=$(HOME) docker-compose run --rm chirpstack-api-java-current-user
