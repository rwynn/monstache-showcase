#!/bin/bash

echo "************************************************************"
echo "Importing Data"
echo "************************************************************"

set -eo pipefail;

mongoimport --authenticationDatabase "admin" --host "mongo-0" -u "root-user" -p "password" -d "chicago" -c "crimes" --type "csv" --columnsHaveTypes --parseGrace=skipField --fields="ID.string(),Case Number.string(),Date.date(01/02/2006 03:04:05 PM MST),Block.string(),IUCR.string(),Primary Type.string(),Description.string(),Location Description.string(),Arrest.boolean(),Domestic.boolean(),Beat.string(),District.string(),Ward.string(),Community Area.string(),FBI Code.string(),X Coordinate.int64(),Y Coordinate.int64(),Year.int32(),Updated On.date(01/02/2006 03:04:05 PM MST),Latitude.double(),Longitude.double(),Location.auto()" --file="data/crimes.csv"

