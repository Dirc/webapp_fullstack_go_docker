# Drafts

## skeleton

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

## ToDo

- [x] Intro more to the point. It is about Docker, not Go.
- [x] Fix all LINK, LINKS
- [x] Fix initdb.sql
- [x] Verify if all code is working with current blog commit
- [x] Capital caracters for Docker and Postgres,
- [x] search for Postgress (double ss), Dockerize
- [x] Spelling checker!
- [x] update docker-compose.yml and verify if it still works!
- [x] Rewrite first paragraph/intro
- [x] Write conclusion
- [ ] Notify Soham Kamani of how he inspired me

## Old snippets

### Intro

I finaly had some evenings to spend some time on learning GoLang. After the Go tutorial [Games with Go](link), I wanted to learn to ceate a simple webapp. I found this excelent [blog of Soham Kamani](https://www.sohamkamani.com/blog/2017/09/13/how-to-build-a-web-application-in-golang/) where he creates a simple webapp in Go including a database and with proper unit tests. This got me exited and after following his tutorial, I decided to improve this fullstack webapp setup by dockerizing it. It turned out to be quite easy and I will show you how we can achieve this.

For my daily work I'm used to work with the Jetbains IDE's. But since they do not have a community version for Go development I tried VS Code which has a nice Go plugin.

### sql

```sql
-- Create the database
CREATE DATABASE bird_encyclopedia;
-- Enter the database
\c bird_encyclopedia
```