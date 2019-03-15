@ECHO OFF

SET COMPOSE_FILE="docker-compose.sc.yml"
SET COMPOSE_PROJECT_NAME="monstache_showcase"

docker-compose down -v --remove-orphans
