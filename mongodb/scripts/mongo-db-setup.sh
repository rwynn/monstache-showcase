#!/bin/bash

echo "************************************************************"
echo "Setting up database"
echo "************************************************************"

set -eo pipefail;

./mongo-engine-wait.sh

./mongo-rep-set-wait.sh

mongo  admin -u "$MONGO_USER_ROOT_NAME" -p "$MONGO_USER_ROOT_PASSWORD" --eval "db.adminCommand({replSetResizeOplog: 1, size: 8000})"
