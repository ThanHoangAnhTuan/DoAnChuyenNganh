# Makefile for the first time setup of the project
run-docker-build:
    $(MAKE) -C golang docker_build

run-db-up:
	$(MAKE) -C golang db_up

run-npm-init:
	$(MAKE) -C angular npm_init

run-golang:
    $(MAKE) -C golang dev

run-angular:
	$(MAKE) -C angular dev


# Makefile for the next time setup of the project
run-docker-up:
    $(MAKE) -C golang docker_up

run: run-golang && run-angular

.PHONY: build golang angular run