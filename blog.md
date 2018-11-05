# Idea for blog

## Dockerize Fullstack Go webapp

While I was following a nice tutorial on creating a webapp in Go, I decided to improve this setup by dockerizing it. It turned out to be quite easy and I will show you how you can do this.

## What to do

What are we going to do? The blog is about dockerizing a simple Go webapp. This webapp exist of some Go code and a Postgres database. We can brake the task down into four parts:

- Get Go code runnning in a Go container
- Get database running in a Postgres container
- Manage both using Docker Compose
- Ensure webapp is using the Postgres container as database.

Before we proceed we need a little more knowledge of this webapp we are working with.

We can start the webapp.

```bash
# Run web app
go run main.go bird_handlers.go store.go
```

We can verify if it is working on some url's.

```bash
wget localhost:8080
wget localhost:8080/hello
wget localhost:8080/assets/
```

There is a database with a table called `bird_encyclopedia` and with two columns `species` and `description`.

More details of this webapp can be found in this great [blog/tutorial of Soham Kamani](https://www.sohamkamani.com/blog/2017/09/13/how-to-build-a-web-application-in-golang/).

## Get GO code runnning in a Go container

To get the webapp running in a container, we will us the [official GoLang image](https://hub.docker.com/_/golang/). Since we have a simple webapp, we do not need to add much to the example `Dockerfile`. We only add a more recent version of the golang image and expose a port.

```docker
FROM golang:1.11

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 8080

CMD ["app"]
```

The `go get -d` command downloads all dependencies of the webapp and `go install` instals the webapp. After installing the webapp it becomes available as an executable `app`, i.e. the workdir directory name.

We can now try to build the image and run it.

```bash
docker build -t gowebapp .
docker run -it --rm -p 8080:8080 --name my-running-gowebapp gowebapp
```

Verify the webapp by browse to one of the uri's

# Get database running in a Postgres container

We start again with [the official Postgres image](https://hub.docker.com/_/postgres/).

We can start a Postgres container.

```bash
# Start postgress container
docker run --rm --name some-postgres -e POSTGRES_PASSWORD=mysecretpassword -d postgres
```

After it is started, we can and connect to it with psql and play around with SQL.

```bash
# Connect using psql
docker run -it --rm --link some-postgres:postgres postgres psql -h postgres -U postgres
#> password: mysecretpassword
```

We can initialise the database.

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
```

And we can insert a value and retreive information.

```sql
-- Insert a value
INSERT INTO birds (species, description) VALUES ('kanarie', 'Small yellow brid');
select * from birds;
select species, description from birds;

-- Get info
\d
\d birds
select * from birds;
```

In other words, we can manage the database the way we need it, except for one last thing. We want the database to initalise on startup. This can be achieved by adding a initdb.sql file and put it in `/docker-entrypoint-initdb.d` in the container. A simple way to do this is to add all files of the current directory using \`pwd\`.

```bash
docker run --rm --name db -v `pwd`:/docker-entrypoint-initdb.d -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=secret -e POSTGRES_DB=bird_encyclopedia -d postgres
```

We can verify the result again with psql.

```bash
# Connect using psql
docker run -it --rm --link db:postgres postgres psql -h postgres -U postgres bird_encyclopedia
#> password: secret
```

## Docker Compose

```bash
docker-compose up
```

You can verify if the database is initialized using psql:

```bash
docker exec -it webappfullstackgodocker_db_1 psql -d bird_encyclopedia -U postgres -c "select * from birds;"
```

Cleaning things:

```bash
docker-compose -f docker-compose.yml rm --force
```

## skeleton

- Intro: Ready to learn some Go Lang
  - (After Games With Go)
  - Wanted to create a simple webapp and found this excelent blog.
  - I got exited and decided to spice things up a bit by dockerise this webapp and database setup.
  - Which was surprisingly easy!

- Tool of choice: VS Code. Since it has a nice Go plugin.

- Steps:
  - Follow blog. It starts with a simple `Hello World`. Note that this example is still available at the `/hello` uri. Hence we can use it to verify if the webapp is running properly in the container.
  - Get webapp runnning in Go Docker
  - Get database running in Postgres Docker
  - Connect both using Docker Compose
  - Ensure webapp is using the Postgres container as database.


## ToDo
- [x] Intro more to the point. It is about Docker, not Go.
- [ ] Fix all LINKS
- [ ] Verify if all code is working with current blog commit.


## Old snippets

### Intro

I finaly had some evenings to spend some time on learning GoLang. After the Go tutorial [Games with Go](link), I wanted to learn to ceate a simple webapp. I found this excelent [blog of Soham Kamani](https://www.sohamkamani.com/blog/2017/09/13/how-to-build-a-web-application-in-golang/) where he creates a simple webapp in Go including a database and with proper unit tests. This got me exited and after following his tutorial, I decided to improve this fullstack webapp setup by dockerizing it. It turned out to be quite easy and I will show you how we can achieve this.

For my daily work I'm used to work with the Jetbains IDE's. But since they do not have a community version for Go development I tried VS Code which has a nice Go plugin.

