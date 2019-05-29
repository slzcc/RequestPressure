SHELL := /bin/bash
NAME := "rq"
VERSION := "latest"

build:
	@docker build -t $(NAME):$(VERSION) .
	@docker run --name $(NAME)_$(VERSION) -d --rm --entrypoint "" $(NAME):$(VERSION) sleep 10
	@docker cp $(NAME)_$(VERSION):/rp .
	@docker cp $(NAME)_$(VERSION):/rpp .
