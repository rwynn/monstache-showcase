@ECHO OFF

:: Use this script after you have successfully run create-showcase.sh
:: This brings up only Elasticsearch and Kibana to visualize the data

SET COMPOSE_FILE="docker-compose.analyze.yml"
SET COMPOSE_PROJECT_NAME="monstache_showcase"

docker-compose down --remove-orphans
docker-compose up
