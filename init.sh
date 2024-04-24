#!/bin/bash

docker network inspect "armoury-network" > /dev/null
if [ $? -ne 0 ]; then
    docker network create -d bridge armoury-network
else 
    echo "network already created"
fi

sh db/init-db.sh && \

sh backend/init-backend.sh