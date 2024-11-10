# ENT
## REGENERATE SCHEMA CODE
```bash
ent genereate ./ent/schema
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
# DOCKER
## CREATE DOCKER IMAGE
```bash
sudo docker build --tag voting-system:latest .
```
## CREATE AND RUN CONTAINER
```bash
sudo docker run --detach -p 3000:3000 --name voting-system voting-system:latest
```
## LIST ALL CONTAINERS
```bash
sudo docker container ls --all
```
## ATTACH TO CONTAINER IN TTY
```bash
sudo docker exec --interactive --tty voting-system /bin/sh
```
# DOCKER COMPOSE
## CREATE AND RUN CONTAINERS
```bash
sudo docker compose up --build --detach
```