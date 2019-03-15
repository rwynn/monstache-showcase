#!/bin/bash

# Use this script after you have successfully run create-showcase.sh
# This brings up only Elasticsearch and Kibana to visualize the data

export COMPOSE_FILE=docker-compose.analyze.yml
export COMPOSE_PROJECT_NAME=monstache_showcase

docker-compose down --remove-orphans ; docker-compose up
