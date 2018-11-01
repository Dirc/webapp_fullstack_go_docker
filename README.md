# GoLang Webapp

Based on this great [blog of Soham Kamani](https://www.sohamkamani.com/blog/2017/09/13/how-to-build-a-web-application-in-golang/).

```bash
# Run unit tests
go test -v

# Run web app
go run main.go bird_handlers.go store.go
```

Browse to

```bash
wget localhost:8080
wget localhost:8080/hello
wget localhost:8080/assets/
```

## GoLang Docker image

```bash
docker build -t gowebapp .
docker run -it --rm -p 8080:8080 --name my-running-gowebapp gowebapp
```

When you develop new code, you will need to remove the image in order to get it rebuild. (COPY in the Dockerfile does not detect changes?)

```bash
docker image rm gowebapp
```

## Postgres Docker image

```bash
docker pull postgres
# Start postgress container
docker run --name some-postgres -e POSTGRES_PASSWORD=mysecretpassword -d postgres
# Connect using psql
docker run -it --rm --link some-postgres:postgres postgres psql -h postgres -U postgres
#> password: mysecretpassword
```

Play around with SQL.

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

## Docker-compose

```bash
docker-compose -f docker-compose.yml up
```

You can verify if the database is initialized using psql:

```bash
docker exec -it webappfullstackgodocker_db_1 psql -d bird_encyclopedia -U postgres -c "select * from birds;"
```

Cleaning things:

```bash
docker-compose -f docker-compose.yml rm --force
```

## Connect Go webapp to Postgres container

```go
connString := "host=db user=postgres password=secret dbname=bird_encyclopedia sslmode=disable"
```

## Reference

- [Blog Soham Kamani](https://www.sohamkamani.com/blog/2017/09/13/how-to-build-a-web-application-in-golang/)
- [Use GoLang image](https://hub.docker.com/_/golang/)
- [Use Postgres image](https://hub.docker.com/_/postgres/)


## ToDo
- [ ] Add unit test for store
- [ ] Add CI flow, Travis CI or Cirlce CI?
