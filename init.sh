#!/bin/bash

docker network inspect "armoury-network" > /dev/null
if [ $? -ne 0 ]; then
    docker network create -d bridge armoury-network
else 
    echo "network already created"
fi

docker inspect "db" > /dev/null
if [ $? -ne 0 ]; then
    sh db/init-db.sh
else 
    echo "db already created"
fi

sh backend/init-backend.sh