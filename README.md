# Web app example

```bash
# Run web app
go run main.go bird_handlers.go

# Run unit tests
go test -v
```

Browse to

```bash
wget localhost:8080
wget localhost:8080/hello
wget localhost:8080/assets/
```

# GoLang Docker image

```bash
docker build -t gowebapp .
docker run -it --rm -p 8080:8080 --name my-running-app gowebapp
```

# Postgres Docker image

## Setup

```bash
docker pull postgres
# Start postgress container
docker run --name some-postgres -e POSTGRES_PASSWORD=mysecretpassword -d postgres
# Connect using psql
docker run -it --rm --link some-postgres:postgres postgres psql -h postgres -U postgres
#> password: mysecretpassword
```

```sql
-- Create the database
CREATE DATABASE bird_encyclopedia;
-- Enter the database
\c bird_encyclopedia

-- Create table
CREATE TABLE birds (
  id SERIAL PRIMARY KEY,
  species VARCHAR(256),
  description VARCHAR(1024)
);

-- Get info
\d
\d birds
select * from birds;

-- Insert a value
INSERT INTO birds (species, description) VALUES ('kanarie', 'Small yellow brid');
select * from birds;
select species, description from birds;

```

# Docker-compose

```bash
docker-compose -f docker-compose.yml up
```


# Reference
- https://www.sohamkamani.com/blog/2017/09/13/how-to-build-a-web-application-in-golang/
- https://hub.docker.com/_/golang/
- https://hub.docker.com/_/postgres/
