# ENT

## REGENERATE SCHEMA CODE

```bash
ent generate --target ./ent/generated ./ent/schema
```

## DESCRIBE GRAPH SCHEMA

```bash
ent describe ./ent/schema
```

# ATLAS

## GENERATE DATABASE SCHEMA

```bash
atlas schema inspect --url "postgres://postgres@localhost:5432/database?sslmode=disable" --format '{{ sql . }}' > schema.sql
```

# POSTGRES

## DUMP TABLES DATA

```bash
pg_dump --quote-all-identifiers --column-inserts --no-comments --host localhost --port 5432 --username postgres --no-password --dbname database --file /app/dump.sql
```

# DOCKER

## CREATE DOCKER IMAGE

```bash
docker build --tag voting-system:latest .
```

## CREATE AND RUN CONTAINER

```bash
docker run --detach -p 3000:3000 --name voting-system voting-system:latest
```

## LIST ALL CONTAINERS

```bash
docker container ls --all
```

## ATTACH TO CONTAINER IN TTY

```bash
docker exec --interactive --tty voting-system /bin/sh
```

# DOCKER COMPOSE

## CREATE AND RUN CONTAINERS

```bash
docker compose up --build --detach
```

# KONG

## PREPARE THE KONG DATABASE

```bash
docker run --network voting-system_default --rm --env "KONG_DATABASE=postgres" --env "KONG_PG_HOST=kong-database" --env "KONG_PG_PASSWORD=kongpass" kong:latest kong migrations bootstrap
```

## BACK UP CONFIGURATION

```bash
deck gateway dump -o kong.yaml
```

## RECREATE ENTITIES FROM BACKUP

```bash
deck gateway sync kong.yaml
```
