# Idea for blog

## Dockerize Fullstack Go webapp

While I was following a nice tutorial on creating a webapp in Go, I decided to improve this setup by dockerizing it. It turned out to be quite easy and I will show you how you can do this.

## What to do

What are we going to do? The blog is about dockerizing a Go webapp. We can brake the task down into four parts:
  - Get webapp runnning in a Go container
  - Get database running in a Postgres container
  - Manage both using Docker Compose
  - Ensure webapp is using the Postgres container as database.

Before we proceed we need a little more knowledge of this webapp we are working with.

We can start the webapp.

```bash
# Run web app
go run main.go bird_handlers.go store.go
```

We can verify if it is working on these url's.

```bash
wget localhost:8080
wget localhost:8080/hello
wget localhost:8080/assets/
```

Database table is called `bird_encyclopedia` and has two columns `species` and `description`.

More details of this webapp can be found in the blog of Kamani (LINKS).

## Get webapp runnning in a Go container

We will us the [official GoLang image](https://hub.docker.com/_/golang/). Since we have a simple webapp, we do not need to add much to the example `Dockerfile`. We only add a more recent version of the golang image and expose a port.

```docker
FROM golang:1.11

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 8080

CMD ["app"]
```

The `go get -d` command downloads all dependencies of the webapp and `go install` instals the webapp. After installing the webapp it becomes available as an executable `app`, i.e. the workdir directory.

We can now try to build the image and run it.

```bash
docker build -t gowebapp .
docker run -it --rm -p 8080:8080 --name my-running-gowebapp gowebapp
```

Verify the webapp by browse to one of the uri's

# Get database running in a Postgres container


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
- [ ] Intro more to the point. It is about Docker, not Go.
- [ ] Fix all LINKS

## Old snippets

### Intro

I finaly had some evenings to spend some time on learning GoLang. After the Go tutorial [Games with Go](link), I wanted to learn to ceate a simple webapp. I found this excelent [blog of Soham Kamani](https://www.sohamkamani.com/blog/2017/09/13/how-to-build-a-web-application-in-golang/) where he creates a simple webapp in Go including a database and with proper unit tests. This got me exited and after following his tutorial, I decided to improve this fullstack webapp setup by dockerizing it. It turned out to be quite easy and I will show you how we can achieve this.

For my daily work I'm used to work with the Jetbains IDE's. But since they do not have a community version for Go development I tried VS Code which has a nice Go plugin.

