docker-build:
	$(MAKE) -C backend docker-build

docker-up:
	$(MAKE) -C backend docker-up

db-up:
	$(MAKE) -C backend db-up

npm-init:
	$(MAKE) -C frontend npm-init

backend:
	$(MAKE) -C backend dev

frontend:
	$(MAKE) -C frontend dev

.PHONY: build backend frontend npm-init