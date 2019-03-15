#!/bin/bash

# Start with this script to populate the Elasticsearch volume with data
# This script tails MongoDB with monstache while the data import is running
# All data imported into MongoDB will be indexed in Elasticsearch.

# Note that you will need sufficient memory and disk space.  At least 16GB RAM
# and 20 GB available disk is recommended.  

# If monstache reports errors about losing the capped position during import then
# you may need to increase the size of the oplog in mongo-db-setup.sh. The default
# is 8GB oplog size.

# It may take a while for all data to reach Elasticsearch even after the mongo
# import is completed.  You may need to periodically check the count of the mongo
# collection db.crimes against the count of the index db.crimes.  Eventually, these
# should become equal. The lag is partially due to the setting of the refresh_interval
# to a high value of 30s such as not to overly tax Elasticsearch. The total document 
# count is 6,820,155. After the mongodb import a script will periodically check the size
# of the Elasticsearch index.

export COMPOSE_FILE=docker-compose.sc.yml
export COMPOSE_PROJECT_NAME=monstache_showcase

docker-compose down --remove-orphans ; docker-compose up --force-recreate --build
