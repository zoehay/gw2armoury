#!/bin/sh

docker run -d \
	--name db \
    --network armoury-network \
    -p 5432:5432 \
	-v /Users/zoehay/Projects/gw2armoury/secrets/db_name.txt:/run/secrets/db_name \
	-v /Users/zoehay/Projects/gw2armoury/secrets/db_password.txt:/run/secrets/db_password \
	-v /Users/zoehay/Projects/gw2armoury/secrets/db_user.txt:/run/secrets/db_user \
	-e POSTGRES_PASSWORD_FILE=/run/secrets/db_password \
    -e POSTGRES_USER_FILE=/run/secrets/db_user \
    -e POSTGRES_DB_FILE=/run/secrets/db_name \
	-e PGDATA=/var/lib/postgresql/data/pgdata \
	-v pgdata:/var/lib/postgresql/data \
	postgres
