#!/bin/bash

export COMPOSE_FILE=docker-compose.sc.yml
export COMPOSE_PROJECT_NAME=monstache_showcase

docker-compose down -v --remove-orphans
