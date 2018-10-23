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
  bird VARCHAR(256),
  description VARCHAR(1024)
);

-- Get info
\d
\d birds
select * from birds;

-- Insert a value
INSERT INTO birds (bird, description) VALUES ('kanarie', 'Small yellow brid');
select * from birds;
select bird, description from birds;

```


# Reference
- https://www.sohamkamani.com/blog/2017/09/13/how-to-build-a-web-application-in-golang/
- https://hub.docker.com/_/postgres/
