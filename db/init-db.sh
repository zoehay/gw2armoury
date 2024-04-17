#!/bin/sh

docker run -d \
	--name db \
    --network armoury-network \
    -p 5432:5432 \
	-e POSTGRES_PASSWORD_FILE=/run/secrets/db_password \
    -e POSTGRES_USER_FILE=/run/secrets/db_user \
    -e POSTGRES_DB_FILE=/run/secrets/db_name \
	-e PGDATA=/var/lib/postgresql/data/pgdata \
	-v ./data:/var/lib/postgresql/data \
	postgres