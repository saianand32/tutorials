
1. Install and run postgress container

$ docker pull postgres:15.1-alpine

$ docker run --env POSTGRES_PASSWORD=foobaz --publish 5432:5432 postgres:15.1-alpine