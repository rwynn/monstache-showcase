# Monstache showcase

This project shows how monstache can be applied to real data from data.gov.  The `mongoimport` tool will be used
to import 6.5 million records of crime data.

During the import monstache will be listening for change events on the entire MongoDB deployment and indexing 
those documents into Elasticsearch.  Before importing monstache will do a little bit of transformation on the 
data using a golang plugin to enable certain aggregations in Kibana. 

The golang plugin was used over a Javascript plugin after noticing a dramatic performance increase.

I recommend that your machine has at least 16GB RAM, 20GB free disk, and 4 or more CPU cores. You may be able to 
get away with less by decreasing the heap sizes for Elasticsearch in the docker-compose files.

First you will need to make sure you have `docker` and `docker-compose` installed.  On desktop systems like 
Docker Desktop for Mac and Windows, Docker Compose is included as part of those desktop installs.

The versions at this project creation time were:

```
Client:
 Version:           18.09.3
 API version:       1.39
 Go version:        go1.10.8
 Git commit:        774a1f4
 Built:             Thu Feb 28 06:40:58 2019
 OS/Arch:           linux/amd64
 Experimental:      false

Server: Docker Engine - Community
 Engine:
  Version:          18.09.3
  API version:      1.39 (minimum version 1.12)
  Go version:       go1.10.8
  Git commit:       774a1f4
  Built:            Thu Feb 28 05:59:55 2019
  OS/Arch:          linux/amd64
  Experimental:     false

docker-compose version 1.23.1, build b02f1306
docker-py version: 3.5.0
CPython version: 3.6.7
OpenSSL version: OpenSSL 1.1.0f  25 May 2017
```

Next you will want to download the public [dataset](https://catalog.data.gov/dataset/crimes-2001-to-present-398a4). You will
want the .CSV format.  Please read all the rules and caveats associated with the public dataset before proceeding.

When you have downloaded this large 1.5GB file you should copy it to the following location:

```
monstache-showcase/mongodb/scripts/data/crimes.csv
```

You are now ready to run docker-compose and start the import. 

```
cd monstache-showcase
./import-showcase.sh
```

The import will take a while.  During the process you will a see line like this coming from `mongoimport`:

```
c-data       | 2019-03-12T20:34:57.586+0000     imported 6820156 documents
```

That means that all the data has been loaded into MongoDB.  Now you must wait for the indexing to complete in 
Elasticsearch.  The process will periodically query the document count in Elasticsearch.  

You will see lines like this repeating forever:

```
c-config     | [
c-config     |   {
c-config     |     "health" : "green",
c-config     |     "status" : "open",
c-config     |     "index" : "chicago.crimes",
c-config     |     "uuid" : "4wShbV-LTq6-6paRsWataQ",
c-config     |     "pri" : "1",
c-config     |     "rep" : "0",
c-config     |     "docs.count" : "1198982",
c-config     |     "docs.deleted" : "0",
c-config     |     "store.size" : "359mb",
c-config     |     "pri.store.size" : "359mb"
c-config     |   }
c-config     | ]

```

The `doc_count` should approach 6820156.  If the document count only reaches 6820155 then that
means that MongoDB imported a record for the first CSV line (the headers) but this was not stored in Elasticsearch. This
line is actually not needed because the fields with types are defined outside the file.  You can delete the first header 
line before running the import if you want the document count to match exactly.

Once all the data is loaded into Elasticsearch you can bring down the containers with Ctrl-C or:

```
cd monstache-showcase
./stop-showcase.sh
```

At this point you have indexed all the data and no longer should run `import-showcase.sh` as that will index all the data
again. The import process stores the Elasticsearch data in a docker volume so it will persist between runs until you 
delete the volume.

The last step is to fire up Kibana to analyze it. To do this start only Elasticsearch and Kibana with:

```
cd monstache-showcase
./view-showcase.sh
```

Once the containers are up and healthy you can go to http://localhost:5601 on the host to load Kibana and explore data.  

In Kibana you can start from scratch and define an index-pattern. However, I recommend that you import the 
file named `export.json` from the root of monstache-showcase to get a head start.

To import you will want to go to `Management` -> `Saved Objects` and then click `Import` and upload `export.json`.

You will also want to go under `Management` -> `Advanced Settings` in Kibana and set `Timezone for date formatting`
to `UTC` to display dates correctly.

When you are finished analyzing in Kibana you can run `./stop-showcase.sh` to bring down the containers.

If you want to tear down everything and delete all the associated data you can run `./clean-showcase.sh`.  
This stops the containers and deletes the associated docker volumes.  

Please open an issue with any feedback you might have.  Thanks!

