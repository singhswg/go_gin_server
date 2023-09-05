SHELL=bash
HOST_OS=$(shell uname)

.DEFAULT_GOAL := env

env: osdep-$(HOST_OS) venv dep
osdep-Linux:
	sudo apt-get install python3-pip python3-venv virtualenv unzip

osdep-Darwin:
	pip3 install virtualenv

venv:
	rm -rf .env
	python3 -m venv .env
	.env/bin/pip install --upgrade pip

dep: pip
	PIP_CONFIG_FILE=./env/pip.conf .env/bin/pip install -r requirements.txt --upgrade

# Assuming the env variables are already set for make build-image. You can also run the script available in ./postgres
build-image:
	docker build --build-arg="PG_USER=${PG_USER}" --build-arg="PG_PASS=${PG_PASS}" -t api-db:base . 

run-local:
	docker run -it --rm -p 8080:8080 --name api-db api-db:base 

run-local-host-network:
	docker run -it --network host --name api-db api-db:base 

local-postgres:
	cd postgres && docker compose up

# One could also create a custom network and run the containers in that network
