#!/bin/sh

docker network create -d bridge armoury-network

sh db/init-db.sh