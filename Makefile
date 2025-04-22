# Makefile for the first time setup of the project
run-docker-build:
	$(MAKE) -C backend docker_build

run-db-up:
	$(MAKE) -C backend db_up

run-npm-init:
	$(MAKE) -C frontend npm_init

run-golang:
	$(MAKE) -C backend dev

run-angular:
	$(MAKE) -C frontend dev

# Makefile for the next time setup of the project
run-docker-up:
	$(MAKE) -C backend docker_up

run: run-golang && run-angular

.PHONY: build backend frontend run