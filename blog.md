# Dockerize fullstack Go webapp

While I was following a nice tutorial on creating a webapp in Go, I decided to improve this setup by dockerizing it. This turned out to be quite easy, which made me excited to write a blog about it.

## What to do

What are we going to do? The blog is about dockerizing a simple Go webapp. This webapp exist of some Go code and a Postgres database. We can brake the task down into four parts:

- Get Go code running in a Go container
- Get database running in a Postgres container
- Manage both using Docker Compose
- Ensure webapp is using the Postgres container as database.

## Go webapp
Before we proceed we need a little more knowledge of this [webapp](https://github.com/Dirc/webapp_fullstack_go_docker) we are working with.

We can start the webapp.

```bash
# Run web app
go run main.go bird_handlers.go store.go
```

We can verify if it is working on some URLs.

```bash
wget localhost:8080
wget localhost:8080/hello
wget localhost:8080/assets/
```

There is a database called `bird_encyclopedia` with a table `birds` which has two columns `species` and `description`.

More details of this webapp can be found in [this great tutorial of Soham Kamani](https://www.sohamkamani.com/blog/2017/09/13/how-to-build-a-web-application-in-golang/).

## Get Go code running in a Go container

To get the webapp running in a container, we will use the [official GoLang image](https://hub.docker.com/_/golang/). Since we have a simple webapp, we do not need to add much to the example `Dockerfile`. We only add a more recent version of the golang image and expose a port.

```docker
FROM golang:1.11

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 8080

CMD ["app"]
```

The `go get -d` command downloads all dependencies of the webapp and `go install` installs the webapp. After installing the webapp it becomes available as an executable `app`, i.e. the `WORKDIR` directory.

We can now try to build the image and run it.

```bash
docker build -t gowebapp .
docker run -it --rm -p 8080:8080 --name web gowebapp
```

Verify the webapp by browse to one of the URLs.

## Get database running in a Postgres container

For our Postgres database, we take the [the official Postgres image](https://hub.docker.com/_/postgres/).

We can start a Postgres container and add a password and default database.

```bash
# Start Postgres container
docker run --rm --name db -e POSTGRES_PASSWORD=secret -e POSTGRES_DB=bird_encyclopedia -d postgres
```

After it is started, we can connect to it with psql and play around with SQL.

```bash
# Connect using psql
docker run -it --rm --link db:postgres postgres psql -h postgres -U postgres
#> password: secret
```

We can initialize the database.

```sql
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
INSERT INTO birds (species, description) VALUES ('Canary', 'Small yellow bird');
select * from birds;
select species, description from birds;

-- Get info
\d
\d birds
select * from birds;
```

In other words, we can manage the database the way we need it. Except for one last thing. We want the database to initalize with the `birds` table during startup. We can achieve this by creating an `initdb.sql` file which creates this table and put this file in the `/docker-entrypoint-initdb.d` directory in the container. A clumsy but simple way to do this is to add all files of the current directory using \`pwd\`.

```bash
docker run --rm --name db -v `pwd`:/docker-entrypoint-initdb.d -e POSTGRES_PASSWORD=secret -e POSTGRES_DB=bird_encyclopedia -d postgres
```

We can verify the result again with psql.

```bash
# Connect using psql
docker run -it --rm --link db:postgres postgres psql -h postgres -U postgres bird_encyclopedia
#> password: secret
```

## Docker Compose

Now that we have both a standalone Go container and Postgres container running, we can tie them together with Docker Compose.

We take our previous `docker run` commands and put them in a `docker-compose.yml` file. For the `web` service we add `build: .`, so when Docker can not find an image called `gowebapp` locally, it will build the Dockerfile. 

```yaml
version: '3.1'
services:
  web:
    build: .
    image: gowebapp
    ports:
      - 8080:8080
  db:
    image: postgres
    environment:
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: secret
      POSTGRES_DB: bird_encyclopedia
    volumes:
      - ./initdb.sql:/docker-entrypoint-initdb.d/initdb.sql
```

We can now start both services.

```bash
docker-compose up
```

You can verify if the database is initialized using psql and expect an empty `birds` table.

```bash
docker exec -it webappfullstackgodocker_db_1 psql -d bird_encyclopedia -U postgres -c "select * from birds;"
```

## Modify database connection in Go app

Now that we can start both containers together using Docker Compose. To only thing left is to tell the Go webapp where to find the database. For this we need to modify the database connection which is configured in the `main` function in the `main.go` file. The database connection is defined by the keys `host= port= user= password= dbname= sslmode=`. Since we use Docker Compose, the host is equal to the service name `host=db`. Hence we get:

```go
connString := "host=db port=5432 user=postgres password=secret dbname=bird_encyclopedia sslmode=disable"
```

## Ready to Go!

Now that everything is set up. We can start the stack again with `docker-compose up`. Browse to the gui at `http://localhost:8080/assets/` and add some of your favorite birds. Verify if your new birds are really stored in the database by running the psql command again.

## Drafts

### skeleton

- Intro: Ready to learn some Go Lang
  - (After Games With Go)
  - Wanted to create a simple webapp and found this excelent blog.
  - I got exited and decided to spice things up a bit by dockerize this webapp and database setup.
  - Which was surprisingly easy!

- Tool of choice: VS Code. Since it has a nice Go plugin.

- Steps:
  - Follow blog. It starts with a simple `Hello World`. Note that this example is still available at the `/hello` uri. Hence we can use it to verify if the webapp is running properly in the container.
  - Get webapp runnning in Go Docker
  - Get database running in Postgres Docker
  - Connect both using Docker Compose
  - Ensure webapp is using the Postgres container as database.

### ToDo

- [x] Intro more to the point. It is about Docker, not Go.
- [x] Fix all LINK, LINKS
- [ ] Verify if all code is working with current blog commit
- [x] Capital caracters for Docker and Postgres,
- [x] search for Postgress (double ss), Dockerize
- [x] Spelling checker!
- [ ] update docker-compose.yml and verify if it still works!
- [x] Rewrite first paragraph/intro
- [x] Write conclusion
- [ ] Notify Soham Kamani of how he inspired me

### Old snippets

#### Intro

I finaly had some evenings to spend some time on learning GoLang. After the Go tutorial [Games with Go](link), I wanted to learn to ceate a simple webapp. I found this excelent [blog of Soham Kamani](https://www.sohamkamani.com/blog/2017/09/13/how-to-build-a-web-application-in-golang/) where he creates a simple webapp in Go including a database and with proper unit tests. This got me exited and after following his tutorial, I decided to improve this fullstack webapp setup by dockerizing it. It turned out to be quite easy and I will show you how we can achieve this.

For my daily work I'm used to work with the Jetbains IDE's. But since they do not have a community version for Go development I tried VS Code which has a nice Go plugin.

#### sql 

```sql
-- Create the database
CREATE DATABASE bird_encyclopedia;
-- Enter the database
\c bird_encyclopedia
```